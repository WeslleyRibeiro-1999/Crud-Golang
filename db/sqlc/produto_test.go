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

func TestGetProdutos(t *testing.T) {
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
