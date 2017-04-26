ALTER TABLE retailers RENAME TO consign;
ALTER TABLE consign RENAME COLUMN retailer_no TO customer_no;
