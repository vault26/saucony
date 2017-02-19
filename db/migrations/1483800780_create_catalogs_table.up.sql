CREATE TABLE catalogs (
  id                serial primary key,
  item_no           text,
  barcode           text,
  price             numeric,
  model             text,
  color             text,
  size              text,
  gender            text,
  percent_discount  numeric,
  quantity          int,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON catalogs
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
