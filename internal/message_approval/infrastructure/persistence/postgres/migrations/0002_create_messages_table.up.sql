-- Create messages table
CREATE TABLE messages (
                          id VARCHAR PRIMARY KEY,
                          content TEXT NOT NULL,
                          status VARCHAR NOT NULL,
                          sender_id VARCHAR NOT NULL,
                          receiver_id VARCHAR NOT NULL,
                          FOREIGN KEY (sender_id) REFERENCES users(id),
                          FOREIGN KEY (receiver_id) REFERENCES users(id)
);
