package server

import (
	//"crypto/sha1"

	"capregsoft.com/carlocator/service/api"
	"capregsoft.com/carlocator/service/lib"
	"capregsoft.com/carlocator/service/middlewares"
	"github.com/gin-gonic/gin"
)

type ServerImpl interface {
	DealerSignup(c *gin.Context)
	SignIn(c *gin.Context)
	DealerVerfiyOTP(c *gin.Context)
	RenewToken(c *gin.Context)
	DealerForgotPassword(c *gin.Context)
	DealerForgotConfirm(c *gin.Context)
	AddNewVechicle(c *gin.Context)
	DealerSignOut(c *gin.Context)
	VehicleVinNumber(c *gin.Context)
	MultipleVinNumber(c *gin.Context)
	CreateKey(c *gin.Context)
	GetDealerVehicles(c *gin.Context)
	AddMultipleVehicles(c *gin.Context)
	DeleteKey(c *gin.Context)
	StaffSignup(c *gin.Context)
	AssignUnassignKey(c *gin.Context)
	GetVehicleKeys(c *gin.Context)
	StaffResetPassword(c *gin.Context)
	UpdateDealership(c *gin.Context)
	DeleteVehicle(c *gin.Context)
	GetStaffMembers(c *gin.Context)
	CsvToJson(c *gin.Context)
	GetDealerInfo(c *gin.Context)
	UserAssignedKeys(c *gin.Context)
	AddVehicleLocation(c *gin.Context)
	S3UploadFile(c *gin.Context)
	GetAllDealers(c *gin.Context)
	UpdatedVehicle(c *gin.Context)
	UpdateKey(c *gin.Context)
	ResendOTP(c *gin.Context)
	CreateDealerMap(c *gin.Context)
	GetDealerMap(c *gin.Context)
	GetStaffAssignedVehicles(c *gin.Context)
	CheckInOutVehicle(c *gin.Context)
	AddVehicleToMap(c *gin.Context)
}
type Server struct {
	api *api.CarLocatorApiImpl
}

func NewServer() *Server {
	return &Server{
		api: api.NewCarLocatorApiImpl(),
	}
}

func NewServerImpl(r *gin.Engine) *gin.Engine {
	server := NewServer()
	lib.NewCustomValidator()
	dealerGroup := r.Group("/dealer")
	{
		dealerGroup.POST("/SignIn", server.SignIn)
		dealerGroup.POST("/signup", server.DealerSignup)
		dealerGroup.POST("/verifyotp", server.DealerVerfiyOTP)
		dealerGroup.POST("/forgotpassword", server.DealerForgotPassword)
		dealerGroup.POST("/forgotconfirm", server.DealerForgotConfirm)
		dealerGroup.POST("/signout", server.DealerSignOut)
		dealerGroup.GET("/info", middlewares.Auth(), server.GetDealerInfo)
		dealerGroup.PATCH("/update", middlewares.Auth(), server.UpdateDealership)
		dealerGroup.POST("/s3/upload", middlewares.Auth(), server.S3UploadFile)
		dealerGroup.GET("/getalldealer", server.GetAllDealers)
		dealerGroup.POST("/resend/otp", server.ResendOTP)
		staffGroup := dealerGroup.Group("/staff")
		{
			staffGroup.POST("/signup", middlewares.Auth(), server.StaffSignup)
			staffGroup.POST("resetpassword", middlewares.Auth(), server.StaffResetPassword)
			staffGroup.GET("/getdealerstaff", middlewares.Auth(), server.GetStaffMembers)
		}
		mapGroup := dealerGroup.Group("/map")
		{
			mapGroup.POST("/create", middlewares.Auth(), server.CreateDealerMap)
			mapGroup.GET("/get", middlewares.Auth(), server.GetDealerMap)
		}
	}
	vehicleGroup := r.Group("/vehicle")
	{
		vehicleGroup.GET("/decodevin/:vin_number", server.VehicleVinNumber)
		vehicleGroup.POST("/add", middlewares.Auth(), server.AddNewVechicle)
		vehicleGroup.POST("/decodevin/multiple_vin", server.MultipleVinNumber)
		vehicleGroup.GET("/all", middlewares.Auth(), server.GetDealerVehicles)
		vehicleGroup.POST("/multiple", middlewares.Auth(), server.AddMultipleVehicles)
		vehicleGroup.POST("/assingunassignkey", middlewares.Auth(), server.AssignUnassignKey)
		vehicleGroup.POST("/createkey", middlewares.Auth(), server.CreateKey)
		vehicleGroup.DELETE("/key/:key_id", middlewares.Auth(), server.DeleteKey)
		vehicleGroup.GET("/keys/:veh_id", middlewares.Auth(), server.GetVehicleKeys)
		vehicleGroup.POST("/csvtojson", middlewares.Auth(), server.CsvToJson)
		vehicleGroup.DELETE("/:veh_id", middlewares.Auth(), server.DeleteVehicle)
		vehicleGroup.GET("/assignedkeys", middlewares.Auth(), server.UserAssignedKeys)
		vehicleGroup.POST("/veh_loc", middlewares.Auth(), server.AddVehicleLocation)
		vehicleGroup.PATCH("/update/:veh_id", middlewares.Auth(), server.UpdatedVehicle)
		vehicleGroup.PATCH("/key/update/:key_id", middlewares.Auth(), server.UpdateKey)
		vehicleGroup.GET("/assigned", middlewares.Auth(), server.GetStaffAssignedVehicles)
		//added in vehicle-group for checking in and checking out from map slots...
		vehicleGroup.POST("/checkinout", middlewares.Auth(), server.CheckInOutVehicle)
		//new enpoint for adding new vehicle in to the map for ddealer....
		vehicleGroup.POST("/add/map", middlewares.Auth(), server.AddVehicleToMap)
	}
	r.POST("/renewtoken", server.RenewToken)
	r.POST("/signin", server.SignIn)

	return r
}

var _ ServerImpl = &Server{}
