CREATE TABLE wholesales (
  id                  serial primary key,
  retailer_no         text,
  item_no             text,
  collection          text,
  style               text,
  color               text,
  season              text,
  size                text,
  gender              text,
  quantity            int,
  created_at          timestamp with time zone default now(),
  last_modified_at    timestamp with time zone default now()
);

CREATE INDEX retailer_no_wholesales_index ON wholesales (retailer_no);

CREATE TRIGGER update_modtime BEFORE UPDATE ON wholesales
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
