package api

import (
	db "github.com/WeslleyRibeiro-1999/Crud-Golang/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.ExecuteStore
	router *gin.Engine
}

func InstanceServer(store *db.ExecuteStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Aqui vai ficar as rotas
	router.POST("/produto", server.createProduto)
	router.GET("/produto/:id", server.getProduto)
	router.GET("/produtos", server.getProdutos)
	router.DELETE("/produto/:id", server.deleteProduto)
	router.PUT("/produto", server.updateProduto)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api started with a error:": err.Error()}
}
