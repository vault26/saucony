CREATE TABLE order_lines (
  id                serial primary key,
  order_id          int,
  product_id        int,
  size              text,
  quantity          int,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON order_lines
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();

