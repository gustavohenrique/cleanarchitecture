DROP TABLE IF EXISTS todo_items;
CREATE TABLE todo_items (
  id String(36),
  created_at DateTime,
  title String(250),
  is_done Bool
)
ENGINE = MergeTree()
PRIMARY KEY (id);
