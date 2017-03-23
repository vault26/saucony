CREATE TABLE orders (
  id                serial primary key,
  customer_id       int,
  total_price       numeric,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON orders
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();

