package timezonecompanion

//go:generate sqlboiler --no-hooks psql

import (
	"github.com/RhykerWells/yagpdb/v2/common"
	"github.com/RhykerWells/yagpdb/v2/lib/when"
	"github.com/RhykerWells/yagpdb/v2/lib/when/rules"
	"github.com/RhykerWells/yagpdb/v2/timezonecompanion/trules"
	"github.com/RhykerWells/yagpdb/v2/common/templates"
)

type Plugin struct {
	DateParser *when.Parser
}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "TimezoneCompanion",
		SysName:  "timezonecompanion",
		Category: common.PluginCategoryMisc,
	}
}

var logger = common.GetPluginLogger(&Plugin{})

func RegisterPlugin() {

	w := when.New(&rules.Options{
		Distance:     10,
		MatchByOrder: true})

	w.Add(trules.Hour(rules.Override), trules.HourMinute(rules.Override))

	common.InitSchemas("timezonecompanion", DBSchemas...)
	common.RegisterPlugin(&Plugin{
		DateParser: w,
	})
	// register user timezone lookup for templates package to avoid import cycles
	templates.UserTimezoneLookup = GetUserTimezone
}
