-- +migrate Up

CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS findr.map (
	map_id uuid NOT NULL,
	dlr_id uuid NOT NULL,
	map_name text,
	map_polygon GEOMETRY
);

CREATE TABLE IF NOT EXISTS findr.map_parking_lot (
	parking_lot_id uuid NOT NULL,
	map_id uuid NOT NULL,
	parking_name text,
	parking_polygon GEOMETRY
);

CREATE TABLE IF NOT EXISTS findr.map_zone (
	zone_id uuid NOT NULL,
	parking_lot_id uuid NOT NULL,
	zone_name text,
	zone_polygon GEOMETRY
);

CREATE TABLE IF NOT EXISTS findr.map_rows(
	row_id uuid NOT NULL,
	zone_id uuid NOT NULL,
	row_name text,
	row_polygon GEOMETRY
);

CREATE TABLE IF NOT EXISTS findr.map_slot (
	slot_id uuid NOT NULL,
	row_id uuid NOT NULL,
	veh_id uuid NOT NULL,
	slot_name text,
	spacing_no int,
	slot_available bool default true,
	slot_polygon GEOMETRY
);

ALTER TABLE ONLY findr.map
	ADD CONSTRAINT "map_pkey" PRIMARY KEY (map_id),
	ADD CONSTRAINT "dlr_id_fkey" FOREIGN KEY (dlr_id) REFERENCES findr.sys_dealer(dlr_id);
	
ALTER TABLE ONLY findr.map_parking_lot
	ADD CONSTRAINT "parking_lot_id_pkey" PRIMARY KEY (parking_lot_id),
	ADD CONSTRAINT "map_id_fkey" FOREIGN KEY (map_id) REFERENCES findr.map(map_id);

ALTER TABLE ONLY findr.map_zone
	ADD CONSTRAINT "zone_id_pkey" PRIMARY KEY (zone_id),
	ADD CONSTRAINT "parking_lot_id_fkey" FOREIGN KEY (parking_lot_id) REFERENCES findr.map_parking_lot(parking_lot_id);
	

ALTER TABLE ONLY findr.map_rows
	ADD CONSTRAINT "row_id_pkey" PRIMARY KEY (row_id),
	ADD CONSTRAINT "zone_id_fkey" FOREIGN KEY (zone_id) REFERENCES findr.map_zone(zone_id);

ALTER TABLE ONLY findr.map_slot
	ADD CONSTRAINT "slot_id_pkey" PRIMARY KEY (slot_id),
	ADD CONSTRAINT "row_id_fkey" FOREIGN KEY (row_id) REFERENCES findr.map_rows(row_id),
	ADD CONSTRAINT "veh_id_fkey" FOREIGN KEY (veh_id) REFERENCES findr.vehicle(veh_id);

-- +migrate Down

ALTER TABLE ONLY findr.map_slot
    DROP CONSTRAINT IF EXISTS row_id_fkey,
    DROP CONSTRAINT IF EXISTS slot_id_pkey;


ALTER TABLE ONLY findr.map_rows
	DROP CONSTRAINT IF EXISTS zone_id_fkey,
	DROP CONSTRAINT IF EXISTS row_id_pkey;

ALTER TABLE ONLY findr.map_zone
    DROP CONSTRAINT IF EXISTS parking_lot_id_fkey,
    DROP CONSTRAINT IF EXISTS zone_id_pkey;

ALTER TABLE ONLY findr.map_parking_lot
    DROP CONSTRAINT IF EXISTS map_id_fkey,
    DROP CONSTRAINT IF EXISTS parking_lot_id_pkey;

ALTER TABLE ONLY findr.map
    DROP CONSTRAINT IF EXISTS dlr_id_fkey,
    DROP CONSTRAINT IF EXISTS map_pkey;

DROP TABLE IF EXISTS findr.map_slot;

DROP TABLE IF EXISTS findr.map_rows;

DROP TABLE IF EXISTS findr.map_zone;

DROP TABLE IF EXISTS findr.map_parking_lot;

DROP TABLE IF EXISTS findr.map;

DROP EXTENSION IF EXISTS postgis CASCADE;