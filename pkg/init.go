package pkg

import (
	"github.com/mhchlib/mconfig-admin/pkg/config"
	"github.com/mhchlib/mconfig-admin/pkg/model"
)

func init() {
	config.Init()
	model.Init()
}
