package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"capregsoft.com/carlocator/service/models"

	"github.com/gin-gonic/gin"
)

func (api *CarLocatorApiImpl) MultipleVinNumberAPI(c *gin.Context, vinNumbers []string) (map[string]models.VinNumberDecode, error) {
	results := make(map[string]models.VinNumberDecode, 0)
	for _, vinNumber := range vinNumbers {
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

		vinNumberDecode := models.VinNumberDecode{}
		vinNumberDecode.VinNumber = &vehicleResponseModel.SearchCriteria
		vinNumberDecode.Make = &vehicleResponseModel.Results[6].Value
		vinNumberDecode.ManufacturerName = &vehicleResponseModel.Results[7].Value
		vinNumberDecode.Model = &vehicleResponseModel.Results[8].Value
		vinNumberDecode.ModelYear = &vehicleResponseModel.Results[9].Value
		vinNumberDecode.ManufacturedCity = &vehicleResponseModel.Results[10].Value
		vinNumberDecode.VehicleType = &vehicleResponseModel.Results[13].Value
		vinNumberDecode.ManufacturedCountry = &vehicleResponseModel.Results[14].Value
		vinNumberDecode.BodyType = &vehicleResponseModel.Results[22].Value
		vinNumberDecode.TransmissionStyle = &vehicleResponseModel.Results[48].Value
		results[vinNumber] = vinNumberDecode
	}

	return results, nil
}
