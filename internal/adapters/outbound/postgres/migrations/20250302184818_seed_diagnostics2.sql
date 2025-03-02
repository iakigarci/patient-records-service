-- +goose Up
-- +goose StatementBegin
INSERT INTO diagnoses (id, patient_id, diagnosis, prescription, diagnosis_date, created_at, updated_at) VALUES 
    ('d8f2d03e-7e73-4f1b-a7bc-d3a6892e9b6f',
     '77696893-740b-402c-989a-ff699b81853c',
     'Seasonal allergies with rhinitis',
     'Prescribed Cetirizine 10mg daily',
     '2025-03-01 11:30:00', NOW(), NOW()),
    ('e9b5c31a-8d4f-4b2c-9e6d-f8a7b6c5d4e3',
     '77696893-740b-402c-989a-ff699b81853c',
     'Migraine headaches',
     'Sumatriptan 50mg as needed for acute episodes',
     '2025-03-02 16:45:00', NOW(), NOW()),
    
    ('f1c2d3e4-5b6a-7c8d-9e0f-1a2b3c4d5e6f',
     'cc579c63-0b5c-43d4-b449-367fef9ca88f',
     'Lower back strain',
     'Physical therapy recommended, Ibuprofen 400mg as needed',
     '2025-03-01 13:20:00', NOW(), NOW()),
    ('a1b2c3d4-e5f6-7a8b-9c0d-1e2f3a4b5c6d',
     'cc579c63-0b5c-43d4-b449-367fef9ca88f',
     'Hypertension follow-up',
     'Continue Lisinopril 10mg daily, maintain low-sodium diet',
     '2025-03-03 10:45:00', NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM diagnoses 
WHERE id IN (
    'd8f2d03e-7e73-4f1b-a7bc-d3a6892e9b6f',
    'e9b5c31a-8d4f-4b2c-9e6d-f8a7b6c5d4e3',
    'f1c2d3e4-5b6a-7c8d-9e0f-1a2b3c4d5e6f',
    'a1b2c3d4-e5f6-7a8b-9c0d-1e2f3a4b5c6d'
);
-- +goose StatementEnd
