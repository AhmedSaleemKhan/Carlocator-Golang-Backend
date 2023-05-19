-- +migrate Up

CREATE SCHEMA if not exists findr;

CREATE TABLE IF NOT EXISTS findr.sys_dealer (
    dlr_id uuid NOT NULL,
    group_id uuid,
    language_id uuid,
    country_id uuid,
    parent_dlr_id uuid,
	email text,
	dlr_username text NOT NULL,
	dlr_login_pass text NOT NULL,
	created_at bigint NOT NULL,
	updated_at bigint,
	email_verified boolean DEFAULT false,
    region text,
    branch_name text,
    dlr_logo text,
    dlr_name text,
    short_name text,
	street1 text,
	street2 text,
	city text,
	postal_code text,
	telephone text,
	website text,
	sic_code text,
	miles_km bigint,
	dlr_tax_id text,
	currency real,
	dlr_discord text,
	billing_period text,
	invoice_method text,
	payment_method text,
	dlr_cmnt text,
	loc_note text
);

ALTER TABLE ONLY findr.sys_dealer
    ADD CONSTRAINT "sys_dealer_pkey" PRIMARY KEY (dlr_id);

-- +migrate Down

ALTER TABLE ONLY findr.sys_dealer
    DROP CONSTRAINT IF EXISTS sys_dealer_pkey;

DROP TABLE IF EXISTS findr.sys_dealer;

