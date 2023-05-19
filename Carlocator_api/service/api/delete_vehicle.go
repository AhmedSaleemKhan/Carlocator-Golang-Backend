package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) DeleteVehicleAPI(c *gin.Context, vehicleId string) error {
	dealerId := c.MustGet("user-id").(string)
	vehDealerId, err := api.db.GetVehicleDealerIDByVehIdDB(c, vehicleId)
	if *vehDealerId != dealerId || err != nil {
		return fmt.Errorf("not allowed to delete this vehicle")
	}
	err = api.db.DeleteVehicleDB(c, vehicleId)
	if err != nil {
		return err
	}
	return nil
}
