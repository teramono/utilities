-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE workspaces (
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE UNIQUE INDEX idx_workspaces_name ON workspaces(name);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP INDEX idx_workspaces_name;

DROP TABLE workspaces;

