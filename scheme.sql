CREATE TYPE category AS ENUM (
    'harmless',
    'delicate',
    'offensive'
    );

CREATE TABLE statements
(
    id         uuid                     NOT NULL,
    statement  text                     NOT NULL,
    category   category                 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

CREATE MATERIALIZED VIEW game AS
SELECT statements.id,
       row_number() OVER (ORDER BY (random())) AS "position",
       statements.statement,
       statements.category
FROM statements
WITH NO DATA;

CREATE TABLE translations
(
    statement_id uuid                     NOT NULL,
    language     character varying(8)     NOT NULL,
    translation  text                     NOT NULL,
    created_at   timestamp with time zone NOT NULL,
    updated_at   timestamp with time zone NOT NULL
);

ALTER TABLE ONLY statements
    ADD CONSTRAINT statements_pkey PRIMARY KEY (id);

ALTER TABLE ONLY statements
    ADD CONSTRAINT statements_statement_key UNIQUE (statement);

ALTER TABLE ONLY translations
    ADD CONSTRAINT translations_pkey PRIMARY KEY (statement_id, language);

ALTER TABLE ONLY translations
    ADD CONSTRAINT statement_id_fkey FOREIGN KEY (statement_id) REFERENCES public.statements (id) ON UPDATE CASCADE ON DELETE CASCADE;
