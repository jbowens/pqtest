CREATE TABLE foos (
    id   serial NOT NULL PRIMARY KEY,
    name text   NOT NULL
);

CREATE TABLE bars (
    id     serial NOT NULL PRIMARY KEY,
    foo_id integer
);
