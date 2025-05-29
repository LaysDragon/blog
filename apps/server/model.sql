CREATE TABLE
    account (
        id serial PRIMARY KEY,
        created_date timestamptz NOT NULL default current_timestamp,
        updated_date timestamptz NOT NULL default current_timestamp,
        username varchar(20) NOT NULL,
        role varchar(10) NOT NULL,
        email varchar(50) NOT NULL,
        pass_hash varchar(50) NOT NULL
    );

CREATE TABLE
    post (
        id serial PRIMARY KEY,
        created_date timestamptz NOT NULL default current_timestamp,
        updated_date timestamptz NOT NULL default current_timestamp,
        owner_id int NOT NULL REFERENCES account (id),
        content text NOT NULL
    );

CREATE TABLE
    attachtment (
        id serial PRIMARY KEY,
        created_date timestamptz default current_timestamp,
        updated_date timestamptz default current_timestamp,
        owner_id int NOT NULL REFERENCES account (id),
        related_id int NOT NULL REFERENCES post (id),
        url text NOT NULL
    );

CREATE TABLE
    comment (
        id serial PRIMARY KEY,
        created_date timestamptz NOT NULL default current_timestamp,
        post_id int NOT NULL REFERENCES post (id),
        email varchar(50) NOT NULL,
        name varchar(20) NOT NULL,
        content text NOT NULL
    );

CREATE TABLE
    access_log (
        id serial PRIMARY KEY,
        timestamp timestamptz NOT NULL default current_timestamp,
        user_id int REFERENCES account (id),
        method text NOT NULL
    );