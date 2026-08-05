package messagecreator

import (
	"github.com/RhykerWells/yagpdb/v2/common"
)

type Plugin struct{}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "Message Creator",
		SysName:  "messagecreator",
		Category: common.PluginCategoryMisc,
	}
}

func RegisterPlugin() {
	common.RegisterPlugin(&Plugin{})
}
