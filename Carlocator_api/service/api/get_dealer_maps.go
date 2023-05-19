package api

import (
	"encoding/json"

	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func (api *CarLocatorApiImpl) GetDealerMapAPI(c *gin.Context) ([]*models.Map, error) {
	mapsApi := make([]*models.Map, 0)
	userID := c.MustGet("user-id").(string)
	dealerMapsDB, err := api.db.GetDealerMapDB(c, userID)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	if len(dealerMapsDB) == 0 {
		dealerID, err := api.db.GetDealerStaffID(c, userID)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
		dealerMapsDB, err = api.db.GetDealerMapDB(c, *dealerID)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
	}

	for _, dealerMapDB := range dealerMapsDB {
		if dealerMapDB.ByteCoords != nil {
			coords := models.Coords{}
			err = json.Unmarshal(dealerMapDB.ByteCoords, &coords)
			if err != nil {
				glg.Error(err)
				return nil, err
			}

			mapCoords := make([]*models.Coordinates, 0)
			for _, coords := range coords.Coordinates[0] {
				newCoord := models.Coordinates{}

				for i, coord := range coords {
					if i == 0 {
						newCoord.Latitude = coord
					} else {
						newCoord.Longitude = coord
					}
				}
				mapCoords = append(mapCoords, &newCoord)
			}
			dealerMapDB.Coordinates = mapCoords
			dealerMapDB.ByteCoords = nil
		}
		parkingSlotsDB, err := api.db.GetDealerParkingLot(c, dealerMapDB.MapID)
		if err != nil {
			glg.Error(err)
			return nil, err
		}
		for _, parkingLotDB := range parkingSlotsDB {
			if parkingLotDB.ByteCoords != nil {
				coords := models.Coords{}
				err = json.Unmarshal(parkingLotDB.ByteCoords, &coords)
				if err != nil {
					glg.Error(err)
					return nil, err
				}

				parkingLotCoords := make([]*models.Coordinates, 0)
				for _, coords := range coords.Coordinates[0] {
					newCoord := models.Coordinates{}

					for i, coord := range coords {
						if i == 0 {
							newCoord.Latitude = coord
						} else {
							newCoord.Longitude = coord
						}
					}
					parkingLotCoords = append(parkingLotCoords, &newCoord)
				}
				parkingLotDB.Coordinates = parkingLotCoords
				parkingLotDB.ByteCoords = nil
			}
			dealerZonesDB, err := api.db.GetDealerMapZone(c, parkingLotDB.ParkingLotID)
			if err != nil {
				glg.Error(err)
				return nil, err
			}
			for _, zoneDB := range dealerZonesDB {
				if zoneDB.ByteCoords != nil {
					coords := models.Coords{}
					err = json.Unmarshal(zoneDB.ByteCoords, &coords)
					if err != nil {
						glg.Error(err)
						return nil, err
					}

					zoneCoords := make([]*models.Coordinates, 0)
					for _, coords := range coords.Coordinates[0] {
						newCoord := models.Coordinates{}

						for i, coord := range coords {
							if i == 0 {
								newCoord.Latitude = coord
							} else {
								newCoord.Longitude = coord
							}
						}
						zoneCoords = append(zoneCoords, &newCoord)
					}
					zoneDB.Coordinates = zoneCoords
					zoneDB.ByteCoords = nil
				}
				dealerRowdDB, err := api.db.GetDealerMapRow(c, zoneDB.ZoneID)
				if err != nil {
					glg.Error(err)
					return nil, err
				}

				for _, rowDB := range dealerRowdDB {
					if rowDB.ByteCoords != nil {
						coords := models.Coords{}
						err = json.Unmarshal(rowDB.ByteCoords, &coords)
						if err != nil {
							glg.Error(err)
							return nil, err
						}

						rowCoords := make([]*models.Coordinates, 0)
						for _, coords := range coords.Coordinates[0] {
							newCoord := models.Coordinates{}

							for i, coord := range coords {
								if i == 0 {
									newCoord.Latitude = coord
								} else {
									newCoord.Longitude = coord
								}
							}
							rowCoords = append(rowCoords, &newCoord)
						}
						rowDB.Coordinates = rowCoords
						rowDB.ByteCoords = nil
					}
					dealerMapSlotsDB, err := api.db.GetDealerMapSlot(c, rowDB.RowID)
					if err != nil {
						glg.Error(err)
						return nil, err
					}

					for _, slotDB := range dealerMapSlotsDB {
						if slotDB.ByteCoords != nil {
							coords := models.Coords{}
							err = json.Unmarshal(slotDB.ByteCoords, &coords)
							if err != nil {
								glg.Error(err)
								return nil, err
							}

							zoneCoords := make([]*models.Coordinates, 0)
							for _, coords := range coords.Coordinates[0] {
								newCoord := models.Coordinates{}

								for i, coord := range coords {
									if i == 0 {
										newCoord.Latitude = coord
									} else {
										newCoord.Longitude = coord
									}
								}
								zoneCoords = append(zoneCoords, &newCoord)
							}
							slotDB.Coordinates = zoneCoords
							slotDB.ByteCoords = nil
						}
					}
					// As no further iteration is required on slots spo appending thw whole slice
					rowDB.Slots = append(rowDB.Slots, dealerMapSlotsDB...)

					zoneDB.Rows = append(zoneDB.Rows, rowDB)
				}
				parkingLotDB.Zones = append(parkingLotDB.Zones, zoneDB)
			}
			dealerMapDB.ParkingLots = append(dealerMapDB.ParkingLots, parkingLotDB)
		}
		mapsApi = append(mapsApi, dealerMapDB)
	}

	return mapsApi, nil
}
