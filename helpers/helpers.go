package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
	"net/smtp"
	"time"
)

type errorStruct struct {
	Error string
}

func MyAbort(c *gin.Context, str string) {
	c.AbortWithStatusJSON(400, errorStruct{Error: str})
}

var (
	verifier = emailverifier.NewVerifier()
	otpChars = "1234567890"
)

func EmailIsValid(email string) bool {

	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return false
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return false
	}

	fmt.Println("email validation result", ret)

	return true

}

//6 digits code is generated!
func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

//Code will be sent to email.
func SendEmail(code, mail string) {
	//put ur e-mail address that you want to sent e-mail by.
	from := "emailverifiy8@gmail.com"
	//put your email' password!!!
	pass := "emailonayla"

	to := []string{
		mail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("To: " + mail + "\r\n" +
		"Subject: Verification Code\r\n" +
		"\r\n" +
		"Hello dear,\r\n" + "Your code is\n" +
		code)

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Is Successfully sent.")

}

// TokenGenerator  is generated here
func TokenGenerator() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	currentTime := time.Now().Format("Mon Jan _2 15:04:05 2006")
	b = append(b, []byte(currentTime)...)

	return hex.EncodeToString(b)
}
