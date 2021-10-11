package adminPanel

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/database"
	"tubitakPrototypeGo/helpers"
)

func loginStructFunc(c *gin.Context) (loginStruct, error) {
	body := loginStruct{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
		return body, err
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return body, err
	}
	return body, err
}

func getAllPatientRows() ([]allPatientInfo, error) {

	rows, err := database.Db.Query("select patient_tc,patient_name,patient_bd,patient_relative_name,patient_relative_phone_number,patient_relative_name2,patient_relative_phone_number2,patient_gender,patient_address from patient_table")
	if err != nil {
		return nil, err
	}
	var patientsInfo []allPatientInfo
	for rows.Next() {
		var patient allPatientInfo
		if err := rows.Scan(&patient.PatientTc, &patient.PatientName, &patient.PatientBd, &patient.PatientR1Name, &patient.PatientR1Num, &patient.PatientR2Name, &patient.PatientR2Num, &patient.PatientGender, &patient.PatientAddress); err != nil {
			return patientsInfo, err
		}
		patientsInfo = append(patientsInfo, patient)
	}
	return patientsInfo, err

}

func getSinglePatientRows(patientId string) ([]singlePatientTrackingStruct, error) {

	rows, err := database.Db.Query("select bdt.device_id,location,distance,seen_time from patient_tracker_info_table  left join beacon_devices_table as bdt on patient_tracker_info_table.beacon_id = bdt.device_id where patient_id=$1 order by seen_time desc", patientId)
	if err != nil {
		return nil, err
	}
	var allRows []singlePatientTrackingStruct
	for rows.Next() {
		var row singlePatientTrackingStruct
		if err := rows.Scan(&row.BeaconId, &row.BeaconLocation, &row.Distance, &row.SeenTime); err != nil {
			return allRows, err
		}
		allRows = append(allRows, row)
	}
	return allRows, err
}

func getAllBeaconRows() ([]allBeaconInfo, error) {
	rows, err := database.Db.Query("select * from beacon_devices_table")
	if err != nil {
		return nil, err
	}
	var beaconsInfo []allBeaconInfo

	for rows.Next() {
		var pst allBeaconInfo
		if err := rows.Scan(&pst.DeviceId, &pst.Location, &pst.Major, &pst.Minor, &pst.GoogleMapLink); err != nil {
			return beaconsInfo, err
		}
		beaconsInfo = append(beaconsInfo, pst)
	}
	return beaconsInfo, err
}
