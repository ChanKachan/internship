CREATE TABLE IF NOT EXISTS public.address (
    id character varying NOT NULL,
    city character varying NOT NULL,
    street character varying NOT NULL,
    building character varying NOT NULL,
    PRIMARY KEY (id)
);