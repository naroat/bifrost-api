package response

import "bifrost/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
