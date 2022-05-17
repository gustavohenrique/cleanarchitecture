PRAGMA auto_vacuum = 1;

DROP TABLE IF EXISTS agents;
CREATE TABLE agents (
  id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
  serial VARCHAR(36) NOT NULL CHECK (LENGTH(serial) > 0),
  name VARCHAR(30) NOT NULL CHECK (LENGTH(name) > 0),
  hostname VARCHAR(255) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  certificate_expires_at TIMESTAMP,
  certificate_pem VARCHAR NOT NULL DEFAULT '',
  max_jobs INTEGER DEFAULT 10,
  system_stat TEXT NOT NULL DEFAULT '{}'
);
CREATE UNIQUE INDEX idx_agents_serial ON agents (serial);
CREATE INDEX idx_agents_certificate_pem ON agents (certificate_pem);
CREATE INDEX idx_agents_hostname ON agents (hostname);


DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
  id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(100) NOT NULL CHECK (LENGTH(name) > 0),
  permissions TEXT NOT NULL CHECK (LENGTH(permissions) > 0)
);
CREATE UNIQUE INDEX idx_roles_name ON roles(name);


DROP TABLE IF EXISTS groups;
CREATE TABLE groups (
  id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(100) NOT NULL CHECK (LENGTH(name) > 0),
  private_key TEXT NOT NULL DEFAULT '',
  public_key TEXT NOT NULL DEFAULT '',
  roles_ids TEXT NOT NULL
);
CREATE UNIQUE INDEX idx_groups_name ON groups(name);


DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  full_name VARCHAR(30) NOT NULL CHECK (LENGTH(full_name) > 0),
  email VARCHAR(100) NOT NULL,
  avatar VARCHAR(255) NOT NULL DEFAULT '',
  password TEXT NOT NULL DEFAULT '',
  private_key TEXT NOT NULL DEFAULT '',
  public_key TEXT NOT NULL DEFAULT '',
  groups_ids TEXT NOT NULL DEFAULT ''
);
CREATE UNIQUE INDEX idx_users_email ON users(email);
