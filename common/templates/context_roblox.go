package templates

import (
	"github.com/RhykerWells/robloxgo"
)

func (c *Context) tmplGetRobloxUserByID(apiKey interface{}, target interface{}) (interface{}, error) {
	robloxClient, _ := robloxgo.Create(ToString(apiKey))
	user, err := robloxClient.GetUserByID(ToString(target))
	if err != nil {
		return user, nil
	}

	return user, err
}

func (c *Context) tmplGetRobloxUserByUsername(apiKey interface{}, target interface{}) (interface{}, error) {
	robloxClient, _ := robloxgo.Create(ToString(apiKey))
	user, err := robloxClient.GetUserByUsername(ToString(target))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Context) tmplGetRobloxGroupByID(apiKey interface{}, target interface{}) (interface{}, error) {
	robloxClient, _ := robloxgo.Create(ToString(apiKey))
	group, err := robloxClient.GetGroupByID(ToString(target))
	if err != nil {
		return nil, err
	}

	return group, nil // Don't return err, a nil output should indicate unknown/invalid group
}

func (c *Context) tmplUpdateMemberRole(apiKey, group interface{}, target interface{}, role interface{}) (interface{}, error) {
	robloxClient, _ := robloxgo.Create(ToString(apiKey))
	robloxGroup, err := robloxClient.GetGroupByID(ToString(group))
	if err != nil {
		return nil, err
	}

	groupRole, _ := robloxGroup.UpdateUserRole(ToString(target), ToString(role))

	return groupRole, nil // Don't return err, a nil output should indicate unknown/invalid group
}