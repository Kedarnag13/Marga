package controllers

import (
	"database/sql"
	"fmt"
	"github.com/Qwinix/rVidi-Go/api/v1/controllers"
	"github.com/anachronistic/apns"
	_ "github.com/lib/pq"
	"log"
)

func Send_notification(senderid int, recieverid int, message string) (string, string) {
	flag := 1
	var response1 string
	var response2 string

	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("senderid notifications:", senderid)
	fmt.Println("recieverid notifications:", recieverid)

	user_session_existance := controllers.Check_for_user_session(recieverid)
	if senderid == recieverid {
		response1, response2 = "Notification cannot be sent in loop!", ""
		goto end
	} else if user_session_existance == true {
		tokens, err := db.Query("SELECT devise_token FROM devices WHERE user_id=$1", recieverid)
		if err != nil {
			log.Fatal(err)
		}
		for tokens.Next() {
			var devise_token string
			err := tokens.Scan(&devise_token)
			fmt.Println("devise_token in notifications:", devise_token)
			if err != nil {
				log.Fatal(err)
			}
			payload := apns.NewPayload()
			payload.Alert = message
			payload.Sound = "bingbong.aiff"
			pn := apns.NewPushNotification()
			pn.DeviceToken = devise_token
			pn.AddPayload(payload)

			pn.Set("Sender_id", senderid)
			pn.Set("Reciever_id", recieverid)

			client := apns.NewClient("gateway.push.apple.com:2195", "actualcrtAPNS.pem", "actualkeyAPNS.pem")
			resp := client.Send(pn)

			alert, _ := pn.PayloadString()
			if alert == "" {
				log.Fatal(err)
			}

			flag = 0
			resp1 := resp.Success
			if resp1 == true {
				response1, response2 = "true", devise_token
			} else {
				response1, response2 = "false", devise_token
			}
		}
		defer tokens.Close()
	}
	if user_session_existance == false || flag == 1 {
		response1, response2 = "Reciever does not have session", ""
	}
end:
	db.Close()
	fmt.Println("response1 notifications:", response1)
	fmt.Println("response2 notifications:", response2)
	return response1, response2
}
