package server

import (
	"net/http"

	"capregsoft.com/carlocator/service/lib"
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) DealerSignup(c *gin.Context) {
	body := models.NewDealerSignupRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	dealer := models.NewDealer(body.Email, body.Username, body.Password)
	resp, err := s.api.DealerSignupAPI(c, dealer)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) SignIn(c *gin.Context) {
	body := models.NewSignInRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	authResp, err := s.api.SignInAPI(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, authResp)
}

func (s *Server) DealerVerfiyOTP(c *gin.Context) {
	body := models.NewOTPRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	resp, err := s.api.DealerVerifyOTPAPI(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) RenewToken(c *gin.Context) {
	body := models.NewRefreshAccessTokenRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	authresp, err := s.api.RenewTokenAPI(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": authresp})
}

func (s *Server) DealerForgotPassword(c *gin.Context) {
	body := models.NewDealerforgotpasswordRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	resp, err := s.api.DealerForgotPasswordAPI(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message:": resp})
}

func (s *Server) DealerForgotConfirm(c *gin.Context) {
	body := models.NewDealerForgotConfirmRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	err := s.api.DealerForgotConfirmAPI(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message:": "password change successfully"})
}

func (s *Server) AddNewVechicle(c *gin.Context) {
	body := models.NewVehicleRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	vehicle := models.NewVehicle(body.VehicleName, body.VinNumber, body.LicPlate, body.BrandMake, body.Model, body.ManfactureYear, body.VechicleColor, body.TransmissionStyle, nil)
	vehResp, err := s.api.AddNewVechicleAPI(c, vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, vehResp)
}

func (s *Server) DealerSignOut(c *gin.Context) {
	body := models.NewDealerSignOutRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	err := s.api.DealerSignOutAPI(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message:": "signout successfully"})
}

func (s *Server) CreateKey(c *gin.Context) {
	body := models.NewCreateKeyRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	key := models.NewKey(body.VehID, body.KeyName, body.KeyAccessible, body.ReceivedFrom, body.ReceivedBy, body.LocName)
	keyApi, err := s.api.CreateKeyAPI(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	resp := models.CreateKeyResponse{
		Key: keyApi,
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) AddMultipleVehicles(c *gin.Context) {
	body := models.NewAddMultipleVehiclesRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	vehiclesApi, err := s.api.AddMultipleVehiclesApi(c, body.Vehicles)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	vehResp := models.NewMultipleVehiclesResponse(vehiclesApi)
	c.JSON(http.StatusOK, vehResp)
}

func (s *Server) StaffSignup(c *gin.Context) {
	body := models.NewStaffSignupRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	err := lib.StaffRoleValidation(*body.StaffRoleName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	staff := models.NewStaff(body.EmailWork, body.FirstName, body.LastName, body.Username, body.Password, body.PhoneNumber, body.EmailPersonal, body.StaffRoleName, body.StaffInfo, body.ProfileImage, body.Region, body.City, body.State, body.PostalCode, body.Street1Address, body.Street2Address, body.Currency, body.MultiLogin)
	staffResp, err := s.api.StaffSignupAPI(c, staff)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	resp := models.NewStaffSignupResponse(staffResp.StaffId, staffResp.StaffRoleName, staffResp.DealerID, staffResp.EmailWork, staffResp.FirstName, staffResp.LastName, staffResp.Username, staffResp.PhoneNumber, staffResp.EmailPersonal, staffResp.StaffInfo, staffResp.ProfileImage, staffResp.Region, staffResp.City, staffResp.State, staffResp.PostalCode, staffResp.Street1Address, staffResp.Street2Address, staffResp.Currency, staffResp.MaximumKeyStaff, staffResp.MaximumRoleKeys, staffResp.StaffNumber, staffResp.MultiLogin)
	c.JSON(http.StatusOK, resp)
}

func (s *Server) AssignUnassignKey(c *gin.Context) {
	body := models.NewAssignUnassignKeyResquest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	assignUnassignKey := models.NewAssignUnassignKey(body.VehicleID, body.KeyID, body.StaffID, body.KeyAssignStatus)
	assignUnassignResp, err := s.api.AssignUnassignKeyAPI(c, assignUnassignKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp := models.NewAssignUnassignKeyResponse(assignUnassignResp.KeyAssignID, assignUnassignResp.VehicleID, assignUnassignResp.KeyID, assignUnassignResp.StaffID, assignUnassignResp.KeyAssignStatus, assignUnassignResp.TransferID, assignUnassignResp.TransfereID, assignUnassignResp.AssignComment, assignUnassignResp.DataKeyAssignStart, assignUnassignResp.DataKeyAssignStop)
	c.JSON(http.StatusOK, resp)
}

func (s *Server) StaffResetPassword(c *gin.Context) {
	body := models.NewStaffResetPasswordRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	staff := models.NewStaff(nil, nil, nil, nil, body.Password, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	err := s.api.StaffResetPasswordAPI(c, staff)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.ResetPasswordResponse{
		Message: "password reset successfully",
	})
}

func (s *Server) CsvToJson(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	if len(form.File["file"]) > 0 {
		file := form.File["file"][0]
		response, err := s.api.CsvToJsonAPI(c, file)
		if err != nil || response == nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusBadRequest, models.ErrorResponse{
		Error: "please attach the file to the request",
	})
}

func (s *Server) S3UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	if len(form.File["file"]) > 0 {
		file := form.File["file"][0]

		url, err := s.api.S3UploadFileAPI(c, file)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		response := models.S3FileUploadResponse{
			URL: url,
		}
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusBadRequest, models.ErrorResponse{
		Error: "please attach the file to the request",
	})
}

func (s *Server) AddVehicleLocation(c *gin.Context) {
	body := models.NewVehicleLocationRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
	}
	location := models.NewVehicleLocation(nil, body.LocID, body.VehID, body.SpaceID, body.UserIN, body.UserOUT, body.DateInLoc, body.DateOutLoc)
	locResp, err := s.api.VechicleLocationAPI(c, location)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	resp := models.NewVehicleLocationResponse(locResp.VehLocID, locResp.LocID, locResp.VehID, locResp.SpaceID, locResp.UserIN, locResp.UserOUT, locResp.DateInLoc, locResp.DateOutLoc)
	c.JSON(http.StatusOK, resp)
}

func (s *Server) ResendOTP(c *gin.Context) {
	body := models.NewResendOTPRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resendOTPEmail, err := s.api.ResendOTPAPI(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	response := models.ResendOTPResponse{
		Email: resendOTPEmail,
	}
	c.JSON(http.StatusOK, response)
}

func (s *Server) CreateDealerMap(c *gin.Context) {
	body := models.NewMapRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	mapsApi, err := s.api.AddMapAPI(c, body.Maps)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	response := models.NewMapResponse(mapsApi)
	c.JSON(http.StatusOK, response)
}
