// voicerole is a plugin that assigns roles to users when they join voice channels
package voiceroles

//go:generate sqlboiler --no-hooks psql

import (
	"context"

	"github.com/RhykerWells/yagpdb/v2/bot"
	"github.com/RhykerWells/yagpdb/v2/common"
	"github.com/RhykerWells/yagpdb/v2/premium"
	"github.com/RhykerWells/yagpdb/v2/web"
)

type Plugin struct{}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "VoiceRoles",
		SysName:  "voiceroles",
		Category: common.PluginCategoryMisc,
	}
}

var logger = common.GetPluginLogger(&Plugin{})

var (
	_ common.Plugin      = (*Plugin)(nil)
	_ web.Plugin         = (*Plugin)(nil)
	_ bot.BotInitHandler = (*Plugin)(nil)
)

const (
	MaxVoiceRoles        = 1
	MaxVoiceRolesPremium = 10
)

func MaxConfigsForContext(ctx context.Context) int {
	if premium.ContextPremium(ctx) {
		return MaxVoiceRolesPremium
	}
	return MaxVoiceRoles
}

func RegisterPlugin() {
	p := &Plugin{}
	common.RegisterPlugin(p)

	common.InitSchemas("voiceroles", DBSchemas...)
}
