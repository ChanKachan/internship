CREATE TABLE IF NOT EXISTS public.analytics (
    warehouse_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity_of_products_sold integer NOT NULL DEFAULT 0,
    price_of_products_sold integer NOT NULL DEFAULT 0,
    PRIMARY KEY (warehouse_id, product_id)
);