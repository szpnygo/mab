package ffapp

import (
	"github.com/szpnygo/mab/utils/configx"
)

type FFApp struct {
	Config *configx.FFAppConf
	token  string
}

func NewFFAppHelper(config *configx.FFAppConf) *FFApp {
	return &FFApp{
		Config: config,
	}
}
