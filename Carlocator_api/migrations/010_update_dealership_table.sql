
-- +migrate Up

ALTER TABLE ONLY findr.sys_dealer
       ADD COLUMN dealership_logo text,
       ADD COLUMN dealership_name text;

-- +migrate Down

ALTER TABLE ONLY findr.sys_dealer
       DROP COLUMN dealership_logo,
       DROP COLUMN dealership_name;