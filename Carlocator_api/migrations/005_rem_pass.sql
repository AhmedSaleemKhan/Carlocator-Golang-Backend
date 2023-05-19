-- +migrate Up

ALTER TABLE ONLY findr.sys_dealer
    DROP COLUMN dlr_login_pass;


ALTER TABLE ONLY findr.dlr_staff
    DROP COLUMN staff_login_pass;

-- +migrate Down

ALTER TABLE ONLY findr.sys_dealer
    ADD COLUMN dlr_login_pass text;


ALTER TABLE ONLY findr.dlr_staff
    ADD COLUMN staff_login_pass text;