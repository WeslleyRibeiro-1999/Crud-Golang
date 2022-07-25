package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createdProdutoRandom(t *testing.T) Produto {
	arg := CreateProdutoParams{
		Nome:  "Any_name",
		Preco: 1234,
	}

	produto, err := testQueries.CreateProduto(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, produto)

	require.NotEmpty(t, produto.ID)
	require.NotEmpty(t, produto.CriadoEm)
	require.Equal(t, arg.Nome, produto.Nome)
	require.Equal(t, arg.Preco, produto.Preco)

	return produto
}

func TestCreatedProduto(t *testing.T) {
	createdProdutoRandom(t)
}

func TestGetProduto(t *testing.T) {
	produtoCriado := createdProdutoRandom(t)

	produtoAchado, err := testQueries.GetProduto(context.Background(), produtoCriado.ID)

	require.NoError(t, err)
	require.NotEmpty(t, produtoCriado)
	require.NotEmpty(t, produtoAchado)

	require.Equal(t, produtoCriado.Nome, produtoAchado.Nome)
	require.Equal(t, produtoCriado.Preco, produtoAchado.Preco)
	require.Equal(t, produtoCriado.ID, produtoAchado.ID)
	require.Equal(t, produtoCriado.CriadoEm, produtoAchado.CriadoEm)

}

func TestDeleteProduto(t *testing.T) {
	produtoCriado := createdProdutoRandom(t)

	err := testQueries.DeleteProduto(context.Background(), produtoCriado.ID)

	require.NoError(t, err)
}

func TestUpdateProduto(t *testing.T) {
	produtoCriado := createdProdutoRandom(t)

	arg := UpdateProdutoParams{
		ID:    produtoCriado.ID,
		Nome:  "Outro_nome",
		Preco: 9883,
	}

	produtoMudado, err := testQueries.UpdateProduto(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, produtoCriado.ID, produtoMudado.ID)
	require.Equal(t, arg.Nome, produtoMudado.Nome)
	require.Equal(t, arg.Preco, produtoMudado.Preco)
	require.Equal(t, produtoCriado.CriadoEm, produtoMudado.CriadoEm)
}

func TestGetProdutos(t *testing.T) {
	produtoCriado := createdProdutoRandom(t)

	produtoAchado, err := testQueries.GetProdutos(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, produtoCriado)
	require.NotEmpty(t, produtoAchado)

	for _, produtos := range produtoAchado {
		require.NotEmpty(t, produtos.ID)
		require.NotEmpty(t, produtos.Nome)
		require.NotEmpty(t, produtos.Preco)
		require.NotEmpty(t, produtos.CriadoEm)
	}

}
