package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func (db *CarLocatorDBImpl) GetDealerVechiclesDB(c *gin.Context, filter models.GetAllVehiclesFilter) ([]*models.Vehicle, *models.Metadata, error) {
	dealerId := c.MustGet("user-id").(string)
	vehicles := make([]*models.Vehicle, 0)
	countQuery := `SELECT count(*) FROM "findr".vehicle where dlr_id = $1`

	query := `SELECT * FROM "findr".vehicle WHERE dlr_id = $1`

	searchQuery := filter.SearchFilterVehicle
	if searchQuery.Model.Value != nil && len(*searchQuery.Model.Value) > 0 {
		query = fmt.Sprintf(query+` AND model ='%v'`, *searchQuery.Model.Value)
		countQuery = fmt.Sprintf(countQuery+` AND model ='%v'`, *searchQuery.Model.Value)
	}

	if searchQuery.Color.Value != nil && len(*searchQuery.Color.Value) > 0 {
		query = fmt.Sprintf(query+` AND veh_color ='%v'`, *searchQuery.Color.Value)
		countQuery = fmt.Sprintf(countQuery+` AND veh_color ='%v'`, *searchQuery.Color.Value)
	}

	if searchQuery.VehicleAvailable.Value != nil {
		query = fmt.Sprintf(query+` AND veh_available ='%v'`, *searchQuery.VehicleAvailable.Value)
		countQuery = fmt.Sprintf(countQuery+` AND veh_available ='%v'`, *searchQuery.VehicleAvailable.Value)
	}

	if searchQuery.Year.Value != nil && len(*searchQuery.Year.Value) > 0 {
		query = fmt.Sprintf(query+` AND year ='%v'`, *searchQuery.Year.Value)
		countQuery = fmt.Sprintf(countQuery+` AND year ='%v'`, *searchQuery.Year.Value)
	}

	if searchQuery.VehicleName.Value != nil && len(*searchQuery.VehicleName.Value) > 0 {
		query = query + ` AND veh_name LIKE '%` + *searchQuery.VehicleName.Value + `%'`
		countQuery = countQuery + ` AND veh_name LIKE '%` + *searchQuery.VehicleName.Value + `%'`
	}

	query = query + " ORDER BY created_at LIMIT $2 OFFSET $3"

	err := db.conn.Select(&vehicles, query, dealerId, filter.Limit, filter.CalculateOffset())
	if err != nil {
		return nil, nil, err
	}

	var totalRecords int
	err = db.conn.Get(&totalRecords, countQuery, dealerId)
	if err != nil {
		return nil, nil, err
	}

	metadata := models.ComputeMetadata(totalRecords, filter.Page, filter.Limit)
	return vehicles, metadata, nil
}

func (db *CarLocatorDBImpl) GetDealersVehiclesAttributesDB(c *gin.Context) ([]*models.UniqueVehicleAttributesDB, error) {
	dealerId := c.MustGet("user-id").(string)

	distinctValues := make([]*models.UniqueVehicleAttributesDB, 0)

	err := db.conn.Select(&distinctValues, `SELECT
	array_to_json(array(
	  SELECT
		DISTINCT make
	  FROM
		"findr".vehicle
	  WHERE
		dlr_id = $1
	)) as makes,
	array_to_json(array(
	  SELECT
		DISTINCT model
	  FROM
		"findr".vehicle
	  WHERE
		dlr_id = $1
	)) as models,
	array_to_json(array(
	  SELECT
		DISTINCT year
	  FROM
		"findr".vehicle
	  WHERE
		dlr_id = $1
	)) as years 	
	`, dealerId)
	if err != nil {
		return nil, err
	}
	return distinctValues, nil
}

func (db *CarLocatorDBImpl) GetStaffDB(c *gin.Context, staff *models.Staff) (*models.Staff, error) {
	staffData := models.NewStaff(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	rows, err := db.conn.NamedQuery(`SELECT staff_username,staff_email_work FROM findr.dlr_staff WHERE staff_username=:staff_username or staff_email_work=:staff_email_work;`, staff)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&staffData)
		return staffData, err
	}
	return nil, err
}

func (db *CarLocatorDBImpl) GetStaffDealerName(c *gin.Context) (string, error) {
	dealerID := c.MustGet("user-id").(string)

	var dealerName string
	err := db.conn.Get(&dealerName, `SELECT dlr_name FROM findr.sys_dealer WHERE dlr_id=$1`, dealerID)
	if err != nil {
		return dealerName, err
	}
	return dealerName, nil
}

func (db *CarLocatorDBImpl) GetLatestStaffNumber(c *gin.Context) (*string, error) {
	var maxStaffNumber *int64
	userID := c.MustGet("user-id")
	err := db.conn.Get(&maxStaffNumber, `SELECT count(*) FROM "findr".dlr_staff where dlr_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	staffNUmberString := strconv.Itoa(int(*maxStaffNumber))
	return &staffNUmberString, nil
}

func (db *CarLocatorDBImpl) GetDealerDB(c *gin.Context, dealer *models.Dealer) (*models.Dealer, error) {
	dealerData := models.NewDealer(nil, nil, nil)
	rows, err := db.conn.NamedQuery(`SELECT dlr_username,email FROM findr.sys_dealer WHERE dlr_username=:dlr_username or email=:email;`, dealer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&dealerData)
		return dealerData, err
	}
	return nil, err
}

func (db *CarLocatorDBImpl) GetVehicleIDForKeyDB(c *gin.Context, keyID string) (*string, error) {
	var vehicleID *string
	err := db.conn.Get(&vehicleID, `SELECT veh_id from "findr".veh_key where key_id=$1`, keyID)
	if err != nil {
		return nil, err
	}
	return vehicleID, nil
}

func (db *CarLocatorDBImpl) GetVehicleDealerIDForKeyDB(c *gin.Context, vehID string) (*string, error) {
	var vehicleDealerID string
	err := db.conn.Get(&vehicleDealerID, `SELECT dlr_id from "findr".vehicle where veh_id=$1`, vehID)
	if err != nil {
		return nil, err
	}
	return &vehicleDealerID, nil
}

func (db *CarLocatorDBImpl) GetDealerStaffID(c *gin.Context, staffId string) (*string, error) {
	var staffDealerID string
	err := db.conn.Get(&staffDealerID, `SELECT dlr_id from findr.dlr_staff where staff_id=$1`, staffId)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return &staffDealerID, nil
}

func (db *CarLocatorDBImpl) GetDealerIDForKeyDB(c *gin.Context, staffID string) (*string, error) {
	var staffDealerID string
	err := db.conn.Get(&staffDealerID, `SELECT dlr_id from "findr".dlr_staff where staff_id=$1`, staffID)
	if err != nil {
		return nil, err
	}
	return &staffDealerID, nil
}
func (db *CarLocatorDBImpl) GetkeyStatusDB(c *gin.Context, keyID string) (*bool, error) {
	var keyStatus bool
	err := db.conn.Get(&keyStatus, `SELECT key_available from findr.veh_key where key_id=$1  `, keyID)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return &keyStatus, nil
}

func (db *CarLocatorDBImpl) GetVehicleKeysDB(c *gin.Context, vehicleID string) ([]*models.Key, error) {
	keys := make([]*models.Key, 0)

	err := db.conn.Select(&keys, `SELECT * FROM "findr".veh_key WHERE veh_id = $1 `, vehicleID)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (db *CarLocatorDBImpl) GetStaffByIDDB(c *gin.Context, staff *models.Staff) (*models.Staff, error) {
	staffData := models.NewStaff(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	rows, err := db.conn.NamedQuery(`SELECT * FROM findr.dlr_staff WHERE staff_id=:staff_id;`, staff)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&staffData)
		return staffData, err
	}
	return nil, err
}

func (db *CarLocatorDBImpl) GetVehicleDealerIDByVehIdDB(c *gin.Context, vehID string) (*string, error) {
	var vehicleDealerID string
	err := db.conn.Get(&vehicleDealerID, `SELECT dlr_id from findr.vehicle where veh_id=$1`, vehID)
	if err != nil {
		return nil, err
	}
	return &vehicleDealerID, nil
}

func (db *CarLocatorDBImpl) GetStaffMembersDB(c *gin.Context, userID string, filter models.GetStaffMembersFilter) ([]*models.Staff, *models.Metadata, error) {
	staffMembersDb := make([]*models.Staff, 0)
	err := db.conn.Select(&staffMembersDb, `SELECT * FROM findr.dlr_staff WHERE dlr_id=$1 ORDER BY created_at LIMIT $2 OFFSET $3 `, userID, filter.Limit, filter.CountOffset())
	if err != nil {
		return nil, nil, err
	}

	var totalRecords int
	err = db.conn.Get(&totalRecords, `SELECT count(*) FROM "findr".dlr_staff where dlr_id = $1`, userID)
	if err != nil {
		return staffMembersDb, nil, nil
	}
	metadata := models.ComputeMetadata(totalRecords, filter.Page, filter.Limit)
	return staffMembersDb, metadata, nil
}

func (db *CarLocatorDBImpl) GetDealerAllInfoDB(c *gin.Context, userID string) ([]*models.Dealer, error) {

	dealerInfoDb := make([]*models.Dealer, 0)
	err := db.conn.Select(&dealerInfoDb, `SELECT * FROM findr.sys_dealer WHERE dlr_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	return dealerInfoDb, err
}

func (db *CarLocatorDBImpl) GetMaximumRoleKeysDB(c *gin.Context, staffRole string) (*int64, error) {
	var maxRoleKeys int64
	err := db.conn.Get(&maxRoleKeys, `SELECT dflt_max_key_cnt from findr.dlr_role where role_name=$1`, staffRole)
	if err != nil {
		return &maxRoleKeys, err
	}
	return &maxRoleKeys, nil
}

func (db CarLocatorDBImpl) GetStaffRolekeysDB(c *gin.Context, staffID string) (*int64, error) {
	var staffRole int64
	err := db.conn.Get(&staffRole, `SELECT max_role_keys from findr.dlr_staff where staff_id=$1`, staffID)
	if err != nil {
		return &staffRole, err
	}
	return &staffRole, nil
}

func (db *CarLocatorDBImpl) GetMaxkeystaffDB(c *gin.Context, staffID string) (*int64, error) {
	var staffMaxKeys int64
	err := db.conn.Get(&staffMaxKeys, `SELECT max_key_staff from findr.dlr_staff where staff_id=$1`, staffID)
	if err != nil {
		return &staffMaxKeys, err
	}
	return &staffMaxKeys, nil
}

func (db *CarLocatorDBImpl) GetUserAssignedKeysDB(c *gin.Context, dealerID string, filter models.GetAssignedKeyFilter) ([]*models.AssignedKeys, *models.Metadata, error) {
	assignedKeysDb := make([]*models.AssignedKeys, 0)
	err := db.conn.Select(&assignedKeysDb, `SELECT
    Veh.veh_name,
    Veh.stock_no,
    Veh.make,
    Veh.model,
    Veh.year,
    Veh.veh_color,
    Veh.transmission_style,
    VSK.veh_id,
    VSK.key_id,
    VSK.staff_id,
    VSK.key_assign_status,
    VSK.transfer_id,
    DS.staff_role_name,
    DS.nam_first,
    DS.nam_last,
    DS.staff_username
FROM
    findr.vehicle as Veh
    JOIN findr.veh_assign_key as VSK ON VSK.veh_id = Veh.veh_id
    and VSK.key_assign_status = true
    JOIN findr.dlr_staff as DS ON DS.staff_id = VSK.staff_id
	where Veh.dlr_id = $1 
ORDER BY
    veh_id
LIMIT
    $2 OFFSET $3 `, dealerID, filter.Limit, filter.CalculatedOffset())
	if err != nil {
		return nil, nil, err
	}
	var totalRecords int
	err = db.conn.Get(&totalRecords, `SELECT count(*) FROM 
	findr.vehicle         
	JOIN findr.veh_assign_key ON findr.veh_assign_key.veh_id = findr.vehicle.veh_id 
	and findr.veh_assign_key.key_assign_status = true
	JOIN findr.dlr_staff ON findr.dlr_staff.dlr_id = findr.vehicle.dlr_id`)
	if err != nil {
		return assignedKeysDb, nil, err
	}
	metadata := models.ComputeMetadata(totalRecords, filter.Page, filter.Limit)
	return assignedKeysDb, metadata, nil
}

func (db *CarLocatorDBImpl) GetAllDealersDB(c *gin.Context) ([]*models.GetAllDealers, error) {
	getalldealerDb := make([]*models.GetAllDealers, 0)
	err := db.conn.Select(&getalldealerDb, `SELECT 
	findr.sys_dealer.dlr_id,
	findr.sys_dealer.dlr_username,
	findr.sys_dealer.dlr_name
	FROM
	findr.sys_dealer`)
	if err != nil {
		return nil, err
	}
	return getalldealerDb, err
}

func (db *CarLocatorDBImpl) GetNoOfVehicles(c *gin.Context) (*int64, error) {
	var noOfVehicles int64
	userID := c.MustGet("user-id")
	err := db.conn.Get(&noOfVehicles, `SELECT count(*) FROM "findr".vehicle where dlr_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	return &noOfVehicles, nil
}

//this function is used to get the map data belongs to specific dealer
func (db *CarLocatorDBImpl) GetDealerMapDB(c *gin.Context, dealerID string) ([]*models.Map, error) {
	dealerMapDB := make([]*models.Map, 0)
	err := db.conn.Select(&dealerMapDB, `SELECT 
						findr.map.map_id,
						findr.map.map_name,
						findr.map.dlr_id,
						ST_AsGeoJSON(findr.map.map_polygon)
						FROM findr.map WHERE dlr_id = $1 `, dealerID)

	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return dealerMapDB, nil
}

//function used to get the parking lots according to specific map id
func (db *CarLocatorDBImpl) GetDealerParkingLot(c *gin.Context, mapID *string) ([]*models.ParkingLot, error) {
	parkingLotsDB := make([]*models.ParkingLot, 0)
	err := db.conn.Select(&parkingLotsDB, `SELECT 
						findr.map_parking_lot.parking_lot_id,
						findr.map_parking_lot.parking_name,
						findr.map_parking_lot.map_id,
						ST_AsGeoJSON(findr.map_parking_lot.parking_polygon)
						FROM findr.map_parking_lot WHERE map_id =$1 `, mapID)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return parkingLotsDB, nil
}

//function used to get the zones by specific parking-lot-id
func (db *CarLocatorDBImpl) GetDealerMapZone(c *gin.Context, parkingLotID *string) ([]*models.Zone, error) {
	zones := make([]*models.Zone, 0)
	err := db.conn.Select(&zones, `SELECT 
	findr.map_zone.zone_id,
	findr.map_zone.zone_name,
	findr.map_zone.parking_lot_id,
	ST_AsGeoJSON(zone_polygon)
	FROM findr.map_zone WHERE parking_lot_id = $1`, parkingLotID)

	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return zones, nil
}

//function used to get the rows by specific zone-id
func (db *CarLocatorDBImpl) GetDealerMapRow(c *gin.Context, zoneID *string) ([]*models.Row, error) {
	rows := make([]*models.Row, 0)
	err := db.conn.Select(&rows, `SELECT 
	findr.map_rows.row_id,
	findr.map_rows.zone_id,
	findr.map_rows.row_name,
	ST_AsGeoJSON(row_polygon)
	FROM findr.map_rows WHERE zone_id = $1`, zoneID)

	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return rows, nil
}

//this function is suppsoed to return the slots belong to that specific row-id
func (db *CarLocatorDBImpl) GetDealerMapSlot(c *gin.Context, rowID *string) ([]*models.Slot, error) {
	slots := make([]*models.Slot, 0)
	err := db.conn.Select(&slots, `SELECT
    MS.slot_id,
    MS.row_id,
    MS.slot_name,
    MS.spacing_no,
    MS.slot_available,
    MS.veh_id,
    Veh.veh_name,
    Veh.veh_color,
    Veh.lic_plate,
    Veh.year,
    Veh.make,
    Veh.transmission_style,
    ST_AsGeoJSON(slot_polygon)
FROM
    findr.map_slot AS MS
    left JOIN findr.vehicle AS Veh ON Veh.veh_id = MS.veh_id
WHERE
    row_id = $1
ORDER BY
    MS.slot_id ASC`, rowID)
	//we need to return the vheh_if as well in order to check
	//either vehicle is checked or not...
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return slots, nil
}

func (db *CarLocatorDBImpl) GetDealerStaffIDForVehDB(c *gin.Context, staffID string) ([]*models.StaffAssignedVehicleKeys, error) {
	DealerStaffID := make([]*models.StaffAssignedVehicleKeys, 0)
	err := db.conn.Select(&DealerStaffID, `SELECT
    Veh.veh_id,
    Veh.veh_name,
    Veh.dlr_id,
    Veh.attach_id,
    Veh.stock_no,
    Veh.vin,
    Veh.lic_plate,
    Veh.make,
    Veh.model,
    Veh.year,
    Veh.car_status_id,
    Veh.veh_color,
    Veh.transmission_style,
    Veh.created_at,
    VSK.key_id,
    VSK.key_assign_status,
    VSK.staff_id
FROM
    findr.vehicle as Veh
    JOIN findr.veh_assign_key as VSK ON VSK.veh_id = Veh.veh_id
    and VSK.key_assign_status = true
    and staff_id = $1`, staffID)
	if err != nil {
		return nil, err
	}
	return DealerStaffID, nil
}

func (db *CarLocatorDBImpl) GetDealerVinNumbers(c *gin.Context, dealerID string) ([]string, error) {
	vinNumbers := make([]string, 0)
	err := db.conn.Select(&vinNumbers, `SELECT vin FROM "findr".vehicle WHERE dlr_id=$1 `, dealerID)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return vinNumbers, nil
}

func (db *CarLocatorDBImpl) GetAssignedStaffId(c *gin.Context, keyID string) *string {
	var staffId *string
	db.conn.Get(&staffId, `SELECT staff_id from findr.veh_assign_key WHERE key_assign_status= true AND key_id=$1 `, keyID)
	return staffId
}

func (db *CarLocatorDBImpl) GetStaffByKeyId(c *gin.Context, keyID string) (*models.Staff, error) {
	var Staff models.Staff
	staffId := db.GetAssignedStaffId(c, keyID)
	if staffId != nil {
		err := db.conn.Get(&Staff, `SELECT * from findr.dlr_staff where staff_id=$1`, staffId)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
	}
	return &Staff, nil
}

//this function is suppossed to get the dealer id where slot_id = body.slot_id
//to validate the map belongs to the same dealer as logged in (token)
func (db *CarLocatorDBImpl) GetDealerIDFromSlotID(c *gin.Context, slotID string) (*models.CheckInOutVehicleMap, error) {
	mapData := models.CheckInOutVehicleMap{}
	err := db.conn.Get(&mapData, `Select
    Map.dlr_id,
	MS.slot_available,
	MS.veh_id
FROM
    findr.map AS Map
    join findr.map_parking_lot AS MPL ON MPL.map_id = Map.map_id
    join findr.map_zone AS MZ ON MZ.parking_lot_id = MPL.parking_lot_id
    join findr.map_rows AS MR ON MR.zone_id = MZ.zone_id
    join findr.map_slot AS MS ON MS.row_id = MR.row_id
WHERE
    MS.slot_id = $1 `, slotID)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	//will return the dealer id which should be compared with token once..
	return &mapData, nil
}

func (db *CarLocatorDBImpl) GetStaffFromAssignKeyByVehIDDB(c *gin.Context, vehID string) (*string, error) {
	var staffId *string
	err := db.conn.Get(&staffId, `SELECT staff_id from findr.veh_assign_key WHERE veh_id=$1 `, vehID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return staffId, nil
}

func (db *CarLocatorDBImpl) CheckVehExistsInSlotDB(c *gin.Context, vehicleID string) (*models.AddingVehicleToMap, error) {
	slotData := models.AddingVehicleToMap{}
	err := db.conn.Get(&slotData, `SELECT
	slot_id,
	veh_id	
FROM
	findr.map_slot
WHERE
	veh_id=$1 `, vehicleID)
	if err != nil && err != sql.ErrNoRows {
		glg.Error(err)
		return nil, err
	}
	return &slotData, nil
}

func (db *CarLocatorDBImpl) GetVehAssignStatus(c *gin.Context, vehicleID string) (*bool, error) {
	var vehAssignedStatus bool
	staffID := c.MustGet("user-id").(string)
	err := db.conn.Get(&vehAssignedStatus, `SELECT 
	key_assign_status	
FROM
	findr.veh_assign_key
WHERE
	veh_id=$1 and staff_id=$2`, vehicleID, staffID)
	if err != nil && err != sql.ErrNoRows {
		glg.Error(err)
		return nil, err
	}
	return &vehAssignedStatus, nil
}

func (db *CarLocatorDBImpl) GetSlotAvailableStatus(c *gin.Context, slotID string) (*bool, error) {
	var vehAvailableStatus bool
	err := db.conn.Get(&vehAvailableStatus, `SELECT 
	slot_available	
FROM
	findr.map_slot
WHERE
	slot_id=$1`, slotID)
	if err != nil && err != sql.ErrNoRows {
		glg.Error(err)
		return nil, err
	}
	return &vehAvailableStatus, nil
}
