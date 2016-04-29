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
		panic(err)
	}
	reg := []string{"+91", mobile_number}
	mobile_number = strings.Join(reg, "")
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 9)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	var (
		AccountSid = "AC996c79a79cbd18129c4cb47edd03870c"
		AuthToken  = "43fae8d7421e2b52556f4e69f8087cf4"
		From       = "+13052406907"
		To         = mobile_number
		)
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	user_exist, err := db.Query("SELECT mobile_number FROM users where mobile_number = $1", u.MobileNumber)
	if err != nil {
		panic(err)
	}
	flag := 1
	for user_exist.Next() {
		var mobile_number int
		err = user_exist.Scan(&mobile_number)
		if err != nil {
			panic(err)
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

			b, err := json.Marshal(models.LogErrorMessage{
				Success: "false",
				Error:   err,
				})
			if err != nil {
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			fmt.Println("Mobile number already exist")
			flag = 0
			goto end

		} else {
			fmt.Println("From :", s)
			fmt.Println("Response :", resp)
		}

		key := []byte("traveling is fun")
		db_password := []byte(string(b))
		encrypt_password := controllers.Encrypt(key, db_password)
		update_pawssword, err := db.Query("UPDATE users set password = $1 where mobile_number = $2", encrypt_password, u.MobileNumber)
		if err != nil || update_pawssword == nil {
			panic(err)
		}
		fmt.Println("Temporary Password sent successfully")
		b, err := json.Marshal(models.SuccessMessage{
			Success: "true",
			Message: "Temporary Password sent successfully",
			})
		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		flag = 0
		goto end
	} else {
		fmt.Println("User does not exist")
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   "User does not exist",
			})
		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		fmt.Println("Mobile number already exist")
		flag = 0
		goto end
	}
	end:
}

func (f *forgotpasswordController) ResetPassword(rw http.ResponseWriter, req *http.Request) {

	var u models.ResetPassword
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	get_password, err := db.Query("SELECT password FROM users where id = $1", u.User_id)
	if err != nil {
		panic(err)
	}
	for get_password.Next() {
		var old_password string
		err = get_password.Scan(&old_password)
		key := []byte("traveling is fun")
		db_password := old_password
		decrypt_password := controllers.Decrypt(key, db_password)
		if err != nil {
			panic(err)
		}

		if decrypt_password == u.OldPassword {
			update_pawssword, err := db.Query("UPDATE users set password = $1 where id = $2", u.NewPassword, u.User_id)
			if err != nil || update_pawssword == nil {
				panic(err)
			}
			fmt.Println("New password updated successfully")
			b, err := json.Marshal(models.SuccessMessage{
				Success: "true",
				Message: "New password updated successfully",
				})
			if err != nil {
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto reset_password_end
		} else {
			fmt.Println("Password reset failed")
			b, err := json.Marshal(models.ErrorMessage{
				Success: "false",
				Error:   "Password reset failed",
				})
			if err != nil {
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto reset_password_end
		}
	}
	defer get_password.Close()
	reset_password_end:
}
