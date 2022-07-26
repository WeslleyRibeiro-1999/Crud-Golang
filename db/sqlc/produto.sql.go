// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: produto.sql

package db

import (
	"context"
)

const createProduto = `-- name: CreateProduto :one
INSERT INTO produtos (
    nome,
    preco
) VALUES (
    $1, $2
) RETURNING id, nome, preco, criado_em
`

type CreateProdutoParams struct {
	Nome  string `json:"nome"`
	Preco int32  `json:"preco"`
}

func (q *Queries) CreateProduto(ctx context.Context, arg CreateProdutoParams) (Produto, error) {
	row := q.db.QueryRowContext(ctx, createProduto, arg.Nome, arg.Preco)
	var i Produto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Preco,
		&i.CriadoEm,
	)
	return i, err
}

const deleteProduto = `-- name: DeleteProduto :exec
DELETE FROM produtos
WHERE id = $1
`

func (q *Queries) DeleteProduto(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProduto, id)
	return err
}

const getProduto = `-- name: GetProduto :one
SELECT id, nome, preco, criado_em FROM produtos
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduto(ctx context.Context, id int32) (Produto, error) {
	row := q.db.QueryRowContext(ctx, getProduto, id)
	var i Produto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Preco,
		&i.CriadoEm,
	)
	return i, err
}

const getProdutos = `-- name: GetProdutos :many
SELECT id, nome, preco, criado_em FROM produtos
`

func (q *Queries) GetProdutos(ctx context.Context) ([]Produto, error) {
	rows, err := q.db.QueryContext(ctx, getProdutos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Produto{}
	for rows.Next() {
		var i Produto
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Preco,
			&i.CriadoEm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduto = `-- name: UpdateProduto :one
UPDATE produtos
SET nome = $2, preco = $3
WHERE id = $1 RETURNING id, nome, preco, criado_em
`

type UpdateProdutoParams struct {
	ID    int32  `json:"id"`
	Nome  string `json:"nome"`
	Preco int32  `json:"preco"`
}

func (q *Queries) UpdateProduto(ctx context.Context, arg UpdateProdutoParams) (Produto, error) {
	row := q.db.QueryRowContext(ctx, updateProduto, arg.ID, arg.Nome, arg.Preco)
	var i Produto
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Preco,
		&i.CriadoEm,
	)
	return i, err
}
