package api

import (
	"net/http"

	db "github.com/WeslleyRibeiro-1999/Crud-Golang/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createProdutoRequest struct {
	Nome  string `json:"nome" binding:"required"`
	Preco int32  `json:"preco" binding:"required"`
}

func (server *Server) createProduto(ctx *gin.Context) {
	var req createProdutoRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateProdutoParams{
		Nome:  req.Nome,
		Preco: req.Preco,
	}

	produto, err := server.store.CreateProduto(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, produto)
}
