package login

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/database"
	"tubitakPrototypeGo/helpers"
)

func SetupLogin(rg *gin.RouterGroup) {
	rg.POST("/login", login)

}

func login(c *gin.Context) {
	body := loginStruct{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		fmt.Println(body.PatientTc)
		helpers.MyAbort(c, "Bad Input")
		return
	}
	var patientId, patientName string
	err = database.Db.QueryRow("select patient_id,patient_name,patient_tc from patient_table where patient_tc=$1", body.PatientTc).Scan(&patientId, &patientName)
	if err != nil {
		helpers.MyAbort(c, "There is no such a patient")
		return
	}
	c.JSON(200, gin.H{
		"patientId":   patientId,
		"patientName": patientName,
	})

}
