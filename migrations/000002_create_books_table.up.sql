CREATE TABLE IF NOT EXISTS books(
   id SERIAL PRIMARY KEY,
   name VARCHAR (255) NOT NULL,
   description VARCHAR (1024) NOT NULL,
   price INT NOT NULL,
   user_id INT NOT NULL REFERENCES users(id),
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP DEFAULT NULL
);
