CREATE TABLE customers (
  id                serial primary key,
  firstname         text,
  lastname          text,
  email             text,
  phone             text,
  address           text,
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON customers
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();

