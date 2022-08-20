CREATE TABLE IF NOT EXISTs search_type (
    id serial PRIMARY KEY,
    type VARCHAR(20) NOT NULL DEFAULT 'car'
);

CREATE TABLE IF NOT EXISTS search (
    search_id serial PRIMARY KEY,
    search_type_id integer not null,
    date_in timestamp not null DEFAULT NOW(),
    date_out timestamp not null DEFAULT NOW(),
    FOREIGN KEY (search_type_id) REFERENCES search_type (id) ON DELETE CASCADE
);

INSERt INTO search_type (type) VALUES ('car'), ('flight'), ('hotel');



