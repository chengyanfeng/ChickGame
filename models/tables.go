package models

import "time"

type User struct {
	Id          int       `json:"size:32;column:id;auto_increment"`
	OpenId      string    `json:"size:512;column:openid"`
	NickName    string    `json:"size:64;column:nickname"`
	Sex         string    `json:"size:32;column:sex"`
	Province    string    `json:"size:32;column:province"`
	City        string    `json:"size:32;column:city"`
	Country     string    `json:"size:32;column:country"`
	HeadimgUrl  string    `json:"size:512;column:headimgurl"`
	Privilege   string    `json:"size:32;column:privilege"`
	PhoneNumber int       `json:"size:15;column:phoneNumber"`
	Time        time.Time `json:"column:time"`
}



