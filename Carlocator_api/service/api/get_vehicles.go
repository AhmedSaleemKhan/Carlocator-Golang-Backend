package api

import (
	"encoding/json"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func (api *CarLocatorApiImpl) GetVehicles(c *gin.Context, filter models.GetAllVehiclesFilter) ([]*models.Vehicle, *models.Metadata, *models.UniqueVehicleAttributes, error) {
	vehiclesApi := make([]*models.Vehicle, 0)
	vehiclesDB, metadata, err := api.db.GetDealerVechiclesDB(c, filter)
	if err != nil {
		return nil, nil, nil, err
	}
	var allVehicles *models.Vehicle

	//validating if keys bool==true show keys..
	if filter.Keys {
		for _, allVehicles = range vehiclesDB {
			keysDB, err := api.db.GetVehicleKeysDB(c, *allVehicles.VehicleID)
			if err != nil {
				return nil, nil, nil, err
			} else if filter.Staff { //if staff filter == true show staff info else --appened keys
				var staffDB *models.Staff
				for _, staffApi := range keysDB {
					staffDB, err = api.db.GetStaffByKeyId(c, *staffApi.KeyID)
					if err != nil {
						glg.Error(err)
						return nil, nil, nil, err
					}
					if staffDB != nil {
						staffApi.Staff = staffDB
					}
				}
			}
			allVehicles.Keys = append(allVehicles.Keys, keysDB...)

			vehiclesApi = append(vehiclesApi, allVehicles)
		}
	} else { //else appened all the vehicles and return it ...
		vehiclesApi = append(vehiclesApi, vehiclesDB...)
	}

	uniqueAttributes, err1 := api.db.GetDealersVehiclesAttributesDB(c)
	if err != nil {
		return nil, nil, nil, err1
	}
	uniqueAttributesParsed := models.UniqueVehicleAttributes{}

	if len(uniqueAttributes) > 0 {
		uniqueAttribute := uniqueAttributes[0]
		var makes, years, models []string
		json.Unmarshal([]byte(*uniqueAttribute.BrandMakes), &makes)
		json.Unmarshal([]byte(*uniqueAttribute.ManfactureYears), &years)
		json.Unmarshal([]byte(*uniqueAttribute.Models), &models)

		uniqueAttributesParsed.BrandMakes = makes
		uniqueAttributesParsed.Models = models
		uniqueAttributesParsed.ManfactureYears = years

	}
	return vehiclesApi, metadata, &uniqueAttributesParsed, nil

}
