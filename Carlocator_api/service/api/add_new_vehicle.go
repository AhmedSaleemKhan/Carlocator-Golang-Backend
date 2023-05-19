package api

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func (api *CarLocatorApiImpl) AddNewVechicleAPI(c *gin.Context, vehicle *models.Vehicle) (*models.VehicleResponse, error) {
	userID := c.MustGet("user-id").(string)
	*vehicle.DealerID = userID
	*vehicle.VehicleID = CreateUUID()
	*vehicle.CarStatusID = CreateUUID()
	*vehicle.VechicleTypeID = CreateUUID()
	*vehicle.AttachId = CreateUUID()
	*vehicle.CreatedAt = api.clock.NowUnix()
	lastestStockNumberDB, err := api.db.GetNoOfVehicles(c)
	if err != nil {
		return nil, err
	}
	*vehicle.StockNumber = *lastestStockNumberDB + 1
	vinNumbersDB, err := api.db.GetDealerVinNumbers(c, userID)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	for _, vinNumbers := range vinNumbersDB {
		if vinNumbers == *vehicle.VinNumber {
			return nil, fmt.Errorf("vin number already exist for dealer-id")
		}
	}
	vehResp, err := api.db.AddNewVechicleDB(c, vehicle)
	if err != nil {
		return nil, err
	}
	return vehResp, nil
}
