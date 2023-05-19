package api

import (
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) UpdateDealershipAPI(c *gin.Context, dealerShip *models.Dealer) (*models.Dealer, error) {
	dealerID := c.MustGet("user-id").(string)
	*dealerShip.DealerId = dealerID
	*dealerShip.UpdatedAt = api.clock.NowUnix()

	dealerShipDB, err := api.db.UpdateDealershipDB(c, dealerShip, dealerID)
	if err != nil {
		return nil, err
	}
	return dealerShipDB, nil

}
