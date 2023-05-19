package api

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) AddMultipleVehiclesApi(c *gin.Context, vehicles []*models.Vehicle) ([]*models.Vehicle, error) {

	var err error
	for _, vehicle := range vehicles {
		vehicle.DealerID = new(string)
		vehicle.VehicleID = new(string)
		vehicle.CarStatusID = new(string)
		vehicle.CreatedAt = new(int64)
		vehicle.VechicleTypeID = new(string)
		vehicle.AttachId = new(string)
		vehicle.MapColour = new(string)
		vehicle.Note = new(string)
		vehicle.StockNumber = new(int64)

		lastestStockNumberDB, err := api.db.GetNoOfVehicles(c)
		if err != nil {
			return nil, err
		}
		*vehicle.StockNumber = *lastestStockNumberDB + 1
		*vehicle.DealerID = c.MustGet("user-id").(string)
		*vehicle.VehicleID = CreateUUID()
		*vehicle.CarStatusID = CreateUUID()
		*vehicle.CreatedAt = api.clock.NowUnix()
		*vehicle.AttachId = CreateUUID()
		*vehicle.VechicleTypeID = CreateUUID()

		_, DBerr := api.db.AddNewVechicleDB(c, vehicle)
		if DBerr != nil {
			return nil, err
		}
	}
	return vehicles, err
}
