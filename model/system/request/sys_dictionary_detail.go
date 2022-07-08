package request

import (
	"bifrost/server/model/common/request"
	"bifrost/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
