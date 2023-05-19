
-- +migrate Up

ALTER TABLE findr.dlr_staff ALTER COLUMN dlr_staff_no SET DATA TYPE text;

-- +migrate Down

ALTER TABLE findr.dlr_staff ALTER COLUMN dlr_staff_no SET DATA TYPE bigint;
