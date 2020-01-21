
-- +migrate Up
ALTER TABLE company
RENAME COLUMN nome TO name;

-- +migrate Down
