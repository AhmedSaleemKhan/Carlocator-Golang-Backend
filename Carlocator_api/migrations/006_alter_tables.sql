
-- +migrate Up

ALTER TABLE ONLY findr.vehicle
    ADD COLUMN veh_available bool default true;

ALTER TABLE ONLY findr.veh_key
    ADD COLUMN key_available bool default true;

-- +migrate Down
ALTER TABLE ONLY findr.vehicle
    DROP COLUMN  veh_available;
	
ALTER TABLE ONLY findr.veh_key
    DROP COLUMN  key_available;