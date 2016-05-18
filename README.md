# Marga -- Way to solve all your problems.

*This application is for the use of the people of Mysore, A community Application that has been designed keeping in mind, to help solve day to day problems of a common man. We are using iOS for the Front-end and Golang for the Back-end.*

This README would normally document whatever steps are necessary to get the
application up and running.

# Golang installation

Refer the following video to install golang
https://www.youtube.com/watch?v=2PATwIfO5ag

Modify the path as per your defined name.
```
export GOROOT=/Users/{account_name}/Documents/WorkSpace/GoLang/go

export GOPATH=/Users/{account_name}/Documents/WorkSpace/GoLang/gopath

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
## Install Godep

Command godep helps build packages reproducibly by fixing their dependencies.

This tool assumes you are working in a standard Go workspace, as described in http://golang.org/doc/code.html.

To install godep run the following command
```
go get github.com/tools/godep

godep save
```
---
## Schema definitions

### create users table
```
CREATE TABLE users (
id SERIAL,
name varchar(100),
username varchar(100),
email varchar(320),
mobile_number varchar(100),
latitude float,
longitude float,
password varchar(100),
password_confirmation varchar(100),
city varchar(100) DEFAULT ‘MYSURU’,
device_token varchar(320),
created_at timestamptz,
my_points int,
type varchar(100),
PRIMARY KEY(id),
UNIQUE (username, mobile_number));
```

### create devices table
```
CREATE TABLE devices (
id int,
devise_token varchar(320),
user_id int,
CONSTRAINT devices_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
PRIMARY KEY(devise_token));
```

### create sessions table
```

CREATE TABLE sessions (
id int,
user_id int,
CONSTRAINT sessions_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
devise_token varchar(320),
CONSTRAINT sessions_devices_key FOREIGN KEY(devise_token)
REFERENCES devices(devise_token));
```

### create issues table
```

CREATE TABLE issues (
id SERIAL,
name varchar(320),
type varchar(100),
description varchar(2083),
latitude float,
longitude float,
image varchar(2083),
status boolean,
address varchar(2083),
created_at timestamptz,
user_id int,
CONSTRAINT issues_users_key FOREIGN KEY (user_id)
REFERENCES users(id),
PRIMARY KEY(id));
```

### create comments table
```

CREATE TABLE comments (
id SERIAL,
description varchar(2083),
user_id int,
CONSTRAINT comments_users_key FOREIGN KEY(user_id)
REFERENCES users(id),
issue_id int,
CONSTRAINT comments_issues_key FOREIGN KEY(issue_id)
REFERENCES issues(id),
PRIMARY KEY(id));
```

### create notifications table
```

CREATE TABLE notifications (
id SERIAL,
message varchar(2083),
sender_id int,
CONSTRAINT notifications_users_sender_key FOREIGN KEY(sender_id)
REFERENCES users(id),
reciever_id int,
CONSTRAINT notifications_users_reciever_key FOREIGN KEY(reciever_id)
REFERENCES users(id));
```


---
# Database instructions

To edit the users table schema you have to drop the tables in the following order

DROP TABLE devises;

DROP TABLE sessions;

DROP TABLE notifications;

DROP TABLE issues;

DROP TABLE comments;

DROP TABLE users;

---
# Inputs data and fields required for running API


### Create User

URL - http://localhost:3000/sign_up

Method POST

Data has to be sent in raw format
```
{"name":"steve","username":"jobs","email":"steve@example.com","password":"password","password_confirmation":"password","city":"mysore","mobile_number":"123456789","latitude":12345,"longitude":12345,"type":"user","devise_token":"039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"}

Response
```
{
	"Success": "true",
	"Message": "User created Successfully!",
	"User": {
		"Id": 1,
		"Name": "steve",
		"Username": "jobs",
		"Email": "steve@example.com",
		"Mobile_number": "123456789",
		"Latitude": 12345,
		"Longitude": 12345,
		"Password": "password",
		"Password_confirmation": "password",
		"City": "mysore",
		"Devise_token": "039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2",
		"Type": "user"
	},
	"Session": {
		"SessionId": 1,
		"DeviseToken": "039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"
	}
}

The user is logged in as he signs up into rVidi. A session is created as soon as he signs up.
```

### Sign In

URL - http://localhost:3000/log_in

Method POST

Data has to be sent in raw format
```
{"password":"password","mobile_number":"123456789","devise_token":"039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"}

Response
```
{
	"Success": "true",
	"Message": "Logged in Successfully",
	"User": {
		"Id": 1,
		"Name": "steve",
		"Username": "jobs",
		"Email": "steve@example.com",
		"Mobile_number": "123456789",
		"Latitude": 12345,
		"Longitude": 12345,
		"Password": "",
		"Password_confirmation": "",
		"City": "mysore",
		"Devise_token": "039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2",
		"Type": "user"
	},
	"Session": {
		"SessionId": 1,
		"DeviseToken": "039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"
	}
}
```

### Sign Out

URL - http://localhost:3000/log_out/039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2

Method GET

Response
```
{
	"Success": "true",
	"Message": "Logged out Successfully",
	"User": {
		"Id": 0,
		"Name": "",
		"Username": "",
		"Email": "",
		"Mobile_number": "",
		"Latitude": 0,
		"Longitude": 0,
		"Password": "",
		"Password_confirmation": "",
		"City": "",
		"Devise_token": "039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2",
		"Type": ""
	}
}
```

### Create Issue

URL - http://localhost:3001/create_issue

Method POST

Data has to be sent in raw format
```
{"name":"hi", "type":"street light", "description":"no strret light", "latitude":12345, "longitude":34567, "image":"sajgdjsahdgjsahdg", "status":true, "address":"1st main", "user_id":1}
```

Response
```

{
	"Success": "true",
	"Message": "Issue created Successfully!",
	"Issue": {
		"Id": 1,
		"Name": "hi",
		"Type": "street light",
		"Description": "no strret light",
		"Latitude": 12345,
		"Longitude": 34567,
		"Image": "sajgdjsahdgjsahdg",
		"Status": true,
		"Address": "1st main",
		"User_id": 1,
		"Corporator_id": 0,
		"Created_at": "2016-04-14T12:01:34.983278213+05:30"
	}
}
```
There are 7 Issues available: 'sew', 'elec', 'dog', 'watelec', 'lights', 'pot', 'garbage'

### List Issues

URL - http://localhost:3001/issues

Method GET

Response
```

{
	"Success": "true",
	"No_Of_Issues": 1,
	"Issue_Details": [
		{
			"Issue_id": 1,
			"Name": "hi",
			"Type": "street light",
			"Description": "no strret light",
			"Latitude": 12345,
			"Longitude": 34567,
			"Image": "sajgdjsahdgjsahdg",
			"Status": true,
			"Address": "1st main",
			"User_id": 1
		}
	]
}
```

### Listing the issue of a particular user

URL - http://localhost:3001/user/1/issues/1

Method GET

Response
```

{
	"Success": "true",
	"No_Of_Issues": 1,
	"Issue_Details": [
		{
			"Issue_id": 1,
			"Name": "hi",
			"Type": "street light",
			"Description": "no strret light",
			"Latitude": 12345,
			"Longitude": 34567,
			"Image": "sajgdjsahdgjsahdg",
			"Status": true,
			"Address": "1st main",
			"User_id": 1
		}
	]
}
```

### MyIssues

URL - http://localhost:3009//user/1/issues

Method GET

Response
```

{
	"Success": "true",
	"No_Of_Issues": 1,
	"Issue_Details": [
		{
			"Issue_id": 1,
			"Name": "hi",
			"Type": "street light",
			"Description": "no strret light",
			"Latitude": 12345,
			"Longitude": 34567,
			"Image": "sajgdjsahdgjsahdg",
			"Status": true,
			"Address": "1st main",
			"User_id": 1
		}
	]
}
```

### To create comments for an issue

URL -http://localhost:3000/create_comment

Method POST

Data has to be sent in raw format

```
{"User_id":1,"Issue_id":5,"Description":"I have the same issue"}
```

Response  

```
{
	"Success": "true",
	"Message": "Comment added successfully"
}
```

### To get comments for a particular issue

URL -http://localhost:3000//comment/issues/5

Method GET

Response  

```
{
	"Success": "true",
	"No_of_comments": 3,
	"Comment_details": [
		{
			"Comment_message": "Transformer issue near Vijaynagar 2nd stage.please take necessary action ASAP",
			"User_id": 10,
			"Name": "KISHAN"
		},
		{
			"Comment_message": "Facing issue",
			"User_id": 17,
			"Name": "KISHAN"
		},
		{
			"Comment_message": "I have the same issue",
			"User_id": 1,
			"Name": "KISHAN"
		}
	]
}
```

### Forgot password (sends new password to the given mobile number)

URL - http://localhost:3000/forgot_password

Method POST

Data has to be sent in raw format. Mobile number is of type integer

```
{"MobileNumber": 9916854333}
```

### Reset password (Update your password)

URL - http://localhost:3000/reset_password

Method POST

Data has to be sent in raw format.

```
{"User_id": 1, "OldPassword": passowrd, "NewPassword": Qwinix123 }
```

### To segrigate the cluster of issues

URL -http://localhost:3000/cluster/issues

Method POST

Data has to be sent in raw format

```
{"Issues":[1,22,3,4,515]}
```

Response  

```
{
	"Success": "true",
	"No_Of_Issues": 4,
	"Issue_Details": [
		{
			"Issue_id": 1,
			"Name": "KISHAN",
			"Type": "sew",
			"Description": "HEBBAL",
			"Latitude": 12.35936039128561,
			"Longitude": 76.61824557006044,
			"Image": "1454503290820.25.png",
			"Status": true,
			"Address": "HEBBAL\n",
			"User_id": 1
		},
		{
			"Issue_id": 22,
			"Name": "APEKSHA",
			"Type": "lights",
			"Description": "STREET LIGHTS NOT WORKING",
			"Latitude": 12.35916855900871,
			"Longitude": 76.61823935799828,
			"Image": "1457501102185.19.png",
			"Status": true,
			"Address": "VIJAYNAGAR",
			"User_id": 18
		},
		{
			"Issue_id": 3,
			"Name": "SAJITH",
			"Type": "sew",
			"Description": "NVKVBL",
			"Latitude": 12.35915686951162,
			"Longitude": 76.61823252682942,
			"Image": "1454506498023.37.png",
			"Status": true,
			"Address": "VKBKB",
			"User_id": 2
		},
		{
			"Issue_id": 4,
			"Name": "SAJITH",
			"Type": "sew",
			"Description": "VKBKB",
			"Latitude": 12.3591568321169,
			"Longitude": 76.61823267589534,
			"Image": "1454506677063.38.png",
			"Status": true,
			"Address": "VKBKB",
			"User_id": 2
		}
	]
}
```


README.md
