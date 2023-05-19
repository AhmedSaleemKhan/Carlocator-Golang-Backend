package api

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) AddVehicleToMapAPI(c *gin.Context, vehicleMap *models.AddingVehicleToMap) (*models.AddingVehicleToMap, error) {

	// check the vehicle already exist at slot or not
	slotData, err := api.db.CheckVehExistsInSlotDB(c, *vehicleMap.VehicleID)
	if err != nil {
		return nil, err
	}
	if len(*slotData.VehicleID) > 0 {
		return nil, fmt.Errorf("veh id %v already exists at slot id %v", *slotData.VehicleID, *slotData.SlotID)
	}
	//get the slot available status if alredy occcupied then genreate an error
	slotAvailableStatus, err := api.db.GetSlotAvailableStatus(c, *vehicleMap.SlotID)
	if err != nil || *slotAvailableStatus != true {
		return nil, fmt.Errorf("slot is not available %v ", *vehicleMap.SlotID)
	}
	//else add vehicle to that slot
	addVehicleToMapDB, err := api.db.AddNewVehicleToMapDB(c, vehicleMap)
	if err != nil {
		return nil, err
	}
	return addVehicleToMapDB, nil

}
