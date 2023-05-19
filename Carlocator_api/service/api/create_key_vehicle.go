package api

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) CreateKeyAPI(c *gin.Context, Key *models.Key) (*models.Key, error) {
	*Key.KeyID = CreateUUID()
	*Key.KeyAvailable = true
	dealerIDDB, err := api.db.GetVehicleDealerIDForKeyDB(c, *Key.VehID)
	dealerID := c.MustGet("user-id").(string)

	if *dealerIDDB != dealerID || err != nil {
		return nil, fmt.Errorf("not allowed for this vehicle")
	}
	assignedKeys, err := api.db.GetVehicleKeysDB(c, *Key.VehID)
	if len(assignedKeys) >= 3 || err != nil {
		return nil, fmt.Errorf("not allowed to create more than 3 keys for a vehicle")
	}

	keyDB, err := api.db.CreateKeyDB(c, Key)
	if err != nil {
		return nil, err
	}
	return keyDB, nil
}
