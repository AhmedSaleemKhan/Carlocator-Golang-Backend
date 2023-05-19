package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) VehicleVinNumberAPI(c *gin.Context, vinNumber string) (*models.VehicleVinNumberResponse, error) {
	url := fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/decodevinextended/%v?format=JSON", vinNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	vehicleResponseModel := models.NHTSAVinDecodeModel{}
	err = json.Unmarshal(body, &vehicleResponseModel)
	if err != nil {
		return nil, err
	}
	response := models.NewVehicleNumberResponse()
	response.VinNumber = &vehicleResponseModel.SearchCriteria
	response.Make = &vehicleResponseModel.Results[7].Value
	response.ManufacturerName = &vehicleResponseModel.Results[8].Value
	response.Model = &vehicleResponseModel.Results[9].Value
	response.ModelYear = &vehicleResponseModel.Results[10].Value
	response.ManufacturedCity = &vehicleResponseModel.Results[11].Value
	response.VehicleType = &vehicleResponseModel.Results[14].Value
	response.ManufacturedCountry = &vehicleResponseModel.Results[15].Value
	response.BodyType = &vehicleResponseModel.Results[23].Value
	response.TransmissionStyle = &vehicleResponseModel.Results[49].Value
	response.FuelType = &vehicleResponseModel.Results[80].Value
	return response, nil
}

