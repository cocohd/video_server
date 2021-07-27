create table video_info (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    author_id INT,  // 此处不使用外键，遵从第三范式，减少约束
    name TEXT,
    display_ctime TEXT, //显示在用户界面的时间格式，需要额外处理
    create_time DATETIME
);