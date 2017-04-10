CREATE TABLE stores (
  id                serial primary key,
  retailer_no       text,
  retailer_no_2     text,
  name              text,
  city_th           text,
  phone             text,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE INDEX retailer_no_stores_index ON stores (retailer_no);
CREATE INDEX retailer_no_2_stores_index ON stores (retailer_no);

CREATE TRIGGER update_modtime BEFORE UPDATE ON stores
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
