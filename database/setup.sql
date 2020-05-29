-- An account wouldn't be deleted once created.
CREATE TABLE account (
    username VARCHAR(20) PRIMARY KEY,
    hashed_password CHAR(60) NOT NULL,
    email VARCHAR(256) NOT NULL CONSTRAINT email_unique UNIQUE,
    karma INT NOT NULL,
    bio VARCHAR(60) NOT NULL,
    join_time TIME NOT NULL
);

-- A subreddit wouldn't be deleted once created.
CREATE TABLE subreddit (
    sub_name VARCHAR(16) PRIMARY KEY,
    description VARCHAR(512) NOT NULL
);

CREATE TABLE join_sub (
    username VARCHAR(20) NOT NULL,
    sub_name VARCHAR(16) NOT NULL,
    PRIMARY KEY (username, sub_name),
    FOREIGN KEY (username) REFERENCES account (username),
    FOREIGN KEY (sub_name) REFERENCES subreddit (sub_name)
);

-- Articles are allowed to be deleted after creation.
CREATE TABLE article (
    sub_name VARCHAR(16) NOT NULL,
    aid VARCHAR(16) NOT NULL CONSTRAINT aid_unique UNIQUE,
    type VARCHAR(8) NOT NULL,
    body VARCHAR(1024) NOT NULL,
    title VARCHAR(128) NOT NULL,
    points INT NOT NULL,
    posted_by VARCHAR(20) NOT NULL,
    posted_time TIME NOT NULL,
    PRIMARY KEY (sub_name, aid),
    FOREIGN KEY (sub_name) REFERENCES subreddit (sub_name),
    FOREIGN KEY (posted_by) REFERENCES account (username)
);

-- Comments are allowed to be deleted after creation.
-- If the article the comment is in is deleted, the comment will also be deleted.
CREATE TABLE comment (
    sub_name VARCHAR(16) NOT NULL,
    aid VARCHAR(16) NOT NULL,
    cid VARCHAR(16) NOT NULL CONSTRAINT cid_unique UNIQUE,
    body VARCHAR(512) NOT NULL,
    points INT NOT NULL,
    posted_by VARCHAR(20) NOT NULL,
    posted_time TIME NOT NULL,
    PRIMARY KEY (sub_name, aid, cid),
    FOREIGN KEY (sub_name) REFERENCES subreddit (sub_name),
    FOREIGN KEY (aid) REFERENCES article (aid) ON DELETE CASCADE,
    FOREIGN KEY (posted_by) REFERENCES account (username)
);

-- If the saved article is deleted, the save_article entry will also be deleted.
CREATE TABLE save_article (
    username VARCHAR(20) NOT NULL,
    sub_name VARCHAR(16) NOT NULL,
    aid VARCHAR(16) NOT NULL,
    PRIMARY KEY (username, sub_name, aid),
    FOREIGN KEY (username) REFERENCES account (username),
    FOREIGN KEY (sub_name) REFERENCES subreddit (sub_name),
    FOREIGN KEY (aid) REFERENCES article (aid) ON DELETE CASCADE
);
