package api

import (
	//"encoding/json"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) GetAllDealersApi(c *gin.Context) ([]*models.GetAllDealers, error) {
	dealerDB, err := api.db.GetAllDealersDB(c)
	if err != nil {
		return nil, err
	}
	return dealerDB, err
}
