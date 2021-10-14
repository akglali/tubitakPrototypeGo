package adminPanelDatabase

import (
	"database/sql"
	"tubitakPrototypeGo/database"
)

//adminDb is used for admin database calls.

// LoginDb loginDb
func LoginDb(username string, password *string) error {
	err := database.Db.QueryRow("select admin_password  from admin_table where admin_username=$1", username).Scan(&*password)
	return err
}
func SignUpDb(username, password string, savedUsername *string) error {
	err := database.Db.QueryRow("insert into admin_table(admin_username, admin_password) VALUES($1,$2) returning admin_username", username, password).Scan(&*savedUsername)
	return err

}

//Patient Calls

func GetAllPatientDb() (*sql.Rows, error) {
	rows, err := database.Db.Query("select patient_tc,patient_name,patient_bd,patient_relative_name,patient_relative_phone_number,patient_relative_name2,patient_relative_phone_number2,patient_gender,patient_address from patient_table")
	return rows, err
}

func GetSinglePatientRowsDb(patientId string) (*sql.Rows, error) {
	rows, err := database.Db.Query("select bdt.device_id,location,distance,seen_time,bdt.google_map_link from patient_tracker_info_table  left join beacon_devices_table as bdt on patient_tracker_info_table.beacon_id = bdt.device_id left join patient_table as pt on pt.patient_id=patient_tracker_info_table.patient_id where patient_tc=$1  order by seen_time desc", patientId)
	return rows, err
}

//Devices Calls

func GetAllBeaconRowsDb() (*sql.Rows, error) {
	rows, err := database.Db.Query("select * from beacon_devices_table")
	return rows, err
}

func GetSingleBeaconIdDb(beaconId string) (*sql.Rows, error) {
	rows, err := database.Db.Query("select pt.patient_tc,seen_time,distance,bdt.location,bdt.google_map_link,bdt.minor,bdt.major from patient_tracker_info_table left join patient_table pt on patient_tracker_info_table.patient_id = pt.patient_id left join beacon_devices_table as bdt  on  bdt.device_id=patient_tracker_info_table.beacon_id where beacon_id=$1 order by seen_time", beaconId)
	return rows, err
}
