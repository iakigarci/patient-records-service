-- +goose Up
-- +goose StatementBegin
INSERT INTO patients (id, name, dni, email, phone, address, created_at, updated_at) VALUES 
    ('501377fc-3977-43da-a171-1ae7a7a69607',
     'John Smith',
     '12345678A',
     'john.smith@example.com',
     '555-0123',
     '{"street": "456 Oak Avenue", "city": "Metropolis", "state": "NY", "zip": "10001"}',
     NOW(), NOW()),
    ('77696893-740b-402c-989a-ff699b81853c',
     'Maria Garcia',
     '87654321B',
     'maria.garcia@example.com',
     '555-0124',
     '{"street": "789 Pine Street", "city": "Riverside", "state": "CA", "zip": "92501"}',
     NOW(), NOW()),
    ('cc579c63-0b5c-43d4-b449-367fef9ca88f',
     'David Johnson',
     '98765432C',
     'david.johnson@example.com',
     '555-0125',
     '{"street": "321 Elm Road", "city": "Springfield", "state": "IL", "zip": "62701"}',
     NOW(), NOW());

INSERT INTO diagnoses (id, patient_id, diagnosis, prescription, diagnosis_date, created_at, updated_at) VALUES 
    ('a32b509a-6adf-4ae8-be04-b443ebe3133b', 
     '501377fc-3977-43da-a171-1ae7a7a69607', 
     'High blood glucose levels detected. Blood sugar at 180 mg/dL.',
     'Prescribed Metformin 500mg twice daily',
     NOW(), NOW(), NOW()),
    ('9a311928-8160-4032-b227-a29d97103fcc',
     '501377fc-3977-43da-a171-1ae7a7a69607',
     'Complete Blood Count shows mild anemia',
     'Iron supplements recommended: Ferrous sulfate 325mg daily',
     NOW(), NOW(), NOW()),
    ('b6e6003a-8c6c-48ac-ac18-e1f630ab5e95',
     '501377fc-3977-43da-a171-1ae7a7a69607', 
     'Chest X-Ray reveals mild bronchitis',
     'Prescribed antibiotic course: Amoxicillin 500mg three times daily for 7 days',
     NOW(), NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM patients 
WHERE id IN (
    '501377fc-3977-43da-a171-1ae7a7a69607',
    '77696893-740b-402c-989a-ff699b81853c',
    'cc579c63-0b5c-43d4-b449-367fef9ca88f'
);

DELETE FROM diagnoses 
WHERE id IN (
    'a32b509a-6adf-4ae8-be04-b443ebe3133b',
    '9a311928-8160-4032-b227-a29d97103fcc',
    'b6e6003a-8c6c-48ac-ac18-e1f630ab5e95'
);
-- +goose StatementEnd
