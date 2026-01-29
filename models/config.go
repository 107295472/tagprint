package models

type ServerConfig struct {
	Host                              *string `json:"Host,omitempty"`
	User                              *string `json:"User,omitempty"`
	Pass                              *string `json:"Pass,omitempty"`
	Database                          *string `json:"Database,omitempty"`
	RedisAddr                         *string `json:"RedisAddr,omitempty"`
	RedisPassword                     *string `json:"RedisPassword,omitempty"`
	RedisDB                           *int64  `json:"RedisDB,omitempty"`
	LnhxURL                           *string `json:"LnhxUrl,omitempty"`
	Username                          *string `json:"Username,omitempty"`
	Password                          *string `json:"Password,omitempty"`
	Yzm                               *string `json:"Yzm,omitempty"`
	ManageHxSetTokenPosturl           *string `json:"manage_hx_set_token_posturl,omitempty"`
	ManageReportScreenStatePosturl    *string `json:"manage_report_screen_state_posturl,omitempty"`
	ManageLoginPosturl                *string `json:"manage_login_posturl,omitempty"`
	ManageSetReportScreenStatePosturl *string `json:"manage_set_report_screen_state_posturl,omitempty"`
	ManageLoginName                   *string `json:"manage_login_name,omitempty"`
	ManageLoginPass                   *string `json:"manage_login_pass,omitempty"`
	ManageURL                         *string `json:"manage_url,omitempty"`
	Interval                          *int64  `json:"interval,omitempty"`
	HttpPort                          *int64  `json:"HttpPort,omitempty"`
}
type GlobalApp struct {
	ManageToken string
	AlctToken   string
}
