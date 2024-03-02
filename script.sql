CREATE TABLE clientes (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  limite INTEGER NOT NULL
);

CREATE TABLE transacoes (
  id SERIAL PRIMARY KEY,
  valor INTEGER,
  tipo CHAR(1) CHECK (type IN ('c', 'd')),
  descricao VARCHAR(10),
  cliente_id INTEGER REFERENCES clientes(id),
  realizado_em VARCHAR(27)
);

ALTER TABLE
  transacoes
SET
  (autovacuum_enabled = false);


INSERT INTO clientes (nome, limite)
  VALUES
    ('o barato sai caro', 1000 * 100),
    ('zan corp ltda', 800 * 100),
    ('les cruders', 10000 * 100),
    ('padaria joia de cocaia', 100000 * 100),
    ('kid mais', 5000 * 100);
SET GLOBAL max_connections = 10000;

DO $$
BEGIN
  INSERT INTO clientes (nome, limite)
  VALUES
    ('o barato sai caro', 1000 * 100),
    ('zan corp ltda', 800 * 100),
    ('les cruders', 10000 * 100),
    ('padaria joia de cocaia', 100000 * 100),
    ('kid mais', 5000 * 100);
END; $$