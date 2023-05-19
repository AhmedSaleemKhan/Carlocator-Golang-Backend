package models

type Vehicle struct {
	VehicleID         *string `json:"veh_id,omitempty" db:"veh_id"`
	DealerID          *string `json:"dlr_id,omitempty" db:"dlr_id"`
	CarStatusID       *string `json:"car_status_id,omitempty" db:"car_status_id"`
	VechicleTypeID    *string `json:"veh_type_id,omitempty" db:"veh_type_id"`
	AttachId          *string `json:"attach_id,omitempty" db:"attach_id"`
	VehicleName       *string `json:"veh_name,omitempty" db:"veh_name"`
	StockNumber       *int64  `json:"stock_no,omitempty" db:"stock_no"`
	VinNumber         *string `json:"vin,omitempty" db:"vin"`
	LicPlate          *string `json:"lic_plate,omitempty" db:"lic_plate"`
	BrandMake         *string `json:"make,omitempty" db:"make"`
	Model             *string `json:"model,omitempty" db:"model"`
	ManfactureYear    *string `json:"year,omitempty" db:"year"`
	VechicleColor     *string `json:"veh_color,omitempty"  db:"veh_color"`
	TransmissionStyle *string `json:"transmission_style,omitempty" db:"transmission_style"`
	CreatedAt         *int64  `json:"created_at,omitempty" db:"created_at"`
	MapColour         *string `json:"map_color,omitempty"  db:"map_color"`
	Note              *string `json:"note,omitempty" db:"note"`
	VehAvailable      *bool   `json:"veh_available,omitempty" db:"veh_available"`
	Keys              []*Key  `json:"keys,omitempty" db:"keys"`
}

func NewVehicle(vehicalName, vinNumber, licPlate, brandMake, model, manfactureYear, vechicleColor, TransmissionStyle *string, vehicleAvailable *bool) *Vehicle {
	return &Vehicle{
		VehicleID:         new(string),
		DealerID:          new(string),
		CarStatusID:       new(string),
		VechicleTypeID:    new(string),
		AttachId:          new(string),
		VehicleName:       vehicalName,
		StockNumber:       new(int64),
		VinNumber:         vinNumber,
		LicPlate:          licPlate,
		BrandMake:         brandMake,
		Model:             model,
		ManfactureYear:    manfactureYear,
		VechicleColor:     vechicleColor,
		TransmissionStyle: TransmissionStyle,
		CreatedAt:         new(int64),
		MapColour:         new(string),
		Note:              new(string),
		VehAvailable:      vehicleAvailable,
	}
}

type Dealer struct {
	DealerId       *string `json:"dlr_id,omitempty" db:"dlr_id"`
	Email          *string `json:"email,omitempty" db:"email"`
	Username       *string `json:"username,omitempty" db:"dlr_username"`
	Password       *string `json:"password,omitempty" db:"dlr_login_pass"`
	EmailVerified  *bool   `json:"email_verified,omitempty" db:"email_verified"`
	RoleName       *string `json:"role_name,omitempty" db:"role_name"`
	CreatedAt      *int64  `json:"created_at,omitempty" db:"created_at"`
	GroupId        *string `json:"group_id,omitempty" db:"group_id"`
	LanguageId     *string `json:"language_id,omitempty" db:"language_id"`
	CountryId      *string `json:"country_id,omitempty" db:"country_id"`
	ParentDealerId *string `json:"parent_dlr_id,omitempty" db:"parent_dlr_id"`
	UpdatedAt      *int64  `json:"updated_at,omitempty" db:"updated_at"`
	Region         *string `json:"region,omitempty" db:"region"`
	BranchName     *string `json:"branch_name,omitempty" db:"branch_name"`
	DealerLogo     *string `json:"dlr_logo,omitempty" db:"dlr_logo"`
	DealerName     *string `json:"dlr_name,omitempty" db:"dlr_name"`
	ShortName      *string `json:"short_name,omitempty" db:"short_name"`
	Street1Address *string `json:"street1_address,omitempty" db:"street1_address"`
	Street2Address *string `json:"street2_address,omitempty" db:"street2_address"`
	City           *string `json:"city,omitempty" db:"city"`
	Telephone      *int64  `json:"telephone,omitempty" db:"telephone"`
	PostalCode     *int64  `json:"postal_code,omitempty" db:"postal_code"`
	Website        *string `json:"website,omitempty" db:"website"`
	Sic_code       *int64  `json:"sic_code,omitempty" db:"sic_code"`
	MilesKm        *int64  `json:"miles_km,omitempty" db:"miles_km"`
	DealerTaxId    *string `json:"dlr_tax_id,omitempty" db:"dlr_tax_id"`
	Currency       *string `json:"currency,omitempty" db:"currency"`
	DealerDiscord  *string `json:"dlr_discord,omitempty" db:"dlr_discord"`
	BilingPeriod   *string `json:"billing_period,omitempty" db:"billing_period"`
	InvoiceMethod  *string `json:"invoice_method,omitempty" db:"invoice_method"`
	PaymentMethod  *string `json:"payment_method,omitempty" db:"payment_method"`
	DealerComment  *string `json:"dlr_cmnt,omitempty" db:"dlr_cmnt"`
	LocNote        *string `json:"loc_note,omitempty" db:"loc_note"`
	State          *string `json:"state,omitempty" db:"state"`
	DealershipName *string `json:"dealership_name,omitempty" db:"dealership_name"`
	DealershipLogo *string `json:"dealership_logo,omitempty" db:"dealership_logo"`
}

func NewDealer(email, username, password *string) *Dealer {
	return &Dealer{
		DealerId:      new(string),
		Email:         email,
		Username:      username,
		Password:      password,
		EmailVerified: new(bool),
		RoleName:      new(string),
		CreatedAt:     new(int64),
	}
}

func NewUpdateDealer(currency, region, city, State, street1Address, street2Address, branchName, dealerLogo, dealerName, shortName, website, deealerDiscord, billingPeriod, invoiceMethod, paymentMethod, dealerComment, locnote, dealershipName, dealershipLogo *string, postalCode *int64, milesKm *int64, sicCode *int64) *Dealer {
	return &Dealer{
		DealerId:       new(string),
		GroupId:        new(string),
		Region:         region,
		City:           city,
		State:          State,
		PostalCode:     postalCode,
		Street1Address: street1Address,
		Street2Address: street2Address,
		UpdatedAt:      new(int64),
		Currency:       currency,
		LanguageId:     new(string),
		CountryId:      new(string),
		ParentDealerId: new(string),
		BranchName:     branchName,
		DealerLogo:     dealerLogo,
		DealerName:     dealerName,
		ShortName:      shortName,
		Website:        website,
		DealerTaxId:    new(string),
		DealerDiscord:  deealerDiscord,
		BilingPeriod:   billingPeriod,
		InvoiceMethod:  invoiceMethod,
		PaymentMethod:  paymentMethod,
		DealerComment:  dealerComment,
		LocNote:        locnote,
		MilesKm:        milesKm,
		Sic_code:       sicCode,
		DealershipName: dealershipName,
		DealershipLogo: dealershipLogo,
	}
}

type Key struct {
	KeyID         *string `json:"key_id,omitempty" db:"key_id"`
	VehID         *string `json:"veh_id,omitempty" db:"veh_id"`
	KeyName       *string `json:"key_name,omitempty" db:"key_name"`
	KeyAccessible *string `json:"key_accessible,omitempty" db:"key_accessible"`
	ReceivedFrom  *string `json:"received_from,omitempty" db:"received_from"`
	ReceivedBy    *string `json:"received_by,omitempty" db:"received_by"`
	KeyAvailable  *bool   `json:"key_available,omitempty" db:"key_available"`
	LocName       *string `json:"loc_name,omitempty" db:"loc_name"`
	//updated the models because we need nested staff in get keys....
	Staff *Staff `json:"staff,omitempty" db:"staff"`
}

func NewKey(vehId, keyName, keyAccessible, receivedFrom, receivedBy, locName *string) *Key {
	return &Key{
		KeyID:         new(string),
		VehID:         vehId,
		KeyName:       keyName,
		KeyAccessible: keyAccessible,
		ReceivedFrom:  receivedFrom,
		ReceivedBy:    receivedBy,
		KeyAvailable:  new(bool),
		LocName:       locName,
	}
}

type UniqueVehicleAttributesDB struct {
	BrandMakes      *string `json:"makes" db:"makes"`
	Models          *string `json:"models" db:"models"`
	ManfactureYears *string `json:"years" db:"years" `
}

type Staff struct {
	StaffId         *string `json:"staff_id,omitempty" db:"staff_id"`
	StaffRoleName   *string `json:"staff_role_name,omitempty" db:"staff_role_name"`
	DealerID        *string `json:"dlr_id,omitempty" db:"dlr_id"`
	EmailWork       *string `json:"email_work,omitempty" db:"staff_email_work"`
	FirstName       *string `json:"first_name,omitempty" db:"nam_first"`
	LastName        *string `json:"last_name,omitempty" db:"nam_last"`
	Username        *string `json:"username,omitempty" db:"staff_username"`
	Password        *string `json:"password,omitempty" db:"staff_login_pass"`
	PhoneNumber     *string `json:"phone_number,omitempty" db:"staff_phone_cell"`
	EmailPersonal   *string `json:"email_personal,omitempty" db:"staff_email_personal"`
	MaximumKeyStaff *int64  `json:"max_key_staff,omitempty" db:"max_key_staff"`
	MaximumRoleKeys *int64  `json:"max_role_keys,omitempty" db:"max_role_keys"`
	MultiLogin      *bool   `json:"multi_login,omitempty" db:"multi_login"`
	StaffNumber     *string `json:"dlr_staff_no,omitempty" db:"dlr_staff_no"`
	StaffInfo       *string `json:"staff_info,omitempty" db:"staff_info"`
	CreatedAt       *int64  `json:"created_at,omitempty" db:"created_at"`
	ProfileImage    *string `json:"profile_image,omitempty" db:"profile_image"`
	Region          *string `json:"region,omitempty" db:"region"`
	City            *string `json:"city,omitempty" db:"city"`
	State           *string `json:"state,omitempty" db:"state"`
	PostalCode      *string `json:"postal_code,omitempty" db:"postal_code"`
	Street1Address  *string `json:"street1_address,omitempty" db:"street1_address"`
	Street2Address  *string `json:"street2_address,omitempty" db:"street2_address"`
	Currency        *string `json:"currency,omitempty" db:"currency"`
}

func NewStaff(emailWork, firstName, lastName, username, password, phoneNumber, emailPersonal, roleName, staffInfo, profile_image, region, city, State, postal_code, street1_address, street2_address, currency *string, multiLogin *bool) *Staff {
	return &Staff{
		StaffId:         new(string),
		StaffRoleName:   roleName,
		DealerID:        new(string),
		EmailWork:       emailWork,
		FirstName:       firstName,
		LastName:        lastName,
		Username:        username,
		Password:        password,
		PhoneNumber:     phoneNumber,
		EmailPersonal:   emailPersonal,
		MaximumKeyStaff: new(int64),
		MaximumRoleKeys: new(int64),
		MultiLogin:      multiLogin,
		StaffNumber:     new(string),
		StaffInfo:       staffInfo,
		CreatedAt:       new(int64),
		ProfileImage:    profile_image,
		Region:          region,
		City:            city,
		State:           State,
		PostalCode:      postal_code,
		Street1Address:  street1_address,
		Street2Address:  street2_address,
		Currency:        currency,
	}
}

type AssignUnassignKey struct {
	KeyAssignID        *string `json:"key_assign_id" db:"key_assign_id"`
	VehicleID          *string `json:"veh_id" db:"veh_id"`
	KeyID              *string `json:"key_id" db:"key_id"`
	StaffID            *string `json:"staff_id" db:"staff_id"`
	KeyAssignStatus    *bool   `json:"key_assign_status" db:"key_assign_status"`
	TransferID         *string `json:"transfer_id" db:"transfer_id"`
	TransfereID        *string `json:"transfere_id" db:"transfere_id"`
	AssignComment      *string `json:"assign_cmnt" db:"assign_cmnt"`
	DataKeyAssignStart *int64  `json:"date_key_assign_start" db:"date_key_assign_start"`
	DataKeyAssignStop  *int64  `json:"date_key_assign_stop" db:"date_key_assign_stop"`
}

func NewAssignUnassignKey(vehicleID, keyID, staffID *string, keyAssignStatus *bool) *AssignUnassignKey {
	return &AssignUnassignKey{
		KeyAssignID:        new(string),
		VehicleID:          vehicleID,
		KeyID:              keyID,
		StaffID:            staffID,
		KeyAssignStatus:    keyAssignStatus,
		TransferID:         new(string),
		TransfereID:        new(string),
		AssignComment:      new(string),
		DataKeyAssignStart: new(int64),
		DataKeyAssignStop:  new(int64),
	}
}

type StaffAssignedVehicleKeys struct {
	VehicleName       *string `json:"veh_name,omitempty" db:"veh_name"`
	VehicleID         *string `json:"veh_id,omitempty" db:"veh_id"`
	DealerID          *string `json:"dlr_id,omitempty" db:"dlr_id"`
	StaffID           *string `json:"staff_id" db:"staff_id"`
	KeyID             *string `json:"key_id" db:"key_id"`
	KeyAssignStatus   *bool   `json:"key_assign_status" db:"key_assign_status"`
	AttachId          *string `json:"attach_id,omitempty" db:"attach_id"`
	StockNumber       *int64  `json:"stock_no,omitempty" db:"stock_no"`
	VinNumber         *string `json:"vin,omitempty" db:"vin"`
	LicPlate          *string `json:"lic_plate,omitempty" db:"lic_plate"`
	CarStatusID       *string `json:"car_status_id,omitempty" db:"car_status_id"`
	VechicleTypeID    *string `json:"veh_type_id,omitempty" db:"veh_type_id"`
	BrandMake         *string `json:"make,omitempty" db:"make"`
	CreatedAt         *int64  `json:"created_at" db:"created_at"`
	Model             *string `json:"model,omitempty" db:"model"`
	ManfactureYear    *string `json:"year,omitempty" db:"year"`
	VechicleColor     *string `json:"veh_color,omitempty"  db:"veh_color"`
	TransmissionStyle *string `json:"transmission_style,omitempty" db:"transmission_style"`
}

type AssignedKeys struct {
	VehicleID         *string `json:"veh_id" db:"veh_id"`
	KeyID             *string `json:"key_id" db:"key_id"`
	StaffID           *string `json:"staff_id" db:"staff_id"`
	KeyAssignStatus   *bool   `json:"key_assign_status" db:"key_assign_status"`
	TransferID        *string `json:"transfer_id" db:"transfer_id"`
	StaffRole         *string `json:"staff_role_name" db:"staff_role_name"`
	VehName           *string `json:"veh_name" db:"veh_name"`
	VehColor          *string `json:"veh_color" db:"veh_color"`
	StockNo           *string `json:"stock_no" db:"stock_no"`
	Make              *string `json:"make" db:"make"`
	Model             *string `json:"model" db:"model"`
	Year              *string `json:"year" db:"year"`
	TransmissionStyle *string `json:"transmission_style" db:"transmission_style"`
	FirstName         *string `json:"nam_first" db:"nam_first"`
	LastName          *string `json:"nam_last" db:"nam_last"`
	StaffUsername     *string `json:"staff_username" db:"staff_username"`
}

type VehicleLocation struct {
	VehLocID   *string `json:"veh_loc_id" db:"veh_loc_id"`
	LocID      *string `json:"loc_id" db:"loc_id"`
	VehID      *string `json:"veh_id" db:"veh_id"`
	SpaceID    *string `json:"space_id" db:"space_id"`
	UserIN     *string `json:"user_in" db:"user_in"`
	UserOUT    *string `json:"user_out" db:"user_out"`
	DateInLoc  *int64  `json:"date_in_loc" db:"date_in_loc"`
	DateOutLoc *int64  `json:"date_out_loc" db:"date_out_loc"`
}

func NewVehicleLocation(nil, LocID, VehID, SpaceID, UserIN, UserOUT *string, DateInLoc, DateOutLoc *int64) *VehicleLocation {
	return &VehicleLocation{
		VehLocID:   new(string),
		LocID:      LocID,
		VehID:      VehID,
		SpaceID:    SpaceID,
		UserIN:     UserIN,
		UserOUT:    UserOUT,
		DateInLoc:  DateInLoc,
		DateOutLoc: DateOutLoc,
	}
}

type GetAllDealers struct {
	DealerId   *string `json:"dlr_id" db:"dlr_id"`
	Username   *string `json:"username" db:"dlr_username"`
	DealerName *string `json:"dlr_name" db:"dlr_name"`
}

func NewGetAllDealers(dlr_id, username, dlr_name *string) *GetAllDealers {
	return &GetAllDealers{
		DealerId:   new(string),
		Username:   username,
		DealerName: dlr_name,
	}

}

func NewUpdatedVehicle(vehicleID, vehicleName, vinNumber, licPlate, brandMake, model, manfactureYear, vechicleColor, TransmissionStyle, MapColour, Note *string, stockNumber *int64, VehAvailable *bool) *Vehicle {
	return &Vehicle{
		DealerID:          new(string),
		CarStatusID:       new(string),
		VechicleTypeID:    new(string),
		AttachId:          new(string),
		VehicleID:         vehicleID,
		VehicleName:       vehicleName,
		StockNumber:       stockNumber,
		VinNumber:         vinNumber,
		LicPlate:          licPlate,
		BrandMake:         brandMake,
		Model:             model,
		CreatedAt:         new(int64),
		ManfactureYear:    manfactureYear,
		VechicleColor:     vechicleColor,
		TransmissionStyle: TransmissionStyle,
		MapColour:         MapColour,
		Note:              Note,
		VehAvailable:      VehAvailable,
	}
}

func NewUpdatedKey(VehID, KeyID, LocName, KeyName, keyAccessible, recievedFrom, receivedBy *string) *Key {
	return &Key{
		KeyID:         new(string),
		VehID:         new(string),
		KeyName:       KeyName,
		KeyAccessible: keyAccessible,
		ReceivedFrom:  recievedFrom,
		ReceivedBy:    receivedBy,
		KeyAvailable:  new(bool),
		LocName:       LocName,
	}
}

type Map struct {
	MapName     *string        `json:"map_name" binding:"required,map_name" db:"map_name"`
	MapID       *string        `json:"map_id" db:"map_id"`
	DealerID    *string        `json:"dealer_id,omitempty" db:"dlr_id"`
	Coords      *string        `json:"coords,omitempty" db:"map_polygon"`
	ByteCoords  []byte         `json:"byte_coords,omitempty" db:"st_asgeojson"`
	Coordinates []*Coordinates `json:"map_coords,omitempty"`
	ParkingLots []*ParkingLot  `json:"parkings_lots"`
}
type ParkingLot struct {
	ParkingName  *string        `json:"parking_name,omitempty"  binding:"required,parking_name" db:"parking_name"`
	ParkingLotID *string        `json:"parking_lot_id,omitempty" db:"parking_lot_id"`
	MapID        *string        `json:"map_id" db:"map_id"`
	Coords       *string        `json:"coords,omitempty" db:"parking_polygon"`
	ByteCoords   []byte         `json:"byte_coords,omitempty" db:"st_asgeojson"`
	Coordinates  []*Coordinates `json:"parking_coords,omitempty"`
	Zones        []*Zone        `json:"zones"`
}
type Zone struct {
	ZoneName     *string        `json:"zone_name,omitempty"  binding:"required,zone_name" db:"zone_name"`
	ZoneID       *string        `json:"zone_id,omitempty" db:"zone_id"`
	ParkingLotID *string        `json:"parking_lot_id,omitempty" db:"parking_lot_id"`
	Coords       *string        `json:"coords,omitempty" db:"zone_polygon"`
	ByteCoords   []byte         `json:"byte_coords,omitempty" db:"st_asgeojson"`
	Coordinates  []*Coordinates `json:"zone_coords,omitempty"`
	Rows         []*Row         `json:"rows"`
}
type Row struct {
	RowsName    *string        `json:"row_name,omitempty"  binding:"required,row_name" db:"row_name"`
	RowID       *string        `json:"row_id,omitempty" db:"row_id"`
	ZoneID      *string        `json:"zone_id,omitempty" db:"zone_id"`
	Coords      *string        `json:"coords,omitempty" db:"row_polygon"`
	ByteCoords  []byte         `json:"byte_coords,omitempty" db:"st_asgeojson"`
	Coordinates []*Coordinates `json:"row_coords,omitempty"`
	Slots       []*Slot        `json:"slots"`
}
type Slot struct {
	SpacingNumber *int64  `json:"spacing_no,omitempty"  binding:"required,spacing_no" db:"spacing_no"`
	SlotAvailable *bool   `json:"slot_available" db:"slot_available"`
	SlotName      *string `json:"slot_name"  binding:"required,slot_name" db:"slot_name"`
	SlotID        *string `json:"slot_id,omitempty" db:"slot_id"`
	RowID         *string `json:"row_id,omitempty" db:"row_id"`
	ByteCoords    []byte  `json:"byte_coords,omitempty" db:"st_asgeojson"`
	//vehicle id added for check-in check out
	VehicleID   *string        `json:"veh_id,omitempty"  db:"veh_id"`
	CarColor    *string        `json:"car_color,omitempty"  binding:"required,car_color" db:"car_color"`
	Coords      *string        `json:"coords,omitempty" db:"slot_polygon"`
	Coordinates []*Coordinates `json:"slot_coords,omitempty"`
	// vehicle data
	VehicleName       *string `json:"veh_name,omitempty" db:"veh_name"`
	LicPlate          *string `json:"lic_plate,omitempty" db:"lic_plate"`
	BrandMake         *string `json:"make,omitempty" db:"make"`
	ManfactureYear    *string `json:"year,omitempty" db:"year"`
	VechicleColor     *string `json:"veh_color,omitempty"  db:"veh_color"`
	TransmissionStyle *string `json:"transmission_style,omitempty" db:"transmission_style"`
}

type Coordinates struct {
	Latitude  *float64 `json:"latitude,omitempty" db:"latitude"`
	Longitude *float64 `json:"longitude,omitempty" db:"longitude"`
}

//models for check-in out vehicle in map
type CheckInOutVehicleMap struct {
	VehicleID     *string `json:"veh_id,omitempty" db:"veh_id"`
	MapID         *string `json:"map_id,omitempty" db:"map_id"`
	SlotAvailable *bool   `json:"check_in_out_status,omitempty" db:"slot_available"`
	SlotID        *string `json:"slot_id,omitempty" db:"slot_id"`
	DealerID      string  `json:"check_in_out_id,omitempty" db:"dlr_id"`
}
type Coords struct {
	Coordinates [][][2]*float64 `json:"coordinates"`
}

//models created for adding vehicle to map
type AddingVehicleToMap struct {
	SlotAvailable *bool   `json:"slot_available,omitempty" db:"slot_available"`
	SlotID        *string `json:"slot_id,omitempty" db:"slot_id"`
	VehicleID     *string `json:"veh_id,omitempty" db:"veh_id"`
}
