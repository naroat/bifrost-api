package request

import (
	"bifrost/server/model/{{.Package}}"
	"bifrost/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
