package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tunvx/simplebank/common/util"
)

var (
	publicKeyBase64  = "w4z+16OqrZddIkrPPcmnsVHerhZZ8hGPAoOFOrlTpfs="
	privateKeyBase64 = "gMYc2NpTvsyAahM66jzV2V/MAf6CgonGOKJftNhnM4DDjP7Xo6qtl10iSs89yaexUd6uFlnyEY8Cg4U6uVOl+w=="
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(publicKeyBase64, privateKeyBase64)
	require.NoError(t, err)

	agentID := util.RandomInt64ID()
	shardID := int32(1)
	role := util.CustomerRole
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(agentID, shardID, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.Equal(t, agentID, payload.UserID)
	require.Equal(t, shardID, payload.ShardID)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	agentID := util.RandomInt64ID()
	shardID := int32(1)
	role := util.CustomerRole
	duration := time.Minute

	maker, err := NewPasetoMaker(publicKeyBase64, privateKeyBase64)
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(agentID, shardID, role, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
