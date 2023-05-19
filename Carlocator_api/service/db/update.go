package db

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func (db *CarLocatorDBImpl) DealerVerifyOTPDB(c *gin.Context, dealer *models.Dealer) (*models.VerifyOTPResponse, error) {
	resp := models.NewVerifyOTPResponse()
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`UPDATE findr.sys_dealer set email_verified=:email_verified WHERE email=:email Returning email`, dealer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&resp)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
```
func (db *CarLocatorDBImpl) AddVechicleKeyUnassignDB(c *gin.Context, unAssignData *models.AssignUnassignKey) (*models.AssignUnassignKey, error)
```
AddVechicleKeyUnassignDB method need db variable of type *CarLocatorDBImpl struct{} (db *CarLocatorDBImpl) for database connection.
Accept gin.Context and *models.AssignUnassignKey struct{} and write all the required filed in database.The return value is *models.AssignUnassignKey struct{} when all the validation true.
*/

func (db *CarLocatorDBImpl) AddVechicleKeyUnassignDB(c *gin.Context, unAssignData *models.AssignUnassignKey) (*models.AssignUnassignKey, error) {
	tx := db.conn.MustBegin()
	rows, err := tx.NamedExec(`Update findr.veh_assign_key SET key_assign_status = false, date_key_assign_stop=:date_key_assign_stop where key_id=:key_id and key_assign_status=true 
	Returning key_assign_id,veh_id,key_id,staff_id,key_assign_status,transfer_id ,transfere_id,assign_cmnt,date_key_assign_start,date_key_assign_stop;`, unAssignData)
	if err != nil {
		return nil, err
	}
	// Checking rows affected
	affectedResult, err := rows.RowsAffected()
	if err != nil {
		return nil, err
	}
	// If no row affected key is already unassign
	if affectedResult < 1 {
		return nil, fmt.Errorf("invalid operation already unassign")
	} else {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
		// Returning reasponse
		var keyAssignID, vehicleID, keyID, staffID, transferID, transfereID, assignComment *string
		var keyAssignStatus *bool
		var dataKeyAssignStart, dataKeyAssignStop *int64
		tx := db.conn.MustBegin()
		err := tx.QueryRow(`select * from findr.veh_assign_key where date_key_assign_stop=$1`, unAssignData.DataKeyAssignStop).Scan(&keyAssignID, &vehicleID, &keyID, &staffID, &keyAssignStatus, &transferID, &transfereID, &assignComment, &dataKeyAssignStart, &dataKeyAssignStop)
		if err != nil {
			return nil, err
		}
		*unAssignData.KeyAssignID = *keyAssignID
		*unAssignData.VehicleID = *vehicleID
		*unAssignData.KeyID = *keyID
		*unAssignData.StaffID = *staffID
		*unAssignData.TransferID = *transferID
		*unAssignData.TransfereID = *transfereID
		*unAssignData.AssignComment = *assignComment
		*unAssignData.KeyAssignStatus = *keyAssignStatus
		*unAssignData.DataKeyAssignStart = *dataKeyAssignStart
		*unAssignData.DataKeyAssignStop = *dataKeyAssignStop
		return unAssignData, nil

	}

}

func (db *CarLocatorDBImpl) UpdateVehicleKeyStatusDB(c *gin.Context, Status bool, keyID string, vehID string) error {
	tx := db.conn.MustBegin()
	_, err := tx.Exec(`UPDATE findr.veh_key SET key_available =$2 WHERE key_id= $1`, keyID, Status)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE findr.vehicle SET veh_available =$2  WHERE veh_id=$1`, vehID, Status)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("unable to update vehicle available status")
	}

	return nil
}

func (db *CarLocatorDBImpl) UpdateDealershipDB(c *gin.Context, updatedDealer *models.Dealer, dealerID string) (*models.Dealer, error) {
	tx := db.conn.MustBegin()

	rows, err := tx.NamedQuery(`UPDATE 
	findr.sys_dealer 
	SET
	created_at = :created_at,
	region = :region, 
	branch_name = :branch_name, 
	dlr_logo = :dlr_logo, 
	dlr_name = :dlr_name, 
	short_name = :short_name, 
	street1_address = :street1_address, 
	street2_address = :street2_address, 
	city = :city, 
	telephone = :telephone, 
	postal_code = :postal_code, 
	website = :website, 
	sic_code = :sic_code, 
	miles_km = :miles_km,
	currency = :currency,
	dlr_discord = :dlr_discord, 
	billing_period = :billing_period, 
	invoice_method = :invoice_method, 
	payment_method = :payment_method, 
	dlr_cmnt = :dlr_cmnt, 
	loc_note = :loc_note,
	state = :state,
	dealership_name = :dealership_name,
	dealership_logo = :dealership_logo
	WHERE dlr_id = :dlr_id
	RETURNING 
	created_at,
	region,
	branch_name, 
	dlr_logo,
	dlr_name,
	short_name,
	street1_address, 
	street2_address,
	city, 
	telephone,
	postal_code,
	website,
	sic_code,
	miles_km,
	currency,
	dlr_discord,
	billing_period,
	invoice_method,
	payment_method,
	dlr_cmnt,
	loc_note,
	state,
	dealership_logo,
	dealership_name	
  `, updatedDealer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	updatedDealerDB := models.NewUpdateDealer(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	for rows.Next() {
		err := rows.StructScan(&updatedDealerDB)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return updatedDealerDB, nil
}

func (db *CarLocatorDBImpl) UpdateMaxStaffKeyDB(c *gin.Context, maxkey int64, staffID string) error {
	tx := db.conn.MustBegin()
	_, err := tx.Exec(`UPDATE findr.dlr_staff SET max_key_staff=$2 where staff_id=$1`, staffID, maxkey)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db *CarLocatorDBImpl) UpdatedVehicleDB(c *gin.Context, VehicleID string, vehicle *models.Vehicle) (*models.Vehicle, error) {
	tx := db.conn.MustBegin()
	*vehicle.VehicleID = VehicleID
	rows, err := tx.NamedQuery(`UPDATE 
	findr.vehicle 
	SET	
	veh_name = :veh_name,
	stock_no =  :stock_no, 
	vin = :vin, 
	lic_plate = :lic_plate, 
	model = :model, 
	year = :year,
	make = :make,
	veh_color = :veh_color, 
	transmission_style = :transmission_style, 
	map_color = :map_color, 
	note = :note, 
	veh_available = :veh_available 
  	WHERE veh_id = :veh_id
	RETURNING 
	veh_id,
	veh_name, 
	stock_no, 
	vin, 
	make,
	lic_plate, 
	model, 
	year,
	veh_color, 
	transmission_style, 
	map_color, 
	note, 
	veh_available
  `, vehicle)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	UpdatedVehicle := models.NewUpdatedVehicle(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	for rows.Next() {
		err := rows.StructScan(&UpdatedVehicle)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return UpdatedVehicle, nil
}

func (db *CarLocatorDBImpl) UpdatedKeyDB(c *gin.Context, keyID string, key *models.Key) (*models.Key, error) {

	query := `UPDATE findr.veh_key 
		SET 
		loc_name = :loc_name, 
		key_name = :key_name, 
		key_accessible = :key_accessible, 
		received_from = :received_from, 
		received_by = :received_by, 
		key_available = :key_available 
		WHERE 
		key_id = :key_id RETURNING veh_id, 
		key_id, 
		loc_name, 
		key_name, 
		key_accessible, 
		received_from, 
		received_by, 
		key_available
	`
	tx := db.conn.MustBegin()
	*key.KeyID = keyID
	rows, err := tx.NamedQuery(query, key)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Updatedkey := models.NewUpdatedKey(nil, nil, nil, nil, nil, nil, nil)
	for rows.Next() {
		err := rows.StructScan(&Updatedkey)
		if err != nil {
			return nil, err
		}

	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return Updatedkey, nil
}

//the function is called if the vehicle wants to check in ...
func (db *CarLocatorDBImpl) CheckInVehicleDB(c *gin.Context, checkInOutvehicle *models.CheckInOutVehicleMap) (*models.CheckInOutVehicleMap, error) {
	tx := db.conn.MustBegin()
	//set slot status = false because slot is occupied by the specific veh-id....
	rows, err := tx.NamedQuery(`UPDATE 
	findr.map_slot 
  	SET 
	veh_id = :veh_id,
	slot_available = false 
  	WHERE 
	slot_id = :slot_id 
	RETURNING 
	veh_id, 
	slot_id,  
	slot_available
  `, checkInOutvehicle)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&checkInOutvehicle)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return checkInOutvehicle, nil
}

//this function is supposed to checkout the vehicle from table findr.map_slot for specific slot-id
func (db *CarLocatorDBImpl) CheckOutVehicleDB(c *gin.Context, checkInOutvehicle *models.CheckInOutVehicleMap) (*models.CheckInOutVehicleMap, error) {
	tx := db.conn.MustBegin()
	//set the slot - available status =true cooz the vehicle is checked out and availabe status become true
	rows, err := tx.NamedQuery(`UPDATE 
	findr.map_slot 
  	SET 
	veh_id = null,
	slot_available = true 
  	WHERE 
	slot_id = :slot_id 
	RETURNING 
	slot_id,
	veh_id,
	slot_available
  `, checkInOutvehicle)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&checkInOutvehicle)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return checkInOutvehicle, nil
}

func (db *CarLocatorDBImpl) UpdateVechicleKeyAssignDB(c *gin.Context, assignData *models.AssignUnassignKey) (*models.AssignUnassignKey, error) {
	tx := db.conn.MustBegin()
	rows, err := tx.NamedQuery(`Update
    findr.veh_assign_key
SET
    key_assign_status = true,
    date_key_assign_start = :date_key_assign_start,
	date_key_assign_stop=:date_key_assign_stop
where
    key_id = :key_id Returning key_assign_id,
    veh_id,
    key_id,
    staff_id,
    key_assign_status,
    transfer_id,
    transfere_id,
    assign_cmnt,
    date_key_assign_start,
    date_key_assign_stop`, assignData)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&assignData)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return assignData, nil
}

//this func is supposed to add the new vehicle to specific slot in the map
func (db *CarLocatorDBImpl) AddNewVehicleToMapDB(c *gin.Context, vehicleForMap *models.AddingVehicleToMap) (*models.AddingVehicleToMap, error) {
	tx := db.conn.MustBegin()
	//set slot status = false because slot is occupied by the specific vehicle....
	rows, err := tx.NamedQuery(`UPDATE 
	findr.map_slot 
  	SET 
	veh_id = :veh_id,
	slot_available = false 
  	WHERE 
	slot_id = :slot_id 
	RETURNING 
	veh_id, 
	slot_id,  
	slot_available
  `, vehicleForMap)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&vehicleForMap)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return vehicleForMap, nil
}
