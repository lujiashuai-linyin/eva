package gt3

const VERSION string = "3.0.0"

const (
	ClientTypeWeb     string = "web"     // web（pc浏览器）
	ClientTypeH5      string = "h5"      // h5（手机浏览器，包括webview）
	ClientTypeNative  string = "native"  // native（原生app）
	ClientTypeUnknown string = "unknown" // unknown（未知）
)

const (
	FN_CHALLENGE string = "geetest_challenge"
	FN_VALIDATE  string = "geetest_validate"
	FN_SECCODE   string = "geetest_seccode"
)

const (
	API_URL          string = "http://api.geetest.com"
	REGISTER_HANDLER string = "/register.php"
	VALIDATE_HANDLER string = "/validate.php"
	JSON_FORMAT      bool   = false
)

const (
	// GeetestChallenge 极验二次验证表单传参字段 chllenge
	GeetestChallenge string = "geetest_challenge"
	// GeetestValidate 极验二次验证表单传参字段 validate
	GeetestValidate string = "geetest_validate"
	// GeetestSeccode 极验二次验证表单传参字段 seccode
	GeetestSeccode string = "geetest_seccode"
	// GT_STATUS_SESSION_KEY 极验验证API服务状态Session Key
	GT_STATUS_SESSION_KEY string = "gt_server_status"
)
