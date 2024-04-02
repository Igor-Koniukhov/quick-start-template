-- +goose Up
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255)        NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP

);

CREATE TABLE chats
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    admin_id   INTEGER REFERENCES users (id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP

);

CREATE TABLE chat_participants
(
    user_id INTEGER REFERENCES users (id),
    chat_id INTEGER REFERENCES chats (id),
    PRIMARY KEY (user_id, chat_id)

);

CREATE TABLE images
(
    id         SERIAL PRIMARY KEY,
    chat_id    INTEGER REFERENCES chats (id),
    user_id    INTEGER REFERENCES users (id),
    file_path  VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP

);

CREATE TABLE messages
(
    id         SERIAL PRIMARY KEY,
    chat_id    INTEGER REFERENCES chats (id),
    user_id    INTEGER REFERENCES users (id),
    content    TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE chat_images
(
    image_id   INTEGER REFERENCES images (id),
    message_id INTEGER REFERENCES messages (id),
    PRIMARY KEY (image_id, message_id)
);


-- +goose Down
DROP TABLE chat_images;
DROP TABLE images;
DROP TABLE messages;
DROP TABLE chat_participants;
DROP TABLE chats;
DROP TABLE users;
