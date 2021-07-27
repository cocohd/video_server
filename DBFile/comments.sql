create table comments (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    video_id VARCHAR(64),
    author_id INT,
    content TEXT,
    time DATETIME
);