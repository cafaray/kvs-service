-- SELECT 'CREATE DATABASE kvs' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'kvs')\gexec

CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS element (
    id serial NOT NULL,
    user_id int NOT NULL,
    key VARCHAR(30) NOT NULL,
    value text NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_kvs PRIMARY KEY(id),
    CONSTRAINT fk_posts_users FOREIGN KEY(user_id) REFERENCES users(id)
);
