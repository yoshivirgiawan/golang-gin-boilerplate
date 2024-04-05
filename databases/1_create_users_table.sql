CREATE TABLE Users (
    id varchar(36) PRIMARY KEY,
    name varchar(255),
    email varchar(255),
    password varchar(255),
    avatar_file_name varchar(255),
    role varchar(255),
    token varchar(255),
    created_at timestamp,
    updated_at timestamp
);