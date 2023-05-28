PRAGMA auto_vacuum = 1;

DROP TABLE IF EXISTS examples;
CREATE TABLE examples (
  id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(250) NOT NULL CHECK (LENGTH(title) > 0),
  is_done BOOLEAN NOT NULL DEFAULT false
);
CREATE UNIQUE INDEX idx_examples_name ON examples(name);
