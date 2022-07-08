package request

import (
	"bifrost/server/model/common/request"
	"bifrost/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
