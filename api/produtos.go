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

	ctx.JSON(http.StatusCreated, produto)
}

type getProdutoRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getProduto(ctx *gin.Context) {
	var req getProdutoRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	produto, err := server.store.GetProduto(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, produto)
}

type deleteProdutoRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteProduto(ctx *gin.Context) {
	var req deleteProdutoRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteProduto(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

type updateProdutoRequest struct {
	ID    int32  `json:"id" binding:"required"`
	Nome  string `json:"nome"`
	Preco int32  `json:"preco"`
}

func (server *Server) updateProduto(ctx *gin.Context) {
	var req updateProdutoRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateProdutoParams{
		ID:    req.ID,
		Nome:  req.Nome,
		Preco: req.Preco,
	}

	produto, err := server.store.UpdateProduto(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, produto)
}

func (server *Server) getProdutos(ctx *gin.Context) {
	produtos, err := server.store.GetProdutos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, produtos)
}
