CREATE TABLE IF NOT EXISTS categories (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(40) NOT NULL,
    parent_id VARCHAR(36) NULL,
    FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS accounts (
    username VARCHAR(36) NOT NULL,
    password VARCHAR(64) NOT NULL,
    salt VARCHAR(36) NOT NULL,
    created TIMESTAMP NOT NULL,
    role VARCHAR(10) NOT NULL,
    PRIMARY KEY (username)
);

CREATE TABLE IF NOT EXISTS articles (
    id VARCHAR(36) PRIMARY KEY,
    category_id VARCHAR(36) NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS commits (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    body TEXT NOT NULL,
    article_id VARCHAR(36) NOT NULL,
    author VARCHAR(36) NOT NULL,
    created TIMESTAMP NOT NULL,
    FOREIGN KEY (author) REFERENCES accounts(username) ON DELETE NO ACTION,
    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE 
);
