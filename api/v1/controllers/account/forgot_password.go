package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	_ "github.com/lib/pq"
	"github.com/subosito/twilio"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type forgotpasswordController struct{}

var ForgotPassword forgotpasswordController

func (f *forgotpasswordController) SendPassword(rw http.ResponseWriter, req *http.Request) {

	var u models.Password
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	mobile_number := strconv.Itoa(u.MobileNumber)
	if err != nil {
		log.Fatal(err)
	}
	reg := []string{"+91", mobile_number}
	mobile_number = strings.Join(reg, "")
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 9)
	fmt.Println(mobile_number)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	fmt.Println("password", string(b))

	var (
		AccountSid = "AC996c79a79cbd18129c4cb47edd03870c"
		AuthToken  = "43fae8d7421e2b52556f4e69f8087cf4"
		From       = "+13052406907"
		To         = mobile_number
	)
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	user_exist, err := db.Query("SELECT mobile_number FROM users where mobile_number = $1", u.MobileNumber)
	if err != nil {
		log.Fatal(err)
	}
	flag := 1
	for user_exist.Next() {
		var mobile_number int
		err = user_exist.Scan(&mobile_number)
		if err != nil {
			log.Fatal(err)
		}

		if mobile_number == u.MobileNumber {
			flag = 0
		}
	}
	defer user_exist.Close()
	if flag == 0 {
		c := twilio.NewClient(AccountSid, AuthToken, nil)
		params := twilio.MessageParams{
			Body: "Forgot your password! Enter '" + string(b) + "' as your new password in sign up page",
		}
		s, resp, err := c.Messages.Send(From, To, params)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("From :", s)
			fmt.Println("Response :", resp)
		}

		key := []byte("traveling is fun")
		db_password := string(b)
		decrypt_password := controllers.Decrypt(key, db_password)
		update_pawssword, err := db.Query("UPDATE users set password = $1 where mobile_number = $2", decrypt_password, u.MobileNumber)
		if err != nil || update_pawssword == nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("User does not exist")
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   "User does not exist",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		fmt.Println("Mobile number already exist")
		flag = 0
		goto end
	}
end:
}
