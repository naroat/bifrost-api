package global

{{- if .HasGlobal }}

import "bifrost/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}