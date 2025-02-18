package logger

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

// Metadata chứa thông tin cần thiết như User-Agent và Client-IP.
type Metadata struct {
	UserAgent string
	ClientIP  string
}

// extractMetadata trích xuất các metadata cần thiết từ gRPC context, bao gồm User-Agent và Client-IP.
func extractGrpcMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Lấy thông tin User-Agent từ gRPC Gateway hoặc trực tiếp từ header của gRPC
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		} else if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// Lấy địa chỉ IP từ header x-forwarded-for (trong trường hợp có proxy)
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			ips := strings.Split(clientIPs[0], ",")
			mtdt.ClientIP = strings.TrimSpace(ips[0])
		}
	}

	// Lấy địa chỉ IP trực tiếp từ peer nếu chưa tìm thấy trong x-forwarded-for
	if p, ok := peer.FromContext(ctx); ok && mtdt.ClientIP == "" {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}

// GrpcLogger là middleware ghi log cho gRPC requests.
func GrpcLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// Ghi nhận thời điểm bắt đầu request
	startTime := time.Now()

	// Trích xuất metadata từ context và cập nhật context với thông tin IP, User-Agent
	mtdt := extractGrpcMetadata(ctx)

	// Gọi handler để tiếp tục xử lý request
	result, err := handler(ctx, req)

	// Tính toán thời gian xử lý request
	duration := time.Since(startTime)

	// Lấy status code từ error nếu có
	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	// Bắt đầu ghi log với cấp độ Info, chuyển sang Error nếu có lỗi
	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	// Ghi log chi tiết về gRPC request
	logger.Str("protocol", "grpc").
		Str("client_ip", mtdt.ClientIP).
		Str("user_agent", mtdt.UserAgent).
		Str("method", info.FullMethod).
		Int("status_code", int(statusCode)).
		Str("status_text", statusCode.String()).
		Dur("duration", duration).
		Msg("received a gRPC request")

	return result, err
}

// extractHTTPMetadata trích xuất metadata từ HTTP requests.
func extractHTTPMetadata(req *http.Request) *Metadata {
	mtdt := &Metadata{}

	// Lấy User-Agent từ header
	mtdt.UserAgent = req.Header.Get(userAgentHeader)

	// Lấy địa chỉ IP từ X-Forwarded-For (nếu qua proxy) hoặc trực tiếp từ req.RemoteAddr
	if clientIP := req.Header.Get(xForwardedForHeader); clientIP != "" {
		ips := strings.Split(clientIP, ",")
		mtdt.ClientIP = strings.TrimSpace(ips[0])
	} else {
		// Trường hợp không có X-Forwarded-For thì lấy từ RemoteAddr
		mtdt.ClientIP, _, _ = net.SplitHostPort(req.RemoteAddr)
	}
	return mtdt
}

// ResponseRecorder là một custom http.ResponseWriter để ghi lại response status code và body.
type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int    // Ghi lại HTTP status code
	Body       []byte // Ghi lại nội dung response body
}

// WriteHeader ghi lại status code.
func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

// Write ghi lại response body.
func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}

// HttpLogger là middleware ghi log cho HTTP requests.
func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Ghi nhận thời điểm bắt đầu request
		startTime := time.Now()

		// Trích xuất metadata và cập nhật context với thông tin IP, User-Agent
		mtdt := extractHTTPMetadata(req)

		// Tạo ResponseRecorder để ghi lại chi tiết response
		rec := &ResponseRecorder{
			ResponseWriter: res,
			StatusCode:     http.StatusOK, // Mặc định là 200 OK
		}

		// Gọi handler để xử lý request
		handler.ServeHTTP(rec, req)

		// Tính toán thời gian xử lý request
		duration := time.Since(startTime)

		// Bắt đầu ghi log với cấp độ Info, chuyển sang Error nếu status code không phải 200
		logger := log.Info()
		if rec.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", rec.Body)
		}

		// Ghi log chi tiết về HTTP request
		logger.Str("protocol", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Str("client_ip", mtdt.ClientIP).
			Str("user_agent", mtdt.UserAgent).
			Int("status_code", rec.StatusCode).
			Str("status_text", http.StatusText(rec.StatusCode)).
			Dur("duration", duration).
			Msg("received a HTTP request")
	})
}
