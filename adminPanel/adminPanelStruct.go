package adminPanel

type allPatientInfo struct {
	PatientTc      string
	PatientName    string
	PatientBd      string
	PatientR1Name  string
	PatientR1Num   string
	PatientR2Name  string
	PatientR2Num   string
	PatientGender  string
	PatientAddress string
}

type loginStruct struct {
	Username string
	Password string
}

type allBeaconInfo struct {
	DeviceId      string
	Location      string
	Major         string
	Minor         string
	GoogleMapLink string
}

type singlePatientTrackingStruct struct {
	BeaconId       string
	BeaconLocation string
	Distance       string
	SeenTime       string
}
