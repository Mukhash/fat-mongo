create table accounts (
    id SERIAL PRIMARY KEY,
    owner TEXT NOT NULL,
    balance INTEGER NOT NULL,
    currency TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

INSERT INTO ACCOUNTS (owner, balance, currency, created_at)
VALUES ('one', 100, 'USD', current_timestamp);
INSERT INTO ACCOUNTS (owner, balance, currency, created_at)
VALUES ('two', 100, 'USD', current_timestamp);
INSERT INTO ACCOUNTS (owner, balance, currency, created_at)
VALUES ('three', 100, 'USD', current_timestamp);