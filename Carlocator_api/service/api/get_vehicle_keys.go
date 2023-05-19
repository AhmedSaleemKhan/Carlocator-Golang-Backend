package api

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) GetVehcileKeysApi(c *gin.Context, vehicleID string) ([]*models.Key, error) {
	var keysDB []*models.Key
	dealerID := c.MustGet("user-id").(string)
	vehDealerID, err := api.db.GetVehicleDealerIDForKeyDB(c, vehicleID)
	if *vehDealerID != dealerID || err != nil {
		return keysDB, fmt.Errorf("not allowed to get keys for this vehicle")
	}
	keysDB, err = api.db.GetVehicleKeysDB(c, vehicleID)
	if err != nil {
		return nil, err
	}
	return keysDB, nil
}
