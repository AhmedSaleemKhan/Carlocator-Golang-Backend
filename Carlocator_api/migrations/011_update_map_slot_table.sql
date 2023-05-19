
-- +migrate Up

ALTER TABLE ONLY findr.map_slot
    ADD COLUMN car_color text;
ALTER TABLE findr.map_slot ALTER COLUMN veh_id DROP NOT NULL;
-- +migrate Down

ALTER TABLE ONLY findr.map_slot
       DROP COLUMN car_color;
