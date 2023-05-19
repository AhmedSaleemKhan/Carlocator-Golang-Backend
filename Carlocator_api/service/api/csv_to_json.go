package api

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) CsvToJsonAPI(c *gin.Context, requestFile *multipart.FileHeader) ([]*models.Vehicle, error) {

	src, err := requestFile.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	reader := csv.NewReader(src)
	csvData, err := reader.ReadAll()
	if err != nil {
		return nil, err
	} else if csvData == nil {
		return nil, fmt.Errorf("null file attacted")
	}

	var vehicles []*models.Vehicle
	for _, data := range csvData {
		VehicleName := string(data[0])
		VinNumber := data[2]
		LicPlate := data[3]
		BrandMake := data[4]
		ManfactureYear := data[5]
		Model := data[6]
		VechicleColor := data[7]
		TransmissionStyle := data[8]

		vehicle := models.NewVehicle(&VehicleName, &VinNumber, &LicPlate, &BrandMake, &Model, &ManfactureYear, &VechicleColor, &TransmissionStyle, nil)
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}
