package adminPanel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"tubitakPrototypeGo/adminPanel/adminPanelDatabase"
	"tubitakPrototypeGo/helpers"
)

func SetupAdminPanel(rg *gin.RouterGroup) {
	rg.GET("/all_beacons_admin/:page", getAllBeaconInfo)
	rg.GET("/all_patients_admin/:page", getAllPatientInfo)
	rg.POST("/add_admin", signup)
	rg.POST("/login_admin", login)
	rg.GET("/get_info_patient/:patientId/:page", getSinglePatientTrackingInfo)
	rg.GET("/get_single_patient/:singlePatientId", getSinglePatientInfo)
	rg.GET("/get_info_beacon/:beaconId/:page", getSingleBeaconTrackingInfo)

}

const itemsPerPage = 10

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
		c.JSON(200, "Welcome admin "+body.Username)
		return
	} else {
		helpers.MyAbort(c, "Password or username is wrong")
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
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	allBeaconsInfoRows, err := getAllBeaconRows(offSet * itemsPerPage)
	if err != nil {
		fmt.Println(err)
		helpers.MyAbort(c, "Could not reach beacons info")
		return
	}
	c.JSON(200, allBeaconsInfoRows)
}

func getAllPatientInfo(c *gin.Context) {
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	allPatientsInfoRows, err := getAllPatientRows(offSet * itemsPerPage)
	if err != nil {
		helpers.MyAbort(c, "Could not reach patients info")
		return
	}
	c.JSON(200, allPatientsInfoRows)
}

func getSinglePatientTrackingInfo(c *gin.Context) {
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	patientId := c.Param("patientId")
	allPatientTrackInfo, err := getSinglePatientRows(patientId, offSet*itemsPerPage)
	if err != nil {
		helpers.MyAbort(c, "Something went wrong for "+patientId)
		return
	}
	c.JSON(200, allPatientTrackInfo)

}

func getSingleBeaconTrackingInfo(c *gin.Context) {
	offSet, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		helpers.MyAbort(c, "Page Number Format is wrong")
		return
	}
	beaconId := c.Param("beaconId")
	allBeaconTrackingInfo, err := getSingleBeaconId(beaconId, offSet*itemsPerPage)
	if err != nil {
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
