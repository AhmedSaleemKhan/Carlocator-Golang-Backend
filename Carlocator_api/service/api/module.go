package api

import (
	"mime/multipart"

	"capregsoft.com/carlocator/pkg/utils"
	"capregsoft.com/carlocator/service/db"
	"capregsoft.com/carlocator/service/models"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

type CarLocatorApi interface {
	DealerSignupAPI(c *gin.Context, dealer *models.Dealer) (*models.DealerResponse, error)
	DealerVerifyOTPAPI(c *gin.Context, body *models.OTPRequest) (*models.VerifyOTPResponse, error)
	DealerForgotPasswordAPI(c *gin.Context, body *models.DealerForgotPasswordRequest) (*cognito.ForgotPasswordOutput, error)
	RenewTokenAPI(c *gin.Context, body *models.RefreshAccessTokenRequest) (*cognito.InitiateAuthOutput, error)
	SignInAPI(c *gin.Context, body *models.SignInRequest) (*cognito.InitiateAuthOutput, error)
	AddNewVechicleAPI(c *gin.Context, vehicle *models.Vehicle) (*models.VehicleResponse, error)
	MultipleVinNumberAPI(c *gin.Context, vinNumbers []string) (map[string]models.VinNumberDecode, error)
	DeleteKeyAPI(c *gin.Context, keyID string) error
	StaffSignupAPI(c *gin.Context, newStaff *models.Staff) (*models.Staff, error)
	StaffResetPasswordAPI(c *gin.Context, staff *models.Staff) error
	GetStaffMembersAPI(c *gin.Context, filter models.GetStaffMembersFilter) ([]*models.Staff, *models.Metadata, error)
	DeleteVehicleAPI(c *gin.Context, vehicleId string) error
	CreateKeyAPI(c *gin.Context, Key *models.Key) (*models.Key, error)
	CsvToJsonAPI(c *gin.Context, requestFile *multipart.FileHeader) ([]*models.Vehicle, error)
	GetDealerMapAPI(c *gin.Context) ([]*models.Map, error)
	//module added for the map API in api module
	CheckInOutVehicleAPI(c *gin.Context, checkInOutVehicle *models.CheckInOutVehicleMap) (*models.CheckInOutVehicleMap, error)
	AddMapAPI(c *gin.Context, maps []*models.Map) ([]*models.Map, error)
	//new module added nd created api for new vehicle to insert in any slot
	AddVehicleToMapAPI(c *gin.Context, vehicleMap *models.AddingVehicleToMap) (*models.AddingVehicleToMap, error)
}

type CarLocatorApiImpl struct {
	db         *db.CarLocatorDBImpl
	awssession *session.Session
	clock      *utils.ClockImpl
}

func NewCarLocatorApiImpl() *CarLocatorApiImpl {
	dbImpl := db.NewCarLocatorDBImpl()
	clockImp := utils.NewClock()

	return &CarLocatorApiImpl{
		db:         dbImpl,
		awssession: ConnectAws(),
		clock:      clockImp,
	}
}

var _ CarLocatorApi = &CarLocatorApiImpl{}
