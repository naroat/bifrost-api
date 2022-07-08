package initialize

import (
	_ "bifrost/server/source/example"
	_ "bifrost/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
