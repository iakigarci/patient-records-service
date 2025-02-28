-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, email, password_hash, address, phone, created_at, updated_at) VALUES
    ('8eefd5f2-701d-4f6d-a78b-0b4bec6c096e', 'john.doe+1@example.com', '$2a$12$4ebqCzeiKVOWEM3z.7Wv6u7UN2sItkolsNdSwohbYZS4P3WOaVpmm', '{"street": "123 Main St", "city": "Anytown", "state": "CA", "zip": "12345"}', '1234567890', NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users 
WHERE id = '8eefd5f2-701d-4f6d-a78b-0b4bec6c096e';
-- +goose StatementEnd
