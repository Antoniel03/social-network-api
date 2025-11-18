CREATE TABLE IF NOT EXISTS users (
  id SERIAL NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  lastname TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  birth_date DATE NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS posts(
  id SERIAL NOT NULL PRIMARY KEY,
  user_id INT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  text_content TEXT NOT NULL,
  likes INT, 
  dislikes INT,

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS friend_requests(
  id SERIAL NOT NULL PRIMARY KEY,
  sender_user_id INT NOT NULL, 
  receiver_user_id INT NOT NULL, 
  status TEXT,
  request_date TIMESTAMP NOT NULL,
  friends_since TIMESTAMP,

  FOREIGN KEY (sender_user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (receiver_user_id) REFERENCES users(id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS media(
  post_id INT NOT NULL, 
  id SERIAL NOT NULL PRIMARY KEY,
  user_id INT NOT NULL, 
  location text NOT NULL,
  file_size INT NOT NULL,
  uploaded_at TEXT NOT NULL,

  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments(
  id SERIAL NOT NULL PRIMARY KEY,
  user_id INT NOT NULL, 
  post_id INT NOT NULL, 
  text_content TEXT,
  created_at timestamp NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS interactions(
  id SERIAL NOT NULL PRIMARY KEY, 
  user_id INT NOT NULL,
  post_id INT NOT NULL, 
  interacted_at TIMESTAMP NOT NULL,
  interaction_type TEXT NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

