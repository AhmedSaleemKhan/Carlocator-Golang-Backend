package models

type DealerSignupRequest struct {
	Email    *string `json:"email" binding:"required,email"`
	Username *string `json:"username" binding:"required,gt=0"`
	Password *string `json:"password" binding:"required,password"`
}

func NewDealerSignupRequest() *DealerSignupRequest {
	return &DealerSignupRequest{
		Email:    new(string),
		Username: new(string),
		Password: new(string),
	}
}

type OTPRequest struct {
	Email *string `json:"email" binding:"required,email"`
	OTP   *string `json:"otp" binding:"required,len=6"`
}

func NewOTPRequest() *OTPRequest {
	return &OTPRequest{
		Email: new(string),
		OTP:   new(string),
	}
}

type DealerForgotPasswordRequest struct {
	Email *string `json:"email" binding:"required,email"`
}

func NewDealerforgotpasswordRequest() *DealerForgotPasswordRequest {
	return &DealerForgotPasswordRequest{
		Email: new(string),
	}
}

type DealerForgotConfirmRequest struct {
	Email       *string `json:"email" binding:"required,email"`
	OTP         *string `json:"otp" binding:"required,len=6"`
	NewPassword *string `json:"newpassword" binding:"required,password"`
}

func NewDealerForgotConfirmRequest() *DealerForgotConfirmRequest {
	return &DealerForgotConfirmRequest{
		Email:       new(string),
		OTP:         new(string),
		NewPassword: new(string),
	}
}

type SignInRequest struct {
	Email    *string `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required,password"`
}

func NewSignInRequest() *SignInRequest {
	return &SignInRequest{
		Email:    new(string),
		Password: new(string),
	}
}

type RefreshAccessTokenRequest struct {
	Token *string `json:"token"`
}

func NewRefreshAccessTokenRequest() *RefreshAccessTokenRequest {
	return &RefreshAccessTokenRequest{
		Token: new(string),
	}
}

type DealerSignOutRequest struct {
	Token *string `json:"token" binding:"required,jwt"`
}

func NewDealerSignOutRequest() *DealerSignOutRequest {
	return &DealerSignOutRequest{
		Token: new(string),
	}
}

type VinNumberDecode struct {
	VinNumber           *string `json:"vin_number"`
	Make                *string `json:"make"`
	Model               *string `json:"model"`
	ManufacturerName    *string `json:"ManufacturerName"`
	ModelYear           *string `json:"modelYear"`
	ManufacturedCity    *string `json:"ManufacturerCity"`
	VehicleType         *string `json:"VehicleType"`
	ManufacturedCountry *string `json:"ManufacturedCountry"`
	BodyType            *string `json:"BodyType"`
	TransmissionStyle   *string `json:"TransmissionStyle"`
	FuelType            *string `json:"Fuel Type"`
}

type VehicleVinNumberResponse struct {
	VinNumberDecode
}

func NewVehicleNumberResponse() *VehicleVinNumberResponse {
	vehicleVinNumber := VehicleVinNumberResponse{}
	vehicleVinNumber.VinNumber = new(string)
	vehicleVinNumber.Make = new(string)
	vehicleVinNumber.Model = new(string)
	vehicleVinNumber.ManufacturerName = new(string)
	vehicleVinNumber.ModelYear = new(string)
	vehicleVinNumber.ManufacturedCity = new(string)
	vehicleVinNumber.VehicleType = new(string)
	vehicleVinNumber.ManufacturedCountry = new(string)
	vehicleVinNumber.BodyType = new(string)
	vehicleVinNumber.TransmissionStyle = new(string)
	vehicleVinNumber.FuelType = new(string)
	return &vehicleVinNumber
}

type MultipleVinNumberRequest struct {
	VinNumbers []string `json:"vin_numbers"`
}

func NewMultipleVinNumberRequest() *MultipleVinNumberRequest {
	return &MultipleVinNumberRequest{}
}

type MultipleVinNumberResponse struct {
	Results map[string]VinNumberDecode `json:"results"`
}

type VehicleRequest struct {
	VehicleName       *string `json:"veh_name" db:"veh_name" binding:"required,gt=0"`
	VinNumber         *string `json:"vin_no" db:"vin_no" binding:"required,gt=0"`
	LicPlate          *string `json:"lic_plate" db:"lic_plate" binding:"required,gt=0"`
	BrandMake         *string `json:"brand_make" db:"brand_make" binding:"required,gt=0"`
	Model             *string `json:"model" db:"model" binding:"required,gt=0"`
	ManfactureYear    *string `json:"manufacture_year" db:"manufacture_year" binding:"required,gt=0" `
	VechicleColor     *string `json:"veh_color"  db:"veh_color" binding:"required,gt=0"`
	TransmissionStyle *string `json:"transmission_style" db:"transmission_style" binding:"required,gt=0"`
}

func NewVehicleRequest() *VehicleRequest {
	return &VehicleRequest{
		VehicleName:       new(string),
		VinNumber:         new(string),
		LicPlate:          new(string),
		BrandMake:         new(string),
		Model:             new(string),
		ManfactureYear:    new(string),
		VechicleColor:     new(string),
		TransmissionStyle: new(string),
	}
}

type CreateKeyRequest struct {
	KeyID         *string `json:"key_id" `
	VehID         *string `json:"veh_id"`
	KeyName       *string `json:"key_name"  binding:"required,gt=0"`
	KeyAccessible *string `json:"key_accessible"  binding:"required,gt=0"`
	ReceivedFrom  *string `json:"received_from"  binding:"required,gt=0"`
	ReceivedBy    *string `json:"received_by"  binding:"required,gt=0"`
	LocName       *string `json:"loc_name" binding:"required"`
}

type CreateKeyResponse struct {
	Key *Key `json:"key"`
}

func NewCreateKeyRequest() *CreateKeyRequest {
	return &CreateKeyRequest{
		KeyID:         new(string),
		VehID:         new(string),
		KeyName:       new(string),
		KeyAccessible: new(string),
		ReceivedFrom:  new(string),
		ReceivedBy:    new(string),
		LocName:       new(string),
	}
}

type AddMultipleVehiclesRequest struct {
	Vehicles []*Vehicle `json:"vehicles"`
}

func NewAddMultipleVehiclesRequest() *AddMultipleVehiclesRequest {
	return &AddMultipleVehiclesRequest{
		[]*Vehicle{},
	}
}

type StaffSignupRequest struct {
	EmailWork      *string `json:"email_work" binding:"required,email"`
	FirstName      *string `json:"first_name" binding:"required,gt=0"`
	LastName       *string `json:"last_name" binding:"required,gt=0"`
	Username       *string `json:"username" binding:"required,gt=0"`
	Password       *string `json:"password" binding:"required,password"`
	PhoneNumber    *string `json:"phone_number" binding:"required,gt=0"`
	EmailPersonal  *string `json:"email_personal" binding:"required,email"`
	MultiLogin     *bool   `json:"multi_login"`
	StaffRoleName  *string `json:"staff_role_name" binding:"required,gt=0"`
	StaffInfo      *string `json:"staff_info" binding:"required,gt=0"`
	ProfileImage   *string `json:"profile_image" binding:"required,gt=0"`
	Region         *string `json:"region"`
	City           *string `json:"city"`
	State          *string `json:"state"`
	PostalCode     *string `json:"Postal_Code" binding:"required,gt=0"`
	Street1Address *string `json:"street1_address"`
	Street2Address *string `json:"street2_address"`
	Currency       *string `json:"currency"`
}

func NewStaffSignupRequest() *StaffSignupRequest {
	return &StaffSignupRequest{
		EmailWork:      new(string),
		FirstName:      new(string),
		LastName:       new(string),
		Username:       new(string),
		Password:       new(string),
		PhoneNumber:    new(string),
		EmailPersonal:  new(string),
		MultiLogin:     new(bool),
		StaffRoleName:  new(string),
		StaffInfo:      new(string),
		ProfileImage:   new(string),
		Region:         new(string),
		City:           new(string),
		State:          new(string),
		PostalCode:     new(string),
		Street1Address: new(string),
		Street2Address: new(string),
		Currency:       new(string),
	}
}

type AssignUnassignKeyResquest struct {
	KeyAssignID        *string `json:"key_assign_id" validate:"required,uuid4"`
	VehicleID          *string `json:"veh_id"  validate:"required,uuid4"`
	KeyID              *string `json:"key_id"  validate:"required,uuid4"`
	StaffID            *string `json:"staff_id"  validate:"required,uuid4"`
	KeyAssignStatus    *bool   `json:"key_assign_status"`
	TransferID         *string `json:"transfer_id"`
	TransfereID        *string `json:"transfere_id"`
	AssignComment      *string `json:"assign_cmnt "`
	DataKeyAssignStart *int64  `json:"date_key_assign_start"`
	DataKeyAssignStop  *int64  `json:"date_key_assign_stop"`
}

func NewAssignUnassignKeyResquest() *AssignUnassignKeyResquest {
	return &AssignUnassignKeyResquest{
		KeyAssignID:        new(string),
		VehicleID:          new(string),
		KeyID:              new(string),
		StaffID:            new(string),
		KeyAssignStatus:    new(bool),
		TransferID:         new(string),
		TransfereID:        new(string),
		AssignComment:      new(string),
		DataKeyAssignStart: new(int64),
		DataKeyAssignStop:  new(int64),
	}
}

type StaffResetPasswordRequest struct {
	Password *string `json:"password" binding:"required,password"`
}

func NewStaffResetPasswordRequest() *StaffResetPasswordRequest {
	return &StaffResetPasswordRequest{
		Password: new(string),
	}
}

type UpdateDealershipRequest struct {
	Dealer *Dealer `json:"dealer"`
}

func NewUpdateDealerShipRequest() *UpdateDealershipRequest {
	return &UpdateDealershipRequest{
		Dealer: &Dealer{
			DealerId:       new(string),
			GroupId:        new(string),
			Region:         new(string),
			City:           new(string),
			State:          new(string),
			PostalCode:     new(int64),
			UpdatedAt:      new(int64),
			Street1Address: new(string),
			Street2Address: new(string),
			ShortName:      new(string),
			DealerLogo:     new(string),
			DealerName:     new(string),
			DealerTaxId:    new(string),
			ParentDealerId: new(string),
			PaymentMethod:  new(string),
			BranchName:     new(string),
			BilingPeriod:   new(string),
			DealerDiscord:  new(string),
			DealerComment:  new(string),
			LocNote:        new(string),
			LanguageId:     new(string),
			CountryId:      new(string),
			Telephone:      new(int64),
			Currency:       new(string),
			Website:        new(string),
			MilesKm:        new(int64),
			Sic_code:       new(int64),
		},
	}
}

type VehicleLocationRequest struct {
	LocID      *string `json:"loc_id" validate:"required,uuid4"`
	VehID      *string `json:"veh_id" validate:"required,uuid4"`
	SpaceID    *string `json:"space_id" validate:"required,uuid4"`
	UserIN     *string `json:"user_in" validate:"required,uuid4"`
	UserOUT    *string `json:"user_out" binding:"required,gt=0"`
	DateInLoc  *int64  `json:"date_in_loc" binding:"required,gt=0"`
	DateOutLoc *int64  `json:"date_out_loc" binding:"required,gt=0"`
}

type GetAllDealersRequest struct {
	DealerId   *string `json:"dlr_id" validate:"required,uuid4"`
	Username   *string `json:"username" binding:"required,gt=0"`
	DealerName *string `json:"dlr_name" binding:"required,gt=0"`
}

func NewGetAllDealersRequest() *GetAllDealersRequest {
	return &GetAllDealersRequest{
		DealerId:   new(string),
		Username:   new(string),
		DealerName: new(string),
	}
}

func NewVehicleLocationRequest() *VehicleLocationRequest {
	return &VehicleLocationRequest{
		LocID:      new(string),
		VehID:      new(string),
		SpaceID:    new(string),
		UserIN:     new(string),
		UserOUT:    new(string),
		DateInLoc:  new(int64),
		DateOutLoc: new(int64),
	}
}

type UpdatedVehicleRequest struct {
	Vehicle *Vehicle `json:"vehicle"`
}

func NewUpdatedVehicleRequest() *UpdatedVehicleRequest {
	return &UpdatedVehicleRequest{
		Vehicle: &Vehicle{
			DealerID:          new(string),
			VehicleID:         new(string),
			VehicleName:       new(string),
			StockNumber:       new(int64),
			VinNumber:         new(string),
			LicPlate:          new(string),
			BrandMake:         new(string),
			Model:             new(string),
			ManfactureYear:    new(string),
			VechicleColor:     new(string),
			TransmissionStyle: new(string),
			MapColour:         new(string),
			Note:              new(string),
			VehAvailable:      new(bool),
		},
	}
}

type UpdatedKeyRequest struct {
	Key *Key `json:"key"`
}

func NewUpdatedKeyRequest() *UpdatedKeyRequest {
	return &UpdatedKeyRequest{
		Key: &Key{
			KeyID:         new(string),
			VehID:         new(string),
			KeyName:       new(string),
			KeyAccessible: new(string),
			ReceivedFrom:  new(string),
			ReceivedBy:    new(string),
			LocName:       new(string),
		},
	}
}

type ResendOTPRequest struct {
	Email *string `json:"email" binding:"required,email"`
}

func NewResendOTPRequest() *ResendOTPRequest {
	return &ResendOTPRequest{
		Email: new(string),
	}
}

type CreateMapRequest struct {
	Maps []*Map `json:"map"`
}

func NewMapRequest() *CreateMapRequest {
	return &CreateMapRequest{
		[]*Map{},
	}
}

type CheckInOutVehicleRequest struct {
	CheckInOutVehicle *CheckInOutVehicleMap `json:"check_in_out_vehicle"`
}

//request if user wants to checkin or out the vehicle
func NewCheckInOutVehicleRequest() *CheckInOutVehicleRequest {
	return &CheckInOutVehicleRequest{
		&CheckInOutVehicleMap{
			VehicleID: new(string),
			SlotID:    new(string),
		},
	}
}

type AddingVehicleToMapRequest struct {
	AddVehicleToMap *AddingVehicleToMap `json:"vehicle_to_map"`
}

func NewAddingVehicleToMapRequest() *AddingVehicleToMapRequest {
	return &AddingVehicleToMapRequest{
		&AddingVehicleToMap{
			VehicleID: new(string),
			SlotID:    new(string),
		},
	}
}
