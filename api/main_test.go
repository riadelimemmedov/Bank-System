package api

import (
	"os"
	db "simplebank/db/sqlc"
	"simplebank/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// !newTestServer
func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: 15,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

// !TestMain
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
