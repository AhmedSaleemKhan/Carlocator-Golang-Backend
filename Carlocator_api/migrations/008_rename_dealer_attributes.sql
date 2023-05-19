
-- +migrate Up

ALTER TABLE findr.sys_dealer RENAME COLUMN street1 TO street1_address;
ALTER TABLE findr.sys_dealer RENAME COLUMN street2 TO street2_address;
ALTER TABLE findr.sys_dealer ALTER COLUMN currency SET DATA TYPE text;
ALTER TABLE findr.sys_dealer ALTER COLUMN created_at DROP NOT NULL;

-- +migrate Down

ALTER TABLE findr.sys_dealer RENAME COLUMN street1_address TO street1;
ALTER TABLE findr.sys_dealer RENAME COLUMN street2_address TO street2;
ALTER TABLE findr.sys_dealer ALTER COLUMN currency SET DATA TYPE text;
