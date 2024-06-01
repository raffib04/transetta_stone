CREATE DATABASE transetta_stone;

USE transetta_stone;

CREATE TABLE rooms (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  created_at TIMESTAMP
);

CREATE TABLE messages (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  room_id INTEGER,
  text VARCHAR(255),
  original_language VARCHAR(50),
  creator_name VARCHAR(255),
  created_at TIMESTAMP,
  FOREIGN KEY (room_id) REFERENCES rooms(id)
);
