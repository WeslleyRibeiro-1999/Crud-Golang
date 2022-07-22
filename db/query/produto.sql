-- name: CreateProduto :one
INSERT INTO produtos (
    nome,
    preco
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetProduto :one
SELECT * FROM produtos
WHERE id = $1 LIMIT 1;

-- name: GetProdutos :many
SELECT * FROM produtos;

-- name: UpdateProduto :one
UPDATE produtos
SET nome = $2, preco = $3
WHERE id = $1 RETURNING *;

-- name: DeleteProduto :exec
DELETE FROM produtos
WHERE id = $1;