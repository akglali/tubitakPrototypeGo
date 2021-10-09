package patientTracker

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"time"
	"tubitakPrototypeGo/database"
	"tubitakPrototypeGo/helpers"
)

func PatientTrackerSetup(rg *gin.RouterGroup) {
	rg.POST("/send_location", sendLocation)
	rg.GET("/get_all_beacon_info", getAllBeacon)
}

func sendLocation(c *gin.Context) {
	body := locationInfo{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	currentTime := time.Now().Format("Mon Jan _2 15:04:05 2006")
	distance, _ := decimal.NewFromString(body.Distance)

	_, err = database.Db.Query("insert into patient_tracker_info_table( patient_id, beacon_id, seen_time, distance) values ($1,$2,$3,$4)", body.PatientId, body.BeaconId, currentTime, distance)
	c.JSON(200, "New location Is added")

}

func getAllBeacon(c *gin.Context) {
	allBeaconsInfoRows, _ := getAllBeaconRows()
	c.JSON(200, allBeaconsInfoRows)

}

func getAllBeaconRows() ([]allBeaconsInfo, error) {
	rows, err := database.Db.Query("select device_id,location from beacon_devices_table")
	if err != nil {
		return nil, err
	}

	var beaconsInfo []allBeaconsInfo

	for rows.Next() {
		var pst allBeaconsInfo
		if err := rows.Scan(&pst.BeaconId, &pst.LocationOfBeacon); err != nil {
			return beaconsInfo, err
		}
		beaconsInfo = append(beaconsInfo, pst)
	}

	return beaconsInfo, err
}
