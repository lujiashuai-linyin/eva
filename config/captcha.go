package config

type Captcha struct {
	KeyLong     int    `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                // 验证码长度
	ImgWidth    int    `mapstructure:"img-width" json:"img-width" yaml:"img-width"`             // 验证码宽度
	ImgHeight   int    `mapstructure:"img-height" json:"img-height" yaml:"img-height"`          // 验证码高度
	GeeTestID   string `mapstructure:"gee-test-id" json:"gee-test-id" yaml:"gee-test-id"`       // 验证码高度
	GeeTestKey  string `mapstructure:"gee-test-key" json:"gee-test-key" yaml:"gee-test-key"`    // 验证码高度
	GeeTestType string `mapstructure:"gee-test-type" json:"gee-test-type" yaml:"gee-test-type"` // 验证码高度
}
