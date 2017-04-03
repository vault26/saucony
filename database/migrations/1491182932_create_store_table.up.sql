CREATE TABLE stores (
  id                serial primary key,
  name              text,
  city_th           text,
  city              text,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON stores
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();
