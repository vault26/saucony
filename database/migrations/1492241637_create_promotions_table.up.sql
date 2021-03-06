CREATE TABLE promotions (
  id                serial primary key,
  code              text UNIQUE,
  discount_percent  numeric,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON promotions
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();

