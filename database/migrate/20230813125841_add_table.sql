-- +goose Up
CREATE TABLE IF NOT EXISTS notifications (
  id VARCHAR(100) PRIMARY KEY,
  channel VARCHAR(100),
  message TEXT,
  created_at timestamp NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS notifications;
