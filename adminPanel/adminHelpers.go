package adminPanel

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/adminPanel/adminPanelDatabase"
	"tubitakPrototypeGo/helpers"
)

//adminHelpers helps us to simplify the code

// it is the struct that I use signup and login.
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

//Getting all patients information
func getAllPatientRows() ([]allPatientInfo, error) {

	rows, err := adminPanelDatabase.GetAllPatientDb()
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

//getting single patient tracking information to be able to see their path.
func getSinglePatientRows(patientId string) ([]singlePatientTrackingStruct, error) {

	rows, err := adminPanelDatabase.GetSinglePatientRowsDb(patientId)
	if err != nil {
		return nil, err
	}
	var allRows []singlePatientTrackingStruct
	for rows.Next() {
		var row singlePatientTrackingStruct
		if err := rows.Scan(&row.BeaconId, &row.BeaconLocation, &row.Distance, &row.SeenTime, &row.MapInfo); err != nil {
			return allRows, err
		}
		allRows = append(allRows, row)
	}
	return allRows, err
}

//getting single beacon tracking info. So admin can see all patients that are tracked by the beacon
func getSingleBeaconId(beaconId string) ([]singleBeaconTrackingStruct, error) {
	rows, err := adminPanelDatabase.GetSingleBeaconIdDb(beaconId)
	if err != nil {
		return nil, err
	}
	var allRows []singleBeaconTrackingStruct
	for rows.Next() {
		var row singleBeaconTrackingStruct
		if err := rows.Scan(&row.PatientTc, &row.SeenTime, &row.Distance, &row.Location, &row.MapInfo, &row.Minor, &row.Major); err != nil {
			return allRows, err
		}
		allRows = append(allRows, row)
	}
	return allRows, err

}

// getting all beacon list
func getAllBeaconRows() ([]allBeaconInfo, error) {
	rows, err := adminPanelDatabase.GetAllBeaconRowsDb()
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
