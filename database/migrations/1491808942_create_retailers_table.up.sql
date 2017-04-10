CREATE TABLE retailers (
  id                  serial primary key,
  retailer_no         text,
  name                text,
  name_2              text,
  item_category_code  text,
  collection          text,
  item_no             text,
  style               text,
  color               text,
  size                text,
  quantity            int,
  created_at          timestamp with time zone default now(),
  last_modified_at    timestamp with time zone default now()
);

CREATE INDEX retailer_no_retailers_index ON retailers (retailer_no);

CREATE TRIGGER update_modtime BEFORE UPDATE ON retailers
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
