CREATE TABLE customers (
    id integer PRIMARY KEY,
    first_name VARCHAR(255),
    last_name  VARCHAR(255),
    email  VARCHAR(255),
    latitude FLOAT,
    longitude FLOAT,
    block_id BIGINT,
    state_fips BIGINT,
    state_code VARCHAR(10),
    block_pop BIGINT
);

COPY customers(id,first_name,last_name,email,latitude,longitude)
FROM '/tmp/customers.csv' DELIMITER ',' CSV HEADER;
