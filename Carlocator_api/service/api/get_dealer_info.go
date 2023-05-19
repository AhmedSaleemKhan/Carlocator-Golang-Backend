package api

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) GetDealerInfoApi(c *gin.Context) ([]*models.Dealer, error) {

	userID := c.MustGet("user-id").(string)
	dealerInfoDb, err := api.db.GetDealerAllInfoDB(c, userID)
	if err != nil {
		return nil, err
	}
	return dealerInfoDb, err
}
