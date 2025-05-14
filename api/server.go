package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

// Server提供所有HTTP请求
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new Http server and setup routing.
func NewServer(store *db.Store) *Server {
	//创建服务器
	server := &Server{store: store}
	router := gin.Default()

	//创建账户
	router.POST("/accounts", server.createAccount)
	//查询账户
	router.GET("/accounts/:id", server.getAccount)
	//批量查询账户
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// 在输入地址上运行HTTP服务器并监听API请求
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
