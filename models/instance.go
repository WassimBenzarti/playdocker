package models

type Instance struct {
	Hostname    string `json:"hostname"`
	Ip          string `json:"ip"`
	Name        string `json:"name"`
	ProxyHost   string `json:"proxy_host"`
	SessionId   string `json:"session_id"`
	SessionHost string `json:"session_host"`
}
