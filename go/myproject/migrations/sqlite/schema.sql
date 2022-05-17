PRAGMA auto_vacuum = 1;

DROP TABLE IF EXISTS todo_items;
CREATE TABLE todo_items (
  id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  title VARCHAR(250) NOT NULL CHECK (LENGTH(title) > 0),
  is_done BOOLEAN NOT NULL DEFAULT false
);
CREATE UNIQUE INDEX idx_todo_title ON todo_items(title);
