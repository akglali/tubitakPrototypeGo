package relative

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tubitakPrototypeGo/database"
	"tubitakPrototypeGo/helpers"
)

func SetupPatientRelative(rg *gin.RouterGroup) {
	rg.POST("/sign_in", signPatient)

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
	var token string
	err = database.Db.QueryRow("select token from patient_relatives_table where email=$1 and password=$2", body.Email, body.Password).Scan(&token)
	if err != nil {
		helpers.MyAbort(c, "Check Your password.")
		return
	} else {
		c.JSON(200, token)
	}

}
