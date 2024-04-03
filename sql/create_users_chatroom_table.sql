CREATE TABLE users_chatroom (
    id SERIAL PRIMARY KEY,
    CONSTRAINT fk_users_id FOREIGN KEY (id) REFERENCES Users(id),
    CONSTRAINT fk_chatroom_id FOREIGN KEY (id) REFERENCES chatroom(id)
);
