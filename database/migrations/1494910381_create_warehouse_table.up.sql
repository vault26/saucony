CREATE TABLE warehouse (
  id                  serial primary key,
  customer_no         text,
  customer_name       text,
  customer_name_2     text,
  item_category_code  text,
  item_no             text,
  collection          text,
  style               text,
  color               text,
  size                text,
  gender              text,
  quantity            int,
  average_cost        numeric,
  vn_number           text,
  season              text,
  created_at          timestamp with time zone default now(),
  last_modified_at    timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON warehouse
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
