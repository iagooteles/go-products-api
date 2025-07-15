CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    product_name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL
);
