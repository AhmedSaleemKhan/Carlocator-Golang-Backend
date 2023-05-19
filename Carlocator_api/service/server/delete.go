package server

import (
	"net/http"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteKey(c *gin.Context) {
	keyID := c.Param("key_id")
	isValid := IsValidUUID(keyID)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid key id",
		})
		return
	}
	err := s.api.DeleteKeyAPI(c, keyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp := models.DeleteKeyResponse{
		Message: "key deleted successfully",
	}
	c.JSON(http.StatusOK, resp)

}

func (s *Server) DeleteVehicle(c *gin.Context) {
	vehicleId := c.Param("veh_id")
	isValid := IsValidUUID(vehicleId)
	if !isValid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "invalid vehicle id",
		})
		return
	}
	err := s.api.DeleteVehicleAPI(c, vehicleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	resp := models.DeleteVehicleResponse{
		Message: "vehicle deleted successfully",
	}
	c.JSON(http.StatusOK, resp)
}
