package api

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) GetStaffVehiclesAPI(c *gin.Context, staffID string) ([]*models.StaffAssignedVehicleKeys, error) {
	var vehDB []*models.StaffAssignedVehicleKeys

	vehDB, err := api.db.GetDealerStaffIDForVehDB(c, staffID)

	if err != nil {
		return nil, err
	}

	return vehDB, nil
}
