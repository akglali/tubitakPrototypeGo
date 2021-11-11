package relativeDatabase

import (
	"database/sql"
	"tubitakPrototypeGo/database"
)

func SinglePatient(patientId string, offSet int) (*sql.Rows, error) {
	rows, err := database.Db.Query("select location,distance,seen_time,bdt.google_map_link from patient_tracker_info_table  left join beacon_devices_table as bdt on patient_tracker_info_table.beacon_id = bdt.device_id left join patient_table as pt on pt.patient_id=patient_tracker_info_table.patient_id where patient_tc=$1  order by seen_time desc LIMIT 5 OFFSET $2", patientId, offSet)
	return rows, err
}
