package api

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"capregsoft.com/carlocator/service/lib"
	"capregsoft.com/carlocator/service/models"
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func (api *CarLocatorApiImpl) AddMapAPI(c *gin.Context, maps []*models.Map) ([]*models.Map, error) {
	var mapMutex, parkingLotMutex, zoneMutex, rowMutex, slotMutex sync.Mutex
	mapsDB := make([]*models.Map, 0)
	for _, dMap := range maps {
		mapSem := semaphore.NewWeighted(lib.MapMaxWorkers)
		mapEg, mapCtx := errgroup.WithContext(context.Background())

		var mapErr error
		var mapDB *models.Map
		mapEg.Go(func() error {
			err := mapSem.Acquire(mapCtx, 1)
			if err != nil {
				return err
			}
			defer mapSem.Release(1)
			mapDB, mapErr = api.AddMapCoords(c, dMap)
			if mapErr != nil {
				return mapErr
			}
			return nil
		})
		if err := mapEg.Wait(); err != nil {
			return nil, err
		}

		mapDB.ParkingLots = make([]*models.ParkingLot, 0)
		for _, parkingLot := range dMap.ParkingLots {
			LotSem := semaphore.NewWeighted(lib.ParkingLotMaxWorkers)
			lotEg, lotCtx := errgroup.WithContext(context.Background())
			var parkingLotDB *models.ParkingLot
			var lotErr error
			lotEg.Go(func() error {
				err := LotSem.Acquire(lotCtx, 1)
				if err != nil {
					return err
				}
				defer LotSem.Release(1)
				parkingLotDB, lotErr = api.AddParkingLotCoords(c, parkingLot, *mapDB.MapID)
				if lotErr != nil {
					return lotErr
				}
				return nil
			})
			if err := lotEg.Wait(); err != nil {
				return nil, err
			}

			parkingLotDB.Zones = make([]*models.Zone, 0)
			for _, zone := range parkingLot.Zones {
				zoneSem := semaphore.NewWeighted(lib.ZoneMaxWorkers)
				zoneEg, zoneCtx := errgroup.WithContext(context.Background())
				var zoneDB *models.Zone
				var zoneErr error
				zoneEg.Go(func() error {
					err := zoneSem.Acquire(zoneCtx, 1)
					if err != nil {
						return err
					}
					defer zoneSem.Release(1)
					zoneDB, zoneErr = api.AddZoneCoords(c, zone, *parkingLotDB.ParkingLotID)
					if zoneErr != nil {
						return zoneErr
					}
					return nil
				})
				if err := zoneEg.Wait(); err != nil {
					return nil, err
				}

				zoneDB.Rows = make([]*models.Row, 0)
				for _, row := range zone.Rows {
					rowSem := semaphore.NewWeighted(lib.RowMaxWorkers)
					rowEg, rowCtx := errgroup.WithContext(context.Background())
					var rowDB *models.Row
					var rowErr error
					rowEg.Go(func() error {
						err := rowSem.Acquire(rowCtx, 1)
						if err != nil {
							return err
						}
						defer rowSem.Release(1)
						rowDB, rowErr = api.AddRowCoords(c, row, *zoneDB.ZoneID)
						if rowErr != nil {
							return rowErr
						}
						return nil
					})
					if err := rowEg.Wait(); err != nil {
						return nil, err
					}

					rowDB.Slots = make([]*models.Slot, 0)
					for _, slot := range row.Slots {
						slotSem := semaphore.NewWeighted(lib.SlotMaxWorker)
						slotEg, slotCtx := errgroup.WithContext(context.Background())
						var slotDB *models.Slot
						var slotErr error
						slotEg.Go(func() error {
							err := slotSem.Acquire(slotCtx, 1)
							if err != nil {
								return err
							}
							defer slotSem.Release(1)
							slotDB, slotErr = api.AddSlotCoords(c, slot, *rowDB.RowID)
							if slotErr != nil {
								return slotErr
							}
							return nil
						})
						if err := slotEg.Wait(); err != nil {
							return nil, err
						}
						slotMutex.Lock()
						rowDB.Slots = append(rowDB.Slots, slotDB)
						slotMutex.Unlock()
					}
					rowMutex.Lock()
					zoneDB.Rows = append(zoneDB.Rows, rowDB)
					rowMutex.Unlock()
				}
				zoneMutex.Lock()
				parkingLotDB.Zones = append(parkingLotDB.Zones, zoneDB)
				zoneMutex.Unlock()
			}
			parkingLotMutex.Lock()
			mapDB.ParkingLots = append(mapDB.ParkingLots, parkingLotDB)
			parkingLotMutex.Unlock()
		}
		mapMutex.Lock()
		mapsDB = append(mapsDB, mapDB)
		mapMutex.Unlock()
	}
	return mapsDB, nil
}

func (api *CarLocatorApiImpl) AddMapCoords(c *gin.Context, dMap *models.Map) (*models.Map, error) {
	dMap.DealerID = new(string)
	dMap.MapID = new(string)

	*dMap.DealerID = c.MustGet("user-id").(string)
	*dMap.MapID = CreateUUID()

	if dMap.Coordinates != nil {
		dMap.Coords = new(string)
		mapCoords := make([]string, 0)
		for _, coords := range dMap.Coordinates {
			makeCoords := fmt.Sprintf("%v %v", *coords.Latitude, *coords.Longitude)
			mapCoords = append(mapCoords, makeCoords)
		}
		coords := strings.Join(mapCoords, ", ")
		*dMap.Coords = fmt.Sprintf("POLYGON ((%v))", coords)
	}

	mapDB, err := api.db.AddNewMapDB(c, dMap)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return mapDB, nil
}

func (api *CarLocatorApiImpl) AddParkingLotCoords(c *gin.Context, parkingLot *models.ParkingLot, mapID string) (*models.ParkingLot, error) {
	parkingLot.ParkingLotID = new(string)
	*parkingLot.ParkingLotID = CreateUUID()
	parkingLot.MapID = new(string)
	*parkingLot.MapID = mapID

	if parkingLot.Coordinates != nil {
		parkingLot.Coords = new(string)
		parkingCoords := make([]string, 0)
		for _, coords := range parkingLot.Coordinates {
			makeCoords := fmt.Sprintf("%v %v", *coords.Latitude, *coords.Longitude)
			parkingCoords = append(parkingCoords, makeCoords)
		}
		coords := strings.Join(parkingCoords, ", ")
		*parkingLot.Coords = fmt.Sprintf("POLYGON ((%v))", coords)
	}

	parkingLotDB, err := api.db.AddNewMapParkingSlotDB(c, parkingLot)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return parkingLotDB, nil
}

func (api *CarLocatorApiImpl) AddZoneCoords(c *gin.Context, zone *models.Zone, parkingLotID string) (*models.Zone, error) {
	zone.ZoneID = new(string)
	*zone.ZoneID = CreateUUID()
	zone.ParkingLotID = new(string)
	*zone.ParkingLotID = parkingLotID

	if zone.Coordinates != nil {
		zone.Coords = new(string)
		zoneCoords := make([]string, 0)
		for _, coords := range zone.Coordinates {
			makeCoords := fmt.Sprintf("%v %v", *coords.Latitude, *coords.Longitude)
			zoneCoords = append(zoneCoords, makeCoords)
		}
		coords := strings.Join(zoneCoords, ", ")
		*zone.Coords = fmt.Sprintf("POLYGON ((%v))", coords)
	}

	zoneDB, err := api.db.AddNewMapZoneDB(c, zone)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return zoneDB, nil
}

func (api *CarLocatorApiImpl) AddRowCoords(c *gin.Context, row *models.Row, zoneID string) (*models.Row, error) {
	row.RowID = new(string)
	*row.RowID = CreateUUID()
	row.ZoneID = new(string)
	*row.ZoneID = zoneID

	if row.Coordinates != nil {
		row.Coords = new(string)
		rowCoords := make([]string, 0)
		for _, coords := range row.Coordinates {
			makeCoords := fmt.Sprintf("%v %v", *coords.Latitude, *coords.Longitude)
			rowCoords = append(rowCoords, makeCoords)
		}
		coords := strings.Join(rowCoords, ", ")
		*row.Coords = fmt.Sprintf("POLYGON ((%v))", coords)
	}

	rowDB, err := api.db.AddNewMapRowsDB(c, row)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return rowDB, nil
}

func (api *CarLocatorApiImpl) AddSlotCoords(c *gin.Context, slot *models.Slot, rowID string) (*models.Slot, error) {
	slot.SlotID = new(string)
	*slot.SlotID = CreateUUID()
	slot.RowID = new(string)
	*slot.RowID = rowID

	if slot.Coordinates != nil {
		slot.Coords = new(string)
		slotCoords := make([]string, 0)
		for _, coords := range slot.Coordinates {
			makeCoords := fmt.Sprintf("%v %v", *coords.Latitude, *coords.Longitude)
			slotCoords = append(slotCoords, makeCoords)
		}
		coords := strings.Join(slotCoords, ", ")
		*slot.Coords = fmt.Sprintf("POLYGON ((%v))", coords)
	}
	slotDB, err := api.db.AddNewMapSlotDB(c, slot)
	if err != nil {
		glg.Error(err)
		return nil, err
	}
	return slotDB, nil
}
