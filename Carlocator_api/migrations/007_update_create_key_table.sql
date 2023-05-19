
-- +migrate Up

ALTER TABLE ONLY findr.veh_key 
    add COLUMN if not exists loc_name text;

ALTER TABLE ONLY findr.veh_key 
    DROP COLUMN if exists loc_id;
	
ALTER TABLE ONLY findr.veh_key 
    DROP COLUMN if exists received;

	
-- +migrate Down

ALTER TABLE ONLY findr.veh_key 
    DROP COLUMN if exists loc_name;


ALTER TABLE ONLY findr.veh_key 
    add COLUMN if not exists loc_id uuid;
	
ALTER TABLE ONLY findr.veh_key 
    add COLUMN if not exists received varchar;