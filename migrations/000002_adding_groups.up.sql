CREATE TABLE groups (
    name VARCHAR(45) NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE accounts_groups (
    account_id VARCHAR(36) NOT NULL,
    group_id VARCHAR(45) NOT NULL,
    PRIMARY KEY (account_id, group_id),
    FOREIGN KEY (account_id) REFERENCES accounts(username) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(name) ON DELETE CASCADE
);

CREATE TABLE r_articles_groups (
    article_id VARCHAR(36) NOT NULL,
    group_id VARCHAR(45) NOT NULL,
    PRIMARY KEY (article_id, group_id),
    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(name) ON DELETE CASCADE
);

CREATE TABLE w_articles_groups (
    article_id VARCHAR(36) NOT NULL,
    group_id VARCHAR(45) NOT NULL,
    PRIMARY KEY (article_id, group_id),
    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(name) ON DELETE CASCADE
);

ALTER TABLE accounts DROP COLUMN role;
