package relative

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"tubitakPrototypeGo/database"
	"tubitakPrototypeGo/helpers"
)

func SetupPatientRelative(rg *gin.RouterGroup) {
	rg.POST("/sign_in", signPatient)
	rg.POST("/change_password", changePassword)
	rg.POST("/add_patient", addPatient)
	rg.GET("/patient_tracking_info/:patientId/:page", getPatientTrackingInfo)

}

func signPatient(c *gin.Context) {
	body := signRelative{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	if !helpers.EmailIsValid(body.Email) {
		helpers.MyAbort(c, "Check your email type!!!")
		return
	}
	var emailExist bool
	err = database.Db.QueryRow("select exists(select 1 from patient_relatives_table where email=$1)", body.Email).Scan(&emailExist)
	if err != nil {
		helpers.MyAbort(c, "Something went wrong check the server.")
		return
	}
	if !emailExist {
		helpers.MyAbort(c, "Girmis oldugunuz mail adresi gecerli degildir.")
		return
	}
	var token, password, patientTc string
	err = database.Db.QueryRow("select token,password,patient_tc from patient_relatives_table where email=$1 ", body.Email).Scan(&token, &password, &patientTc)
	if !helpers.Checkpassword(body.Password, password) {
		helpers.MyAbort(c, "Check Your password.")
		return
	}
	c.JSON(200, gin.H{
		"token":     token,
		"patientTc": patientTc,
	})

}

func changePassword(c *gin.Context) {
	body := changePasswordSt{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	token := c.GetHeader("token")
	var password string
	err = database.Db.QueryRow("select password from patient_relatives_table where token=$1", token).Scan(&password)
	if err != nil || !helpers.Checkpassword(body.OldPassword, password) {
		helpers.MyAbort(c, "Eski Parolanin dogru oldugundan emin olun. ")
		return
	} else {
		var checkChange bool
		newPassword, _ := helpers.Hashpassword(body.NewPassword)
		err = database.Db.QueryRow("SELECT  * from changepassword($1,$2,$3)", token, true, newPassword).Scan(&checkChange)
		if err != nil || !checkChange {
			helpers.MyAbort(c, "Eski Parolanin dogru oldugundan emin olun. ")
			return
		}
	}

	c.JSON(200, "Parola basariyla degistirilmistir.")

}

func addPatient(c *gin.Context) {
	body := addPatientSt{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Admin could not be found")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	token := c.GetHeader("token")
	var checkTokenExist bool
	err = database.Db.QueryRow("select exists(select 1 from patient_relatives_table where token=$1)", token).Scan(&checkTokenExist)
	if !checkTokenExist {
		helpers.MyAbort(c, "Birseyler hatali gitti lutfen yeniden baglanin!")
		return
	}

	_, err = database.Db.Query("insert into patient_table(patient_bd, patient_relative_name, patient_relative_phone_number, patient_relative_name2, patient_relative_phone_number2, patient_gender, patient_address, patient_tc, patient_name, patient_surname, patient_relative_surname, patient_relative_surname2) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)",
		body.PatientBd, body.PRName, body.PRNum, body.PRName2, body.PRNum2, body.PatientGender, body.PatientAddress, body.PatientTc, body.PatientName, body.PatientSurname, body.PRSurname, body.PRSurname2)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, "Patient Is Added")

}

const itemsPerPage = 5

func getPatientTrackingInfo(c *gin.Context) {
	token := c.GetHeader("token")
	var checkTokenExist bool
	err := database.Db.QueryRow("select exists(select 1 from patient_relatives_table where token=$1)", token).Scan(&checkTokenExist)
	if !checkTokenExist {
		helpers.MyAbort(c, "Birseyler hatali gitti lutfen yeniden baglanin!")
		return
	}
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
