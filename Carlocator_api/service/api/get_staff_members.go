package api

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) GetStaffMembersAPI(c *gin.Context, filter models.GetStaffMembersFilter) ([]*models.Staff, *models.Metadata, error) {
	userID := c.MustGet("user-id").(string)
	staffDb, metadata, err := api.db.GetStaffMembersDB(c, userID, filter)
	if err != nil {
		return nil, nil, err
	}

	return staffDb, metadata, nil
}
