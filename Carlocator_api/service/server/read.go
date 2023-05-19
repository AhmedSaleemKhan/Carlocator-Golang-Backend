package server

import (
	"encoding/json"
	"net/http"

	"capregsoft.com/carlocator/service/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) VehicleVinNumber(c *gin.Context) {
	vinNumber := c.Param("vin_number")
	resp, err := s.api.VehicleVinNumberAPI(c, vinNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) MultipleVinNumber(c *gin.Context) {
	body := models.NewMultipleVinNumberRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsAPI, err := s.api.MultipleVinNumberAPI(c, body.VinNumbers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := models.MultipleVinNumberResponse{
		Results: resultsAPI,
	}
	c.JSON(http.StatusOK, response)
}

func (s *Server) GetDealerVehicles(c *gin.Context) {
	filterJson := c.Query("json")

	filter := models.GetAllVehiclesFilter{
		Limit: 10,
		Page:  1,
	}

	if len(filterJson) > 0 {
		if err := json.Unmarshal([]byte(filterJson), &filter); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "invalid filter request passed",
			})
			return
		}
	}

	vehiclesApi, metadata, uniqueAttributes, err := s.api.GetVehicles(c, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	response := models.GetDealerVehiclesResponse{
		Vehicles:   vehiclesApi,
		Atrributes: uniqueAttributes,
		Metadata:   metadata,
	}
	c.JSON(http.StatusOK, response)
}

func (s *Server) GetVehicleKeys(c *gin.Context) {
	vehicleID := c.Param("veh_id")
	keysApi, err := s.api.GetVehcileKeysApi(c, vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := models.GetVehicleKeysRespone{
		Keys: keysApi,
	}
	c.JSON(http.StatusOK, response)
}

func (s *Server) GetStaffMembers(c *gin.Context) {
	filtrateJson := c.Query("json")
	filter := models.GetStaffMembersFilter{
		Limit: 10,
		Page:  1,
	}
	if len(filtrateJson) > 0 {
		if err := json.Unmarshal([]byte(filtrateJson), &filter); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "invalid filter request passed",
			})
			return
		}
	}

	staffMembersApi, metadata, err := s.api.GetStaffMembersAPI(c, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := models.GetStaffMembersResponse{
		StaffMembers: staffMembersApi,
		Metadata:     metadata,
	}
	c.JSON(http.StatusOK, response)

}

func (s *Server) GetDealerInfo(c *gin.Context) {
	dealerInfoApi, err := s.api.GetDealerInfoApi(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := models.GetDealerInfoResponse{
		DealerInfo: dealerInfoApi,
	}
	c.JSON(http.StatusOK, response)
}

func (s *Server) UserAssignedKeys(c *gin.Context) {
	filtratedJson := c.Query("json")
	filter := models.GetAssignedKeyFilter{
		Limit: 10,
		Page:  1,
	}
	if len(filtratedJson) > 0 {
		if err := json.Unmarshal([]byte(filtratedJson), &filter); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "invalid filter request passed",
			})
			return
		}
	}
	assignedKeysApi, metadata, err := s.api.GetUserAssignedKeysApi(c, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	response := models.GetUserAssignedKeysResponse{
		AssignedKeys: assignedKeysApi,
		Metadata:     metadata,
	}
	c.JSON(http.StatusOK, response)
}

func (s *Server) GetAllDealers(c *gin.Context) {
	getalldealersApi, err := s.api.GetAllDealersApi(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := models.GetDealersResponse{
		GetAllDealers: getalldealersApi,
	}
	c.JSON(http.StatusOK, response)

}

func (s *Server) GetDealerMap(c *gin.Context) {
	getMapsAPI, err := s.api.GetDealerMapAPI(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	getMapResponse := models.GetDealerMapResponse{
		DealerMaps: getMapsAPI,
	}
	c.JSON(http.StatusOK, getMapResponse)

}
func (s *Server) GetStaffAssignedVehicles(c *gin.Context) {
	staffID := c.MustGet("user-id").(string)

	vehiclesApi, err := s.api.GetStaffVehiclesAPI(c, staffID)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
	}

	response := models.StaffAssignedVehicleKeysResponse{
		Vehicles: vehiclesApi,
	}

	c.JSON(http.StatusOK, response)

}
