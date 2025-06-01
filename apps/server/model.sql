CREATE TABLE IF NOT EXISTS account (
    id serial PRIMARY KEY,
    created_at timestamptz NOT NULL default current_timestamp,
    updated_at timestamptz NOT NULL default current_timestamp,
    username varchar(20) NOT NULL,
    role varchar(10) NOT NULL,
    email varchar(50) NOT NULL,
    passwd_hash varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS site (
    id serial PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp,
    name varchar(50)
);

CREATE TABLE IF NOT EXISTS site_role (
    PRIMARY KEY (account_id, site_id),
    account_id int NOT NULL REFERENCES account (id),
    site_id int NOT NULL REFERENCES site (id),
    role varchar(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS post (
    id serial PRIMARY KEY,
    created_at timestamptz NOT NULL default current_timestamp,
    updated_at timestamptz NOT NULL default current_timestamp,
    site_id int NOT NULL REFERENCES site (id),
    content text NOT NULL
);

CREATE TABLE IF NOT EXISTS attachtment (
    id serial PRIMARY KEY,
    created_at timestamptz NOT NULL default current_timestamp,
    updated_at timestamptz NOT NULL default current_timestamp,
    site_id int NOT NULL REFERENCES site (id),
    --TODO: 關聯的文章id，設計有些問題，這可能要從另一個關聯表進行管理，需要程序能在文章儲存時自動檢測引用哪些文件
    related_id int NOT NULL REFERENCES post (id),
    url text NOT NULL
);

CREATE TABLE IF NOT EXISTS comment (
    id serial PRIMARY KEY,
    created_at timestamptz NOT NULL default current_timestamp,
    post_id int NOT NULL REFERENCES post (id),
    email varchar(50) NOT NULL,
    name varchar(20) NOT NULL,
    content text NOT NULL
);

CREATE TABLE IF NOT EXISTS access_log (
    id serial PRIMARY KEY,
    created_at timestamptz NOT NULL default current_timestamp,
    user_id int REFERENCES account (id),
    method text NOT NULL
);