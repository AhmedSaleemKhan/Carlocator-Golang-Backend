package api

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) CheckInOutVehicleAPI(c *gin.Context, checkInOutVehicle *models.CheckInOutVehicleMap) (*models.CheckInOutVehicleMap, error) {
	staffID := c.MustGet("user-id").(string)

	//get the dealer-id from slot-id in order to validate from findr.map table
	mapData, err := api.db.GetDealerIDFromSlotID(c, *checkInOutVehicle.SlotID)
	if err != nil {
		return nil, err
	}
	DealerIDDB, err := api.db.GetDealerStaffID(c, staffID)
	if err != nil {
		return nil, err
	}
	//if dealerID does not belongs to the dealerid of map id generate an error
	if *DealerIDDB != mapData.DealerID {
		return nil, fmt.Errorf("map-id does not belong to your dealer ")
	}
	//this function is supposed to get the dealer-id from the veh--id in order to validate vehicle belongs to token dealer or not
	vehicleDealerID, err := api.db.GetVehicleDealerIDForKeyDB(c, *checkInOutVehicle.VehicleID)
	// If vehicle not belongs to dealer it is not allowed for dealer to assing vehicle key to staff
	if *vehicleDealerID != *DealerIDDB || err != nil {
		return nil, fmt.Errorf("vehicle is not yours")
	}
	//set the dealer-id as check-in-out-id for time being
	checkInOutVehicle.DealerID = *DealerIDDB
	checkInOutVehicle.SlotAvailable = mapData.SlotAvailable

	var checkInOutVehicleDB *models.CheckInOutVehicleMap
	//validation if the vehicle is not assigned to specific staff return error
	vehAssigStatus, err := api.db.GetVehAssignStatus(c, *checkInOutVehicle.VehicleID)
	if *vehAssigStatus != true {
		return nil, fmt.Errorf("vehicle is not assigned for staff")
	}
	//if the slot-available status==true means wants to check in the vehicle

	if *checkInOutVehicle.SlotAvailable {
		// check the vehicle already exist at slot or not
		slotData, err := api.db.CheckVehExistsInSlotDB(c, *checkInOutVehicle.VehicleID)
		if err != nil {
			return nil, err
		}
		if len(*slotData.VehicleID) > 0 {
			return nil, fmt.Errorf("vehicle already exists at slot")
		}

		//check-in the vehicle and update the findr.map_slot table for specfic slot-id
		checkInOutVehicleDB, err = api.db.CheckInVehicleDB(c, checkInOutVehicle)
		if err != nil {
			return nil, err
		}
	} else {
		if *checkInOutVehicle.VehicleID == *mapData.VehicleID {
			// check-out the vehicle and update the findr.map_slot table for the specific slot-id
			checkInOutVehicleDB, err = api.db.CheckOutVehicleDB(c, checkInOutVehicle)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("slot is not available ")
		}
	}
	return checkInOutVehicleDB, nil
}
