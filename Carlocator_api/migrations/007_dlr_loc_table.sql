-- +migrate Up
CREATE TABLE findr.dlr_loc (
    loc_id uuid NOT NULL,
    dlr_id uuid NOT NULL,
    map_id uuid NOT NULL,
    row_id uuid NOT NULL,
    loc_name varchar NOT NULL,
    show_loc_on_map varchar NOT NULL,
    loc_coord varchar NOT NULL,
    loc_has_spaces varchar NOT NULL,
    number_of_spaces int4 NOT NULL,
    loc_checkboxes varchar NOT NULL,
    one_party_transfer varchar NOT NULL,
    max_num_veh int4 NOT NULL,
    check_in_out varchar NOT NULL,
    loc_cmnt varchar NOT NULL,
    CONSTRAINT loc_id_pkey PRIMARY KEY (loc_id)
);
ALTER TABLE ONLY findr.dlr_loc
    ADD CONSTRAINT "dlr_id_fkey" FOREIGN KEY (dlr_id) REFERENCES findr.sys_dealer(dlr_id);
CREATE TABLE findr.dlr_space (
    space_id uuid NOT NULL,
    loc_id uuid NOT NULL,
    space_name varchar NOT NULL,
    space_cmnt varchar NOT NULL,
    CONSTRAINT space_id_pkey PRIMARY KEY (space_id)
);
CREATE TABLE findr.veh_loc (
    veh_loc_id uuid NOT NULL,
    loc_id uuid NOT NULL,
    veh_id uuid NOT NULL,
    space_id uuid NOT NULL,
    user_in varchar NOT NULL,
    user_out varchar NOT NULL,
    date_in_loc bigint NOT NULL,
    date_out_loc bigint NOT NULL,
    CONSTRAINT veh_loc_pkey PRIMARY KEY (veh_loc_id)
);
ALTER TABLE ONLY findr.veh_loc
    ADD CONSTRAINT "veh_id_fkey" foreign key (veh_id) references findr.vehicle(veh_id);


-- +migrate Down
ALTER TABLE ONLY findr.dlr_loc
    DROP CONSTRAINT IF EXISTS loc_id_pkey,
    DROP CONSTRAINT IF EXISTS dlr_id_fkey;

DROP TABLE IF EXISTS findr.dlr_loc;

ALTER TABLE ONLY findr.dlr_space
    DROP CONSTRAINT IF EXISTS space_id_pkey,
    DROP CONSTRAINT IF EXISTS loc_id_fkey;

DROP TABLE IF EXISTS findr.dlr_space;

ALTER TABLE ONLY findr.veh_loc
    DROP CONSTRAINT IF EXISTS veh_id_fkey,
    DROP CONSTRAINT IF EXISTS veh_loc_id_pkey;

DROP TABLE IF EXISTS findr.veh_loc;