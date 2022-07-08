package request

import (
	"bifrost/server/model/common/request"
	"bifrost/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
