-- +migrate Up

CREATE TABLE findr.veh_assign_key(
    key_assign_id uuid NOT NULL,
    veh_id uuid NOT NULL,
    key_id uuid NOT NULL,
    staff_id uuid NOT NULL,
    key_assign_status bool NOT NULL,
    transfer_id uuid NOT NULL,
    transfere_id uuid NOT NULL,
    assign_cmnt text,
    date_key_assign_start bigint,
    date_key_assign_stop bigint,
    CONSTRAINT veh_pkey PRIMARY KEY (key_assign_id)
);
CREATE INDEX fki_veh_id ON findr.vehicle USING btree (veh_id);
CREATE INDEX fki_key_id ON findr.veh_key USING btree (key_id);
 ALTER TABLE ONLY findr.veh_assign_key
    ADD CONSTRAINT "veh_id_fkey" foreign key (veh_id) references findr.vehicle(veh_id);
 ALTER TABLE ONLY findr.veh_assign_key
    ADD CONSTRAINT "key_id_fkey" foreign key (key_id) references findr.veh_key(key_id);
-- +migrate Down
ALTER TABLE ONLY findr.veh_assign_key
    DROP CONSTRAINT IF EXISTS key_assign_id_pkey;
ALTER TABLE ONLY findr.veh_assign_key
    DROP CONSTRAINT IF EXISTS key_id_fkey;
ALTER TABLE ONLY findr.veh_assign_key
    DROP CONSTRAINT IF EXISTS veh_id_fkey;
DROP TABLE IF EXISTS findr.veh_assign_key;