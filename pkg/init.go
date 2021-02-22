package pkg

import (
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/config"
)

func init() {
	config.Init()
	model.Init()
}
