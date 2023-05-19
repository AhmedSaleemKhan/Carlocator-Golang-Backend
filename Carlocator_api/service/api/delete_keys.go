package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) DeleteKeyAPI(c *gin.Context, keyID string) error {
	vehID, err := api.db.GetVehicleIDForKeyDB(c, keyID)
	if err != nil {
		return err
	}
	dealerID := c.MustGet("user-id").(string)
	vehDealerID, err := api.db.GetVehicleDealerIDForKeyDB(c, *vehID)
	if *vehDealerID != dealerID || err != nil {
		return fmt.Errorf("not allowed to delete this key")
	}
	err = api.db.DeleteKeyDB(c, keyID)
	if err != nil {
		return err
	}
	return nil
}
