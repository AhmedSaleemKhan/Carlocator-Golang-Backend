package models

type VehicleResponse struct {
	VehicleID         *string `json:"veh_id" db:"veh_id"`
	DealerID          *string `json:"dlr_id" db:"dlr_id"`
	CarStatusID       *string `json:"car_status_id" db:"car_status_id"`
	VechicleTypeID    *string `json:"veh_type_id" db:"veh_type_id"`
	AttachId          *string `json:"attach_id" db:"attach_id"`
	VehicleName       *string `json:"veh_name" db:"veh_name"`
	StockNumber       *int64  `json:"stock_no" db:"stock_no"`
	VinNumber         *string `json:"vin_no" db:"vin"`
	LicPlate          *string `json:"lic_plate" db:"lic_plate"`
	BrandMake         *string `json:"brand_make" db:"make"`
	Model             *string `json:"model" db:"model"`
	ManfactureYear    *string `json:"manufacture_year" db:"year" `
	VechicleColor     *string `json:"veh_color"  db:"veh_color"`
	TransmissionStyle *string `json:"transmission_style" db:"transmission_style"`
	CreatedAt         *int64  `json:"created_at" db:"created_at"`
}

func NewVehicleResponse() *VehicleResponse {
	return &VehicleResponse{
		VehicleID:         new(string),
		DealerID:          new(string),
		CarStatusID:       new(string),
		VechicleTypeID:    new(string),
		AttachId:          new(string),
		StockNumber:       new(int64),
		VinNumber:         new(string),
		LicPlate:          new(string),
		BrandMake:         new(string),
		Model:             new(string),
		ManfactureYear:    new(string),
		VechicleColor:     new(string),
		TransmissionStyle: new(string),
		CreatedAt:         new(int64),
	}
}

type NHTSAVinDecodeModel struct {
	Count          int    `json:"Count"`
	Message        string `json:"Message"`
	SearchCriteria string `json:"SearchCriteria"`
	Results        []struct {
		Value      string `json:"Value"`
		ValueID    string `json:"ValueId"`
		Variable   string `json:"Variable"`
		VariableID int    `json:"VariableId"`
	} `json:"Results"`
}

type DealerResponse struct {
	DealerId *string `json:"dlr_id" db:"dlr_id"`
	Email    *string `json:"email" db:"email"`
	Username *string `json:"username" db:"dlr_username"`
}

func NewDealerResponse() *DealerResponse {
	return &DealerResponse{
		DealerId: new(string),
		Email:    new(string),
		Username: new(string),
	}
}

type VerifyOTPResponse struct {
	Email *string `json:"email" db:"email"`
}

func NewVerifyOTPResponse() *VerifyOTPResponse {
	return &VerifyOTPResponse{
		Email: new(string),
	}
}

type UniqueVehicleAttributes struct {
	BrandMakes      []string `json:"makes"`
	Models          []string `json:"models"`
	ManfactureYears []string `json:"years"`
}

type GetDealerVehiclesResponse struct {
	Vehicles   []*Vehicle               `json:"vehicles"`
	Atrributes *UniqueVehicleAttributes `json:"attributes"`
	Metadata   *Metadata                `json:"metadata"`
}

type MultipleVehiclesResponse struct {
	Vehicles []*Vehicle `json:"vehicles"`
}

func NewMultipleVehiclesResponse(vehicles []*Vehicle) MultipleVehiclesResponse {
	return MultipleVehiclesResponse{
		Vehicles: vehicles,
	}

}

type StaffSignupResponse struct {
	StaffId         *string `json:"staff_id,omitempty" db:"staff_id"`
	StaffRoleName   *string `json:"staff_role_name,omitempty" db:"staff_role_name"`
	DealerId        *string `json:"dlr_id,omitempty" db:"dlr_id"`
	EmailWork       *string `json:"email_work,omitempty" db:"staff_email_work"`
	FirstName       *string `json:"first_name,omitempty" db:"nam_first"`
	LastName        *string `json:"last_name,omitempty" db:"nam_last"`
	Username        *string `json:"username,omitempty" db:"staff_username"`
	PhoneNumber     *string `json:"phone_number,omitempty" db:"staff_phone_cell"`
	EmailPersonal   *string `json:"email_personal,omitempty" db:"staff_email_personal"`
	MaximumKeyStaff *int64  `json:"max_key_staff,omitempty" db:"max_key_staff"`
	MaximumRoleKeys *int64  `json:"max_role_keys,omitempty" db:"max_role_keys"`
	MultiLogin      *bool   `json:"multi_login,omitempty" db:"multi_login"`
	StaffNumber     *int64  `json:"staff_no,omitempty" db:"dlr_staff_no"`
	StaffInfo       *string `json:"staff_info,omitempty" db:"staff_info"`
	ProfileImage    *string `json:"profile_image,omitempty" db:"profile_image"`
	Region          *string `json:"region,omitempty" db:"region"`
	City            *string `json:"city,omitempty" db:"city"`
	State           *string `json:"state,omitempty" db:"state"`
	PostalCode      *string `json:"postal_code,omitempty" db:"postal_code"`
	Street1Address  *string `json:"street1_address,omitempty" db:"street1_address"`
	Street2Address  *string `json:"street2_address,omitempty" db:"street2_address"`
	Currency        *string `json:"currency,omitempty" db:"currency"`
}

func NewStaffSignupResponse(staffId, roleName, dealerId, emailWork, firstName, lastName, username, phoneNumber, emailPersonal, staffInfo, profile_image, region, city, State, postal_code, street1_address, street2_address, currency *string, maxKey, maxRole *int64, staffNumber *string, multiLogin *bool) *StaffSignupResponse {
	return &StaffSignupResponse{
		StaffId:         new(string),
		StaffRoleName:   roleName,
		DealerId:        dealerId,
		EmailWork:       emailWork,
		FirstName:       firstName,
		LastName:        lastName,
		Username:        username,
		PhoneNumber:     phoneNumber,
		EmailPersonal:   emailPersonal,
		MaximumKeyStaff: new(int64),
		MaximumRoleKeys: new(int64),
		MultiLogin:      multiLogin,
		StaffNumber:     new(int64),
		StaffInfo:       staffInfo,
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

type DeleteKeyResponse struct {
	Message string `json:"message"`
}

type AssignUnassignKeyResponse struct {
	KeyAssignID        *string `json:"key_assign_id"`
	VehicleID          *string `json:"veh_id"`
	KeyID              *string `json:"key_id"`
	StaffID            *string `json:"staff_id"`
	KeyAssignStatus    *bool   `json:"key_assign_status"`
	TransferID         *string `json:"transfer_id"`
	TransfereID        *string `json:"transfere_id"`
	AssignComment      *string `json:"assign_cmnt"`
	DataKeyAssignStart *int64  `json:"date_key_assign_start"`
	DataKeyAssignStop  *int64  `json:"date_key_assign_stop"`
}

func NewAssignUnassignKeyResponse(keyAssignID, vehicleID, keyID, staffID *string, keyAssignStatus *bool, transferID, transfereID, assignComment *string, dataKeyAssignStart, dataKeyAssignStop *int64) *AssignUnassignKeyResponse {
	return &AssignUnassignKeyResponse{
		KeyAssignID:        keyAssignID,
		VehicleID:          vehicleID,
		KeyID:              keyID,
		StaffID:            staffID,
		KeyAssignStatus:    keyAssignStatus,
		TransferID:         transferID,
		TransfereID:        transfereID,
		AssignComment:      assignComment,
		DataKeyAssignStart: dataKeyAssignStart,
		DataKeyAssignStop:  dataKeyAssignStop,
	}
}

type GetVehicleKeysRespone struct {
	Keys []*Key `json:"keys"`
}
type ResetPasswordResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type DeleteVehicleResponse struct {
	Message string `json:"message"`
}
type GetStaffMembersResponse struct {
	StaffMembers []*Staff  `json:"staff"`
	Metadata     *Metadata `json:"metadata"`
}

type GetDealerInfoResponse struct {
	DealerInfo []*Dealer `json:"dealer"`
}

type GetUserAssignedKeysResponse struct {
	AssignedKeys []*AssignedKeys `json:"assignedkey"`
	Metadata     *Metadata       `json:"metadata"`
}

type VehicleLocationResponse struct {
	VehLocID   *string `json:"veh_loc_id" db:"veh_loc_id"`
	LocID      *string `json:"loc_id" db:"loc_id"`
	VehID      *string `json:"veh_id" db:"veh_id"`
	SpaceID    *string `json:"space_id" db:"space_id"`
	UserIN     *string `json:"user_in" db:"user_in"`
	UserOUT    *string `json:"user_out" db:"user_out"`
	DateInLoc  *int64  `json:"date_in_loc" db:"date_in_loc"`
	DateOutLoc *int64  `json:"date_out_loc" db:"date_out_loc"`
}

func NewVehicleLocationResponse(vehLocID, LocID, VehID, SpaceID, UserIN, UserOUT *string, DateInLoc, DateOutLoc *int64) *VehicleLocationResponse {
	return &VehicleLocationResponse{
		VehLocID:   vehLocID,
		LocID:      LocID,
		VehID:      VehID,
		SpaceID:    SpaceID,
		UserIN:     UserIN,
		UserOUT:    UserOUT,
		DateInLoc:  DateInLoc,
		DateOutLoc: DateOutLoc,
	}
}

type S3FileUploadResponse struct {
	URL *string `json:"url"`
}

type GetDealersResponse struct {
	GetAllDealers []*GetAllDealers `json:"getalldealer"`
}
type UpdateDealershipResponse struct {
	Dealer *Dealer `json:"dealer"`
}

func NewUpdateDealershipResponse(dealer *Dealer) *UpdateDealershipResponse {
	return &UpdateDealershipResponse{
		Dealer: dealer,
	}
}

type UpdatedVehicleResponse struct {
	Vehicle *Vehicle `json:"vehicle"`
}

func NewUpdatedVehicleResponse(vehicle *Vehicle) *UpdatedVehicleResponse {
	return &UpdatedVehicleResponse{
		Vehicle: vehicle,
	}
}

type UpdatedKeyResponse struct {
	Key *Key `json:"key"`
}

func NewUpdatedKeyResponse(key *Key) *UpdatedKeyResponse {
	return &UpdatedKeyResponse{
		Key: key,
	}
}

type ResendOTPResponse struct {
	Email *string `json:"email"`
}

func NewResendOTPResponse(email *string) *ResendOTPResponse {
	return &ResendOTPResponse{
		Email: new(string),
	}
}

type CreateMapResponse struct {
	Map []*Map `json:"map"`
}

func NewMapResponse(maps []*Map) *CreateMapResponse {
	return &CreateMapResponse{
		Map: maps,
	}
}

type GetDealerMapResponse struct {
	DealerMaps []*Map `json:"maps"`
}
type StaffAssignedVehicleKeysResponse struct {
	Vehicles []*StaffAssignedVehicleKeys `json:"staff_assigned_vehicles"`
}

type CheckInOutVehicleResponse struct {
	VehicleResponse *CheckInOutVehicleMap `json:"vehicle_checked"`
}

func NewCheckInOutVehicleResponse(response *CheckInOutVehicleMap) *CheckInOutVehicleResponse {
	return &CheckInOutVehicleResponse{
		VehicleResponse: response,
	}
}

type AddingVehicleToMapResponse struct {
	AddVehicleResponse *AddingVehicleToMap `json:"vehicle_added"`
}

func NewAddingVehicleToMapResponse(response *AddingVehicleToMap) *AddingVehicleToMapResponse {
	return &AddingVehicleToMapResponse{
		AddVehicleResponse: response,
	}
}
