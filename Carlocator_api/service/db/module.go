package db

import (
	"fmt"
	"log"
	"os"

	"capregsoft.com/carlocator/config"
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CarLocatorDB interface {
	AddNewVechicleDB(c *gin.Context, vehicle *models.Vehicle) (*models.VehicleResponse, error)
	DealerSignupDB(c *gin.Context, dealer *models.Dealer) (*models.DealerResponse, error)
	DealerVerifyOTPDB(c *gin.Context, dealer *models.Dealer) (*models.VerifyOTPResponse, error)
	GetStaffDB(c *gin.Context, staff *models.Staff) (*models.Staff, error)
	GetStaffDealerName(c *gin.Context) (string, error)
	GetLatestStaffNumber(c *gin.Context) (*string, error)
	StaffSignupDB(c *gin.Context, staff *models.Staff) (*models.Staff, error)
	GetDealerDB(c *gin.Context, dealer *models.Dealer) (*models.Dealer, error)
	GetVehicleIDForKeyDB(c *gin.Context, keyID string) (*string, error)
	GetVehicleDealerIDForKeyDB(c *gin.Context, vehID string) (*string, error)
	GetStaffByIDDB(c *gin.Context, staff *models.Staff) (*models.Staff, error)
	GetVehicleDealerIDByVehIdDB(c *gin.Context, vehID string) (*string, error)
	//modules added for check in and check out the vehicle from map
	GetDealerIDFromSlotID(c *gin.Context, mapID string) (*models.CheckInOutVehicleMap, error)
	CheckOutVehicleDB(c *gin.Context, checkInOutvehicle *models.CheckInOutVehicleMap) (*models.CheckInOutVehicleMap, error)
	CheckInVehicleDB(c *gin.Context, checkInOutvehicle *models.CheckInOutVehicleMap) (*models.CheckInOutVehicleMap, error)
	GetVehAssignStatus(c *gin.Context, vehicleID string) (*bool, error)
	GetSlotAvailableStatus(c *gin.Context, slotID string) (*bool, error)
	AddNewVehicleToMapDB(c *gin.Context, vehicleForMap *models.AddingVehicleToMap) (*models.AddingVehicleToMap, error)
}
type CarLocatorDBImpl struct {
	conn *sqlx.DB
}

func NewCarLocatorDBImpl() *CarLocatorDBImpl {
	cfg := config.Cfg
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cfg.DbUser, cfg.DbPassword, cfg.Host, cfg.Port, cfg.DbName, cfg.SslMode)
	log.Println(psqlInfo)
	conn, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = conn.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("Connected")
	}
	return &CarLocatorDBImpl{
		conn: conn,
	}
}

var _ CarLocatorDB = &CarLocatorDBImpl{}
