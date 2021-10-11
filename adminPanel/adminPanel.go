package adminPanel

import (
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/database"
	"tubitakPrototypeGo/helpers"
)

func SetupAdminPanel(rg *gin.RouterGroup) {
	rg.GET("/all_beacons_admin", getAllBeaconInfo)
	rg.GET("/all_patients_admin", getAllPatientInfo)
	rg.POST("/add_admin", signup)
	rg.POST("/login_admin", login)
	rg.GET("/get_info/:patientId", getSinglePatientTrackingInfo)
}

func login(c *gin.Context) {
	body, err := loginStructFunc(c)
	var password string
	err = database.Db.QueryRow("select admin_password  from admin_table where admin_username=$1", body.Username).Scan(&password)
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
	err = database.Db.QueryRow("insert into admin_table(admin_username, admin_password) VALUES($1,$2) returning admin_username", body.Username, password).Scan(&username)
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
