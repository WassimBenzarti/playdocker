package models

type Session struct {
	Host         string              `json:"host"`
	Id           string              `json:"id"`
	PwdIPAddress string              `json:"pwd_ip_address"`
	Ready        bool                `json:"ready"`
	UserId       string              `json:"user_id"`
	Instances    map[string]Instance `json:"instances"`
}
