package db

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (db *CarLocatorDBImpl) AddNewVechicleDB(c *gin.Context, vehicle *models.Vehicle) (*models.VehicleResponse, error) {
	vehResp := models.NewVehicleResponse()
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.vehicle(veh_id,dlr_id,car_status_id,veh_type_id,attach_id,veh_name,
		stock_no,vin,lic_plate,make,model,year,veh_color,transmission_style,created_at) 
		VALUES(:veh_id,:dlr_id,:car_status_id,:veh_type_id,:attach_id,:veh_name,:stock_no,:vin,:lic_plate,:make,:model,
		:year,:veh_color,:transmission_style,:created_at)Returning veh_id,dlr_id,car_status_id,veh_type_id,attach_id,veh_name,
		stock_no,vin,lic_plate,make,model,year,veh_color,transmission_style,created_at`, vehicle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&vehResp)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return vehResp, nil
}

func (db *CarLocatorDBImpl) DealerSignupDB(c *gin.Context, dealer *models.Dealer) (*models.DealerResponse, error) {
	dealerResp := models.NewDealerResponse()
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.sys_dealer(dlr_id,email,dlr_username,email_verified,role_name,created_at) 
		 VALUES(:dlr_id,:email,:dlr_username,:email_verified,:role_name,:created_at)Returning dlr_id,email,dlr_username`, dealer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&dealerResp)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return dealerResp, nil
}

func (db *CarLocatorDBImpl) CreateKeyDB(c *gin.Context, key *models.Key) (*models.Key, error) {
	keyDB := models.Key{}
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.veh_key(key_id,veh_id,Key_name,key_accessible,received_from,received_by,loc_name) 
		VALUES(:key_id,:veh_id,:key_name,:key_accessible,:received_from,:received_by,:loc_name)Returning key_id,veh_id,key_name,key_accessible,received_from,received_by,loc_name`, key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&keyDB)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &keyDB, nil
}

func (db *CarLocatorDBImpl) StaffSignupDB(c *gin.Context, staff *models.Staff) (*models.Staff, error) {
	staffResp := models.NewStaff(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.dlr_staff(staff_id,staff_role_name,dlr_id,staff_email_work,nam_first,nam_last,staff_username,
		staff_phone_cell,staff_email_personal,max_key_staff,max_role_keys,multi_login,dlr_staff_no,staff_info,created_at,profile_image,region,city,state,postal_code,street1_address,street2_address,currency) 
		VALUES(:staff_id,:staff_role_name,:dlr_id,:staff_email_work,:nam_first,:nam_last,:staff_username,:staff_phone_cell,
		:staff_email_personal,:max_key_staff,:max_role_keys,:multi_login,:dlr_staff_no,:staff_info,:created_at,:profile_image,:region,:city,:state,:postal_code,:street1_address,:street2_address,:currency)
		Returning staff_id,staff_role_name,dlr_id,staff_email_work,nam_first,nam_last,staff_username,staff_phone_cell,staff_email_personal,
		max_key_staff,max_role_keys,multi_login,dlr_staff_no,staff_info,profile_image,region,city,state,postal_code,street1_address,street2_address,currency`, staff)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&staffResp)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return staffResp, nil
}

/*
```
func (db *CarLocatorDBImpl) AddVechicleKeyAssignDB(c *gin.Context, assignUnassignData *models.AssignUnassignKey) (*models.AssignUnassignKey, error)
```
AddVechicleKeyAssignDB method need db variable of type *CarLocatorDBImpl struct{} (db *CarLocatorDBImpl) for database connection.
Accept gin.Context and *models.AssignUnassignKey struct{} and write all the required filed in database.The return value is *models.AssignUnassignKey struct{} when all the validation true.
*/
func (db *CarLocatorDBImpl) AddVechicleKeyAssignDB(c *gin.Context, assignUnassignData *models.AssignUnassignKey) (*models.AssignUnassignKey, error) {
	assignResp := models.NewAssignUnassignKey(nil, nil, nil, nil)
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.veh_assign_key(key_assign_id,veh_id,key_id,staff_id,key_assign_status,transfer_id ,transfere_id,assign_cmnt,date_key_assign_start,date_key_assign_stop) 
		VALUES(:key_assign_id,:veh_id,:key_id,:staff_id,:key_assign_status,:transfer_id ,:transfere_id,:assign_cmnt,:date_key_assign_start,:date_key_assign_stop)
		Returning key_assign_id,veh_id,key_id,staff_id,key_assign_status,transfer_id ,transfere_id,assign_cmnt,date_key_assign_start,date_key_assign_stop`, assignUnassignData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&assignResp)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return assignResp, nil

}

func (db *CarLocatorDBImpl) VechicleLocationDB(c *gin.Context, location *models.VehicleLocation) (*models.VehicleLocation, error) {
	locationResp := models.NewVehicleLocation(nil, nil, nil, nil, nil, nil, nil, nil)
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.veh_loc(veh_loc_Id, loc_id, veh_id, space_id, user_in, user_out, date_in_loc, date_out_loc) 
		VALUES(:veh_loc_id,:loc_id,:veh_id,:space_id,:user_in,:user_out,:date_in_loc,
		:date_out_loc)
		Returning veh_loc_id,loc_id,veh_id,space_id,user_in,user_out,date_in_loc,date_out_loc`, location)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&locationResp)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return locationResp, nil
}

func (db *CarLocatorDBImpl) AddNewMapDB(c *gin.Context, mapData *models.Map) (*models.Map, error) {
	mapDB := models.Map{}
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.map(dlr_id,map_id,map_name,map_polygon) VALUES (:dlr_id,:map_id,:map_name,ST_AsGeoJSON(:map_polygon)) Returning map_id,map_name,dlr_id`, mapData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&mapDB)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &mapDB, nil
}

func (db *CarLocatorDBImpl) AddNewMapParkingSlotDB(c *gin.Context, parkingLot *models.ParkingLot) (*models.ParkingLot, error) {
	parkingSlotDB := models.ParkingLot{}
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.map_parking_lot(parking_lot_id,map_id,parking_name,parking_polygon) VALUES (:parking_lot_id,:map_id,:parking_name,ST_AsGeoJSON(:parking_polygon)) Returning parking_lot_id,parking_name,map_id`, parkingLot)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&parkingSlotDB)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &parkingSlotDB, nil
}

func (db *CarLocatorDBImpl) AddNewMapZoneDB(c *gin.Context, zoneData *models.Zone) (*models.Zone, error) {
	zoneDB := models.Zone{}
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.map_zone(zone_id,zone_name,parking_lot_id,zone_polygon) VALUES (:zone_id,:zone_name,:parking_lot_id,ST_AsGeoJSON(:zone_polygon)) Returning zone_id,zone_name,parking_lot_id`, zoneData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&zoneDB)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &zoneDB, nil
}

func (db *CarLocatorDBImpl) AddNewMapRowsDB(c *gin.Context, rowData *models.Row) (*models.Row, error) {
	rowsDB := models.Row{}
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.map_rows(row_id,row_name,zone_id,row_polygon) VALUES (:row_id,:row_name,:zone_id,ST_AsGeoJSON(:row_polygon)) Returning row_id,row_name,zone_id`, rowData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&rowsDB)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &rowsDB, nil
}

func (db *CarLocatorDBImpl) AddNewMapSlotDB(c *gin.Context, mapData *models.Slot) (*models.Slot, error) {
	slotsDB := models.Slot{}
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`INSERT INTO findr.map_slot(slot_id,slot_name,spacing_no,row_id,slot_polygon) VALUES (:slot_id,:slot_name,:spacing_no,:row_id,ST_AsGeoJSON(:slot_polygon)) Returning slot_id,slot_name,spacing_no,row_id,slot_polygon`, mapData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&slotsDB)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &slotsDB, nil
}
