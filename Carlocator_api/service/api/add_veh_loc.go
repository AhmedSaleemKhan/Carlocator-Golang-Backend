package api

import (
	//"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) VechicleLocationAPI(c *gin.Context, location *models.VehicleLocation) (*models.VehicleLocation, error) {
	*location.VehLocID = CreateUUID()
	*location.SpaceID = CreateUUID()
	locResp, err := api.db.VechicleLocationDB(c, location)
	if err != nil {
		return nil, err
	}
	return locResp, nil
}
