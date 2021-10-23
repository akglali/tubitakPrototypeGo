package adminPanel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/adminPanel/adminPanelDatabase"
	"tubitakPrototypeGo/helpers"
)

func SetupAdminPanel(rg *gin.RouterGroup) {
	rg.GET("/all_beacons_admin", getAllBeaconInfo)
	rg.GET("/all_patients_admin", getAllPatientInfo)
	rg.POST("/add_admin", signup)
	rg.POST("/login_admin", login)
	rg.GET("/get_info_patient/:patientId", getSinglePatientTrackingInfo)
	rg.GET("/get_single_patient/:singlePatientId", getSinglePatientInfo)
	rg.GET("/get_info_beacon/:beaconId", getSingleBeaconTrackingInfo)

}

func login(c *gin.Context) {
	body, err := loginStructFunc(c)
	var password string
	err = adminPanelDatabase.LoginDb(body.Username, &password)
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	passwordTrue := CheckPassword(body.Password, password)

	if passwordTrue {
		c.JSON(200, "Hos geldin admin "+body.Username)
		return
	} else {
		helpers.MyAbort(c, "Password and username is wrong")
		return
	}
}

func signup(c *gin.Context) {
	body, err := loginStructFunc(c)
	password, _ := HashPassword(body.Password)
	var username string
	err = adminPanelDatabase.SignUpDb(body.Username, password, &username)
	if err != nil {
		helpers.MyAbort(c, "Admin Is already exist")
		return
	}
	c.JSON(200, "Admin "+username+" is added ")

}

func getAllBeaconInfo(c *gin.Context) {
	allBeaconsInfoRows, err := getAllBeaconRows()
	if err != nil {
		helpers.MyAbort(c, "Could not reach beacons info")
	}
	c.JSON(200, allBeaconsInfoRows)
}

func getAllPatientInfo(c *gin.Context) {
	allPatientsInfoRows, err := getAllPatientRows()
	if err != nil {
		helpers.MyAbort(c, "Could not reach patients info")
	}
	c.JSON(200, allPatientsInfoRows)
}

func getSinglePatientTrackingInfo(c *gin.Context) {
	patientId := c.Param("patientId")
	allPatientTrackInfo, err := getSinglePatientRows(patientId)
	if err != nil {
		helpers.MyAbort(c, "Something went wrong for "+patientId)
		return
	}
	c.JSON(200, allPatientTrackInfo)

}

func getSingleBeaconTrackingInfo(c *gin.Context) {
	beaconId := c.Param("beaconId")
	allBeaconTrackingInfo, err := getSingleBeaconId(beaconId)
	if err != nil {
		fmt.Println(err.Error())
		helpers.MyAbort(c, "Something went wrong for "+beaconId)
		return
	}
	c.JSON(200, allBeaconTrackingInfo)
}

func getSinglePatientInfo(c *gin.Context) {
	patientId := c.Param("singlePatientId")
	row := getSinglePatientInfoRow(patientId)
	c.JSON(200, row)

}
