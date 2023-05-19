-- +migrate Up
ALTER TABLE ONLY findr.sys_dealer
    ADD COLUMN state text;


ALTER TABLE ONLY findr.veh_key
    DROP CONSTRAINT veh_id_fkey;

ALTER TABLE only findr.veh_key
ADD CONSTRAINT veh_id_fkey
    FOREIGN KEY (veh_id)
    REFERENCES findr.vehicle (veh_id)
    ON DELETE CASCADE;




-- +migrate Down
ALTER TABLE ONLY findr.sys_dealer
       DROP COLUMN state;


ALTER TABLE ONLY findr.veh_key
    DROP CONSTRAINT veh_id_fkey;

ALTER TABLE ONLY findr.veh_key
    ADD CONSTRAINT veh_id_fkey 
    FOREIGN KEY (veh_id) 
    REFERENCES findr.vehicle(veh_id);

