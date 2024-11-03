package gapi

import (
	"context"
	"strings"

	"github.com/tunvx/simplebank/pkg/token"
	"github.com/tunvx/simplebank/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
	internalCall        = "internal-call"
)

func (service *Service) authorizeUser(ctx context.Context, accessibleRoles []string) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	var payload *token.Payload

	// Check metadata "internal-call"
	if values := md.Get(internalCall); len(values) > 0 && values[0] == "true" {
		// If it's an internal call, skip token authentication
		payload = &token.Payload{
			Role: util.IServiceRole, // Assign internal service role
		}
	} else {
		// Handle external call with authorization header
		values := md.Get(authorizationHeader)
		if len(values) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		authHeader := values[0]
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization header format")
		}

		authType := strings.ToLower(fields[0])
		if authType != authorizationBearer {
			return nil, status.Errorf(codes.Unauthenticated, "unsupported authorization type ( %s )", authType)
		}

		accessToken := fields[1]
		var err error
		payload, err = service.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid access token: %s", err)
		}
	}

	// Check if the role has permission
	if !hasPermission(payload.Role, accessibleRoles) {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied as role ( %s )", payload.Role)
	}
	return payload, nil
}

func hasPermission(userRole string, accessibleRoles []string) bool {
	for _, role := range accessibleRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
