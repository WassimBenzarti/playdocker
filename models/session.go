package models

type Session struct {
	// Not used
	Host         string `json:"host"`
	Id           string `json:"id"`
	PwdIPAddress string `json:"pwd_ip_address"`
	Ready        bool   `json:"ready"`
	UserId       string `json:"user_id"`
}
