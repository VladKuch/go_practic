CREATE TABLE IF NOT EXISTs unit_type (
    id serial PRIMARY KEY,
    type VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS unit (
    unit_id serial PRIMARY KEY,
    type_id integer not null,
    unit_name varchar(255) not null,
    weight int not null DEFAULT 1,
    created_at timestamp not null DEFAULT NOW(),
    FOREIGN KEY (type_id) REFERENCES unit_type (id) ON DELETE CASCADE
);

INSERt INTO unit_type (type) VALUES ('Tab'), ('Core'), ('Inline');



