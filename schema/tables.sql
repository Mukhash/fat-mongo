CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    BODY TEXT,
    CONSTRAINT title_unique UNIQUE (title)
);

INSERT INTO tasks (title, body)
VALUES ('Mourning', 'Make bed, brush teeth, have breakfast');
INSERT INTO tasks (title, body)
VALUES ('Workout', 'Pullups, jogging');
INSERT INTO tasks (title, body)
VALUES ('Work', 'Postgre, database/sql');