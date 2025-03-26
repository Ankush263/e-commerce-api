CREATE TYPE user_role AS ENUM ('seller', 'customer');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    phone VARCHAR NOT NULL UNIQUE,
    role user_role DEFAULT 'customer',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)

--  migrate -path ./db/migration -database "postgres://postgres:Ankush%40postgres263@localhost/e-commerce-api?sslmode=disable" up 1
--  migrate create -ext sql -dir db/migration -seq users -format