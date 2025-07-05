CREATE TABLE IF NOT EXISTS public.warehouses (
    id uuid NOT NULL,
    address_id uuid NOT NULL,
    PRIMARY KEY (id)
);