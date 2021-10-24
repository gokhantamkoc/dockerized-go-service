CREATE TABLE landmarks
(
    id      INTEGER NOT NULL,
    name    VARCHAR NOT NULL,
    country VARCHAR NOT NULL,
    CONSTRAINT landmarks_pk PRIMARY KEY (id)
);

INSERT INTO landmarks(id, name, country) values(1, 'Ataturk''s Mausoleum', 'Turkey');
INSERT INTO landmarks(id, name, country) values(2, 'The Great Wall', 'China');