package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func produtoTest(t *testing.T) Produto {
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
