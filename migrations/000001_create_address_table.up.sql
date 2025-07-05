CREATE TABLE IF NOT EXISTS public.address (
    id uuid NOT NULL,
    city character varying NOT NULL,
    street character varying NOT NULL,
    building character varying NOT NULL,
    PRIMARY KEY (id)
);