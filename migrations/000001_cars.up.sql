CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(255) UNIQUE NOT NULL,
    mark VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    year INTEGER NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    owner_surname VARCHAR(255) NOT NULL,
    owner_patronymic VARCHAR(255)
);