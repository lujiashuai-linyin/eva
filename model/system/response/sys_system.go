package response

import "eva/biz/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
