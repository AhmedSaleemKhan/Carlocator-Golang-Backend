package db

import "github.com/gin-gonic/gin"

func (db *CarLocatorDBImpl) DeleteKeyDB(c *gin.Context, keyID string) error {
	tx := db.conn.MustBegin()
	_, err := tx.Exec(`DELETE FROM findr.veh_key WHERE key_id=$1 ;`, keyID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db *CarLocatorDBImpl) DeleteVehicleDB(c *gin.Context, vehicleId string) error {
	tx := db.conn.MustBegin()
	_, err := tx.Exec(`DELETE FROM findr.vehicle WHERE veh_id=$1`, vehicleId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
