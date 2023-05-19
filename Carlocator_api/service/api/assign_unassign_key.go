package api

import (
	"fmt"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func (api *CarLocatorApiImpl) AssignUnassignKeyAPI(c *gin.Context, keyData *models.AssignUnassignKey) (*models.AssignUnassignKey, error) {
	dealerID := c.MustGet("user-id").(string)
	staffDealerID, err := api.db.GetDealerIDForKeyDB(c, *keyData.StaffID)
	// If staff not belongs to dealer it is not allowed to assing key for staff
	if err != nil || *staffDealerID != dealerID {
		return nil, fmt.Errorf("staff is not yours ")
	}
	vehicleDealerID, err := api.db.GetVehicleDealerIDForKeyDB(c, *keyData.VehicleID)
	// If vehicle not belongs to dealer it is not allowed for dealer to assing vehicle key to staff
	if err != nil || *vehicleDealerID != dealerID {
		return nil, fmt.Errorf("vehicle is not yours ")
	}
	vehicleID, err := api.db.GetVehicleIDForKeyDB(c, *keyData.KeyID)
	// If vehicle against key does not match with vehicle that pass through body then it is not allowed to assign key of that vehicle
	if err != nil || *keyData.VehicleID != *vehicleID {
		return nil, fmt.Errorf("not allowed for this vehicle")
	}
	staffMaxRoleKeys, err := api.db.GetStaffRolekeysDB(c, *keyData.StaffID)
	if *staffMaxRoleKeys <= 0 || err != nil {
		return nil, fmt.Errorf("not allowed for this staff")
	}
	maxStaffkey, err := api.db.GetMaxkeystaffDB(c, *keyData.StaffID)
	if err != nil {
		return nil, fmt.Errorf("not allowed for this staff")
	}
	var assignUnAssignDB *models.AssignUnassignKey
	var keyStatus bool
	var keyCount int64
	*keyData.KeyAssignID = CreateUUID()
	*keyData.TransferID = dealerID
	*keyData.TransfereID = *keyData.StaffID
	// Key Assignment
	if *keyData.KeyAssignStatus {
		*keyData.DataKeyAssignStart = api.clock.NowUnix()
		// Get key status
		keyAvailable, err := api.db.GetkeyStatusDB(c, *keyData.KeyID)
		if err != nil {
			return nil, err
		}
		// If keyStatus is not true then not allowed to assign key
		if !*keyAvailable {
			return nil, fmt.Errorf("key is already assign")
		}
		// If already assign is equal to the staff assign key limits then no more keys assign to staff
		if *maxStaffkey == *staffMaxRoleKeys {
			return nil, fmt.Errorf("no more limit to assing key")
		}
		vehStaffID, err := api.db.GetStaffFromAssignKeyByVehIDDB(c, *keyData.VehicleID)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
		if vehStaffID == nil {
			// When key assign add requried feilds in database
			assignUnAssignDB, err = api.db.AddVechicleKeyAssignDB(c, keyData)
			if err != nil {
				return nil, err
			}
			// If key assign successfully then add assign key in staff table
			keyCount = *maxStaffkey + 1
			// When key Assigned  update status of key and vehicle not available
			keyStatus = false
		} else {
			keyData.DataKeyAssignStop = nil

			assignUnAssignDB, err = api.db.UpdateVechicleKeyAssignDB(c, keyData)
			if err != nil {
				return nil, err
			}
			keyCount = *maxStaffkey + 1
			// When key Assigned  update status of key and vehicle not available
			keyStatus = false
		}

	} else { //Key UnAssignment
		*keyData.DataKeyAssignStop = api.clock.NowUnix()
		assignUnAssignDB, err = api.db.AddVechicleKeyUnassignDB(c, keyData)
		if err != nil {
			return nil, err
		}
		// If key unassign successfully then remove assign key in staff table
		keyCount = *maxStaffkey - 1
		// When key unAssigned  update status of key and vehicle available
		keyStatus = true
	}

	err = api.db.UpdateMaxStaffKeyDB(c, keyCount, *keyData.StaffID)
	if err != nil {
		return nil, err
	}
	err = api.db.UpdateVehicleKeyStatusDB(c, keyStatus, *keyData.KeyID, *keyData.VehicleID)
	if err != nil {
		return nil, err
	}

	return assignUnAssignDB, nil
}
