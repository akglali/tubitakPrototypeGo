package adminPanel

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"tubitakPrototypeGo/adminPanel/adminPanelDatabase"
	"tubitakPrototypeGo/database"
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
		if err := rows.Scan(&patient.PatientTc, &patient.PatientName, &patient.PatientSurname, &patient.PatientGender, &patient.PatientAddress, &patient.LastSeenTime); err != nil {
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

//get single patients all information

func getSinglePatientInfoRow(patientId string) allSinglePatientInfo {

	var patientInfo allSinglePatientInfo
	row := database.Db.QueryRow("select patient_tc,patient_name,patient_surname,patient_bd,patient_relative_name2,patient_relative_phone_number,patient_relative_name,patient_relative_phone_number2,patient_gender,patient_address from patient_table where patient_tc=$1", patientId)
	err := row.Scan(&patientInfo.PatientTc, &patientInfo.PatientName, &patientInfo.PatientSurname, &patientInfo.PatientBd, &patientInfo.PatientR1Name, &patientInfo.PatientR1Num, &patientInfo.PatientR2Name, &patientInfo.PatientR2Num, &patientInfo.PatientGender, &patientInfo.PatientAddress)
	if err != nil {
		return patientInfo
	}
	return patientInfo
}

const itemsPerPage = 3

func pagination(c *gin.Context, st interface{}) {
	var allRows []interface{}
	switch reflect.TypeOf(st).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(st)
		for i := 0; i < s.Len(); i++ {
			allRows = append(allRows, s.Index(i).Interface())

		}
	}
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Data could not reached")
		return
	}
	start := ((page) - 1) * itemsPerPage
	stop := start + itemsPerPage
	if start > len(allRows) {
		c.JSON(404, "No More Post")
		return
	}
	if stop > len(allRows) {
		stop = len(allRows)
	}

	if err != nil {
		helpers.MyAbort(c, "We can't get all posts!")
		return
	}
	c.JSON(200, allRows[start:stop])
}
