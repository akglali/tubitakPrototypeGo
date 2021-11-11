package relative

import (
	"tubitakPrototypeGo/relative/relativeDatabase"
)

func getSinglePatientRows(patientId string, offSet int) ([]singlePatient, error) {
	rows, err := relativeDatabase.SinglePatient(patientId, offSet)
	if err != nil {
		return nil, err
	}
	var allRows []singlePatient
	for rows.Next() {
		var row singlePatient
		if err := rows.Scan(&row.Location, &row.Distance, &row.SeenTime, &row.GoogleMapLink); err != nil {
			return allRows, err
		}
		allRows = append(allRows, row)
	}
	return allRows, err

}
