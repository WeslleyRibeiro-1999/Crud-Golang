CREATE TABLE "produtos" (
    "id" serial PRIMARY KEY NOT NULL,
    "nome" varchar NOT NULL,
    "preco" integer NOT NULL,
    "criado_em" timestamptz NOT NULL DEFAULT (now())
);