package api

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api CarLocatorApiImpl) GetUserAssignedKeysApi(c *gin.Context, filter models.GetAssignedKeyFilter) ([]*models.AssignedKeys, *models.Metadata, error) {
	dealerId := c.MustGet("user-id").(string)
	assignedKeysDb, metadata, err := api.db.GetUserAssignedKeysDB(c, dealerId, filter)
	if err != nil {
		return nil, nil, err
	}
	return assignedKeysDb, metadata, err
}
