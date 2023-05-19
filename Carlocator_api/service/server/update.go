package server

import (
	"net/http"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateDealership(c *gin.Context) {
	body := models.NewUpdateDealerShipRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	updatedDealerShip, err := s.api.UpdateDealershipAPI(c, body.Dealer)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	if updatedDealerShip.Region == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "dealer not found",
		})
		return
	}
	resp := models.NewUpdateDealershipResponse(updatedDealerShip)

	c.JSON(http.StatusOK, resp)
}

func (s *Server) UpdatedVehicle(c *gin.Context) {
	vehicleID := c.Param("veh_id")
	body := models.NewUpdatedVehicleRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	updatedVehicle, err := s.api.UpdatedVehicleAPI(c, vehicleID, body.Vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	response := models.NewUpdatedVehicleResponse(updatedVehicle)
	c.JSON(http.StatusOK, response)
}

func (s *Server) UpdateKey(c *gin.Context) {
	keyID := c.Param("key_id")
	body := models.NewUpdatedKeyRequest()
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{

			Error: err.Error(),
		})
		return
	}
	UpdatedKey, err := s.api.UpdatedKeyAPI(c, keyID, body.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	resp := models.NewUpdatedKeyResponse(UpdatedKey)
	c.JSON(http.StatusOK, resp)
}

func (s *Server) CheckInOutVehicle(c *gin.Context) {
	body := models.NewCheckInOutVehicleRequest() //check-in out vehicle request
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	checkInOutResp, err := s.api.CheckInOutVehicleAPI(c, body.CheckInOutVehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	response := models.NewCheckInOutVehicleResponse(checkInOutResp)
	c.JSON(http.StatusOK, response)
}

func (s *Server) AddVehicleToMap(c *gin.Context) {
	body := models.NewAddingVehicleToMapRequest() //check-in out vehicle request
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	addVehicleMapAPI, err := s.api.AddVehicleToMapAPI(c, body.AddVehicleToMap)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	vehicleMapResponse := models.NewAddingVehicleToMapResponse(addVehicleMapAPI)
	c.JSON(http.StatusOK, vehicleMapResponse)
}
