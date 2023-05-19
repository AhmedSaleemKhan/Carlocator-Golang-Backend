package api

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) UpdatedKeyAPI(c *gin.Context, KeyID string, key *models.Key) (*models.Key, error) {

	vehID, err := api.db.GetVehicleIDForKeyDB(c, KeyID)
	if err != nil {
		return nil, err
	}
	dealerIDDB, err := api.db.GetVehicleDealerIDForKeyDB(c, *vehID)
	dealerID := c.MustGet("user-id").(string)

	if *dealerIDDB != dealerID || err != nil {
		return nil, fmt.Errorf("not allowed for this vehicle")
	}
	updatedKeys, err := api.db.UpdatedKeyDB(c, KeyID, key)

	if err != nil {
		return nil, err
	}
	return updatedKeys, err
}
