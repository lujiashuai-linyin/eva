package config

type TencentSMS struct {
	Template Template `json:"template"`
}

type Template struct {
	Login    string `json:"login"`
	Register string `json:"register"`
	RePwd    string `json:"re_pwd"`
}
