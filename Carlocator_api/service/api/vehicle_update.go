package api

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) UpdatedVehicleAPI(c *gin.Context, vehID string, vehicle *models.Vehicle) (*models.Vehicle, error) {
	dealerID := c.MustGet("user-id").(string)
	*vehicle.DealerID = dealerID

	UpdatedVehicleDB, err := api.db.UpdatedVehicleDB(c, vehID, vehicle)
	if err != nil {
		return nil, err
	}

	return UpdatedVehicleDB, nil

}
