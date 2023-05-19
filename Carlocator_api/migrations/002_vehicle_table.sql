-- +migrate Up

CREATE TABLE IF NOT EXISTS findr.vehicle (
    veh_id uuid NOT NULL,
    dlr_id uuid NOT NULL,
    car_status_id uuid,
    veh_type_id uuid,
    attach_id uuid,
	map_color text,
    veh_name text,
	stock_no bigint,
    vin text,
    lic_plate text,
    make text,
    model text,
    year text,
	veh_color text,
    transmission_style text,
    note text,
	created_at bigint NOT NULL
);

ALTER TABLE ONLY findr.vehicle
    ADD CONSTRAINT "vehicle_pkey" PRIMARY KEY (veh_id);


ALTER TABLE ONLY findr.vehicle
    ADD CONSTRAINT "dealer_id_fkey" FOREIGN KEY (dlr_id) REFERENCES findr.sys_dealer(dlr_id);

-- +migrate Down

ALTER TABLE ONLY findr.vehicle
    DROP CONSTRAINT IF EXISTS dealer_id_fkey;

ALTER TABLE ONLY findr.vehicle
    DROP CONSTRAINT IF EXISTS vehicle_pkey;

DROP TABLE IF EXISTS findr.vehicle;