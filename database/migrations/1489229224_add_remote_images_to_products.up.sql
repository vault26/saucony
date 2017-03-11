ALTER TABLE products
  ADD COLUMN primary_remote_image text,
  ADD COLUMN remote_images text[];
