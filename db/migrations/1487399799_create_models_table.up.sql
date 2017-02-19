CREATE TABLE models (
  id                serial primary key,
  name              text,
  gender            text,
  primary_image     text,
  images            text[],
  features          text[],
  types             text[],
  subtypes          text[],
  created_at        timestamp with time zone default now(),
  last_modified_at  timestamp with time zone default now()
);

CREATE TRIGGER update_modtime BEFORE UPDATE ON models
FOR EACH ROW EXECUTE PROCEDURE update_last_modified_at();

