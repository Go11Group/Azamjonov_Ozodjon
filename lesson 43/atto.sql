-- Create users table
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       name VARCHAR(100),
                       phone VARCHAR(20),
                       age INT
);

-- Create card table
CREATE TABLE card (
                      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                      number VARCHAR(20),
                      user_id UUID REFERENCES users(id)
);

-- Create transaction table
CREATE TABLE transaction (
                             id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             card_id UUID REFERENCES card(id),
                             amount DECIMAL(10, 2),
                             terminal_id UUID DEFAULT NULL,
                             transaction_type VARCHAR(6) CHECK (transaction_type IN ('credit', 'debit'))
);

-- Create station table
CREATE TABLE station (
                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         name VARCHAR(100)
);

-- Create terminal table
CREATE TABLE terminal (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          station_id UUID REFERENCES station(id)
);



-- Insert data into users table
INSERT INTO users (id, name, phone, age) VALUES
                                             (uuid_generate_v4(), 'Alice', '555-1234', 30),
                                             (uuid_generate_v4(), 'Bob', '555-5678', 25),
                                             (uuid_generate_v4(), 'Charlie', '555-8765', 35),
                                             (uuid_generate_v4(), 'Diana', '555-4321', 28),
                                             (uuid_generate_v4(), 'Edward', '555-9876', 22);

-- Get user ids
SELECT id INTO TEMP user_ids FROM users;

-- Insert data into card table using user ids
INSERT INTO card (id, number, user_id) VALUES
                                           (uuid_generate_v4(), '1111-2222-3333-4444', (SELECT id FROM user_ids LIMIT 1 OFFSET 0)),
                                           (uuid_generate_v4(), '5555-6666-7777-8888', (SELECT id FROM user_ids LIMIT 1 OFFSET 1)),
                                           (uuid_generate_v4(), '9999-0000-1111-2222', (SELECT id FROM user_ids LIMIT 1 OFFSET 2)),
                                           (uuid_generate_v4(), '3333-4444-5555-6666', (SELECT id FROM user_ids LIMIT 1 OFFSET 3)),
                                           (uuid_generate_v4(), '7777-8888-9999-0000', (SELECT id FROM user_ids LIMIT 1 OFFSET 4));

-- Get card ids
SELECT id INTO TEMP card_ids FROM card;

-- Insert data into station table
INSERT INTO station (id, name) VALUES
                                   (uuid_generate_v4(), 'Station A'),
                                   (uuid_generate_v4(), 'Station B'),
                                   (uuid_generate_v4(), 'Station C'),
                                   (uuid_generate_v4(), 'Station D'),
                                   (uuid_generate_v4(), 'Station E');

-- Get station ids
SELECT id INTO TEMP station_ids FROM station;

-- Insert data into terminal table using station ids
INSERT INTO terminal (id, station_id) VALUES
                                          (uuid_generate_v4(), (SELECT id FROM station_ids LIMIT 1 OFFSET 0)),
                                          (uuid_generate_v4(), (SELECT id FROM station_ids LIMIT 1 OFFSET 1)),
                                          (uuid_generate_v4(), (SELECT id FROM station_ids LIMIT 1 OFFSET 2)),
                                          (uuid_generate_v4(), (SELECT id FROM station_ids LIMIT 1 OFFSET 3)),
                                          (uuid_generate_v4(), (SELECT id FROM station_ids LIMIT 1 OFFSET 4));

-- Get terminal ids
SELECT id INTO TEMP terminal_ids FROM terminal;

-- Insert data into transaction table using card ids and terminal ids
INSERT INTO transaction (id, card_id, amount, terminal_id, transaction_type) VALUES
                                                                                 (uuid_generate_v4(), (SELECT id FROM card_ids LIMIT 1 OFFSET 0), 50.00, (SELECT id FROM terminal_ids LIMIT 1 OFFSET 0), 'credit'),
                                                                                 (uuid_generate_v4(), (SELECT id FROM card_ids LIMIT 1 OFFSET 1), 20.00, (SELECT id FROM terminal_ids LIMIT 1 OFFSET 1), 'debit'),
                                                                                 (uuid_generate_v4(), (SELECT id FROM card_ids LIMIT 1 OFFSET 2), 75.00, (SELECT id FROM terminal_ids LIMIT 1 OFFSET 2), 'credit'),
                                                                                 (uuid_generate_v4(), (SELECT id FROM card_ids LIMIT 1 OFFSET 3), 10.00, (SELECT id FROM terminal_ids LIMIT 1 OFFSET 3), 'debit'),
                                                                                 (uuid_generate_v4(), (SELECT id FROM card_ids LIMIT 1 OFFSET 4), 100.00, (SELECT id FROM terminal_ids LIMIT 1 OFFSET 4), 'credit');
