CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)

--  migrate -path ./db/migration -database "postgres://postgres:Ankush%40postgres263@localhost/e-commerce-api?sslmode=disable" up 1
--  migrate create -ext sql -dir db/migration -seq users -format