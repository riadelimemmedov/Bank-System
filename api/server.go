package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// !NewServer
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts/", server.createAccount)

	server.router = router
	return server
}

// !Start
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// !errorResponse
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
