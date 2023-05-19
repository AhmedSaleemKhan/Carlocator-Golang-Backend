-- +migrate Up
ALTER TABLE ONLY findr.dlr_staff
       ADD COLUMN profile_image text,
       ADD COLUMN region text,
       ADD COLUMN city text,
       ADD COLUMN state text,
       ADD COLUMN postal_code text,
       ADD COLUMN street1_address text, 
       ADD COLUMN street2_address text, 
       ADD COLUMN currency text;    

-- +migrate Down
ALTER TABLE ONLY findr.dlr_staff
       DROP COLUMN profile_image,
       DROP COLUMN region,
       DROP COLUMN city,
       DROP COLUMN state,
       DROP COLUMN postal_code,
       DROP COLUMN street1_address,
       DROP COLUMN street2_address,
       DROP COLUMN currency;
DROP TABLE IF EXISTS findr.dlr_staff;

