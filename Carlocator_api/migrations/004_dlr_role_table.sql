-- +migrate Up

CREATE TABLE IF NOT EXISTS findr.dlr_role (
    role_name text NOT NULL,
    dflt_max_key_cnt bigint
);

ALTER TABLE ONLY findr.dlr_role
    ADD CONSTRAINT "role_name_pkey" PRIMARY KEY (role_name);

INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('Accounting',0);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('Admin',0);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('SalesManager',0);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('ITAdministrator',0);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('InventoryManager',3);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('LotPorter',3);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('Receptionist',0);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('SalesRepresentative',3);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('ServiceManager',2);
INSERT INTO findr.dlr_role (role_name,dflt_max_key_cnt) VALUES ('ServiceTechnician',2);


CREATE TABLE IF NOT EXISTS findr.dlr_staff (
    staff_id uuid NOT NULL,
    dlr_id uuid NOT NULL,
    staff_role_name text NOT NULL,
    dlr_staff_no bigint,
    max_key_staff bigint,
    max_role_keys bigint,
    multi_login boolean,
    nam_first text,
    nam_last text,
    staff_username text,
    staff_login_pass text,
    staff_phone_cell text,
    staff_email_work text,
    staff_email_personal text,
    staff_info text,
    created_at bigint
);

ALTER TABLE ONLY findr.dlr_staff
    ADD CONSTRAINT "dlr_staff_pkey" PRIMARY KEY (staff_id);

ALTER TABLE ONLY findr.dlr_staff
    ADD CONSTRAINT "dealer_id_fkey" FOREIGN KEY (dlr_id) REFERENCES findr.sys_dealer(dlr_id);


ALTER TABLE ONLY findr.dlr_staff
    ADD CONSTRAINT "staff_role_name_fkey" FOREIGN KEY (staff_role_name) REFERENCES findr.dlr_role(role_name);

ALTER TABLE ONLY findr.sys_dealer
    ADD COLUMN role_name text;

-- +migrate Down

ALTER TABLE ONLY findr.sys_dealer
    DROP COLUMN role_name;

ALTER TABLE ONLY findr.dlr_staff
    DROP CONSTRAINT IF EXISTS staff_role_name_fkey;

ALTER TABLE ONLY findr.dlr_staff
    DROP CONSTRAINT IF EXISTS dealer_id_fkey;

ALTER TABLE ONLY findr.dlr_staff
    DROP CONSTRAINT IF EXISTS dlr_staff_pkey;

DROP TABLE IF EXISTS findr.dlr_staff;

ALTER TABLE ONLY findr.dlr_role
    DROP CONSTRAINT IF EXISTS role_name_pkey;

DROP TABLE IF EXISTS findr.dlr_role;