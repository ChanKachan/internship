CREATE TABLE IF NOT EXISTS public.products (
    id character varying NOT NULL,
    product_name character varying NOT NULL,
    description text,
    characteristics jsonb NOT NULL,
    weight integer NOT NULL,
    barcode character varying NOT NULL,
    PRIMARY KEY (id)
);