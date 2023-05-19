-- +migrate Up

CREATE TABLE findr.veh_key (
	key_id uuid NOT NULL,
	veh_id uuid NOT NULL,
	loc_id uuid NOT NULL,
	key_name varchar NOT NULL,
	key_accessible varchar NOT NULL,
	received varchar NOT NULL,
	received_from varchar NOT NULL,
	received_by varchar NOT NULL,
	CONSTRAINT veh_key_pkey PRIMARY KEY (key_id)
);
CREATE INDEX "fki_V" ON findr.veh_key USING btree (veh_id);
CREATE INDEX fki_veh_loc ON findr.veh_key USING btree (loc_id);

ALTER TABLE ONLY findr.veh_key 
    ADD CONSTRAINT "veh_id_fkey" foreign key (veh_id) references findr.vehicle(veh_id);

-- +migrate Down

ALTER TABLE ONLY findr.veh_key
    DROP CONSTRAINT IF EXISTS veh_key_pkey;

ALTER TABLE ONLY findr.veh_key 
    DROP CONSTRAINT IF EXISTS veh_id_fkey;
	
ALTER TABLE ONLY findr.veh_key
    DROP CONSTRAINT IF EXISTS fki_V;

DROP TABLE IF EXISTS findr.veh_key;

