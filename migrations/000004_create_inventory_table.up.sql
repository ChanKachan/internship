CREATE TABLE IF NOT EXISTS public.inventory (
    warehouse_id uuid NOT NULL REFERENCES public.warehouses(id),
    product_id uuid NOT NULL REFERENCES public.products(id),
    quantity_of_product integer NOT NULL CHECK (quantity_of_product >= 0),
    price integer NOT NULL CHECK (price >= 0),
    percentage_discount_from_price integer DEFAULT 0 CHECK (percentage_discount_from_price BETWEEN 0 AND 100),
    PRIMARY KEY (warehouse_id, product_id)
);