package zabbix

// For `RoleObject` field: `Type`
const (
	RoleTypeUser       = 1
	RoleTypeAdmin      = 2
	RoleTypeSuperAdmin = 3
)

// RoleObject
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/role/object
type RoleObject struct {
	RoleID int    `json:"roleid,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   int    `json:"type,omitempty"` // has defined consts, see above

	// Can't be used with role.create
	ReadOnly int `json:"readonly,omitempty"`

	// for role creation
	Rules RoleRuleObject `json:"rules,omitempty"`
}

// For `RoleRuleObject` field: `UIsDefaultAccess`
const (
	UIsAccessTypeDeny  = 0
	UIsAccessTypeAllow = 1
)

// For `RoleRuleObject` field: `ModulesDefaultAccess`
const (
	ModulesAccessTypeDeny  = 0
	ModulesAccessTypeAllow = 1
)

// For `RoleRuleObject` field: `RoleActionsDefaultAccess`
const (
	RoleActionsAccessTypeDeny  = 0
	RoleActionsAccessTypeAllow = 1
)

// For `RoleRuleObject` field: `APIAccess`
const (
	APIAccessTypeDeny  = 0
	APIAccessTypeAllow = 1
)

// For `RoleRuleObject` field: `APIMode`
const (
	APIModeDenyList   = 0
	APIModeAccessList = 1
)

// For `ModuleObject` field: `Status`
const (
	ModuleStatusEnabled  = 0
	ModuleStatusDisabled = 1
)

// For `RoleActionObject` field: `Status`
const (
	RoleActionStatusDisabled = 0
	RoleActionStatusEnabled  = 1
)

// RoleRuleObject
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/role/object#role_rules
type RoleRuleObject struct {
	UIs                      []UIElement        `json:"ui"`
	UIsDefaultAccess         int                `json:"ui.default_access"` // has defined consts, see above
	Modules                  []ModuleObject     `json:"modules"`
	ModulesDefaultAccess     int                `json:"modules.default_access"` // has defined consts, see above
	APIAccess                int                `json:"api.access"`             // has defined consts, see above
	APIMode                  int                `json:"api.mode"`               // has defined consts, see above
	APIMethods               []string           `json:"api,omitempty"`
	RoleActions              []RoleActionObject `json:"actions"`
	RoleActionsDefaultAccess int                `json:"actions.default_access"` // has defined consts, see above
}

// UIElement
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/role/object#ui_element
type UIElement struct {
	Name   string `json:"name,omitempty"`
	Status int    `json:"status,omitempty"`
}

// ModuleObject
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/role/object#module
type ModuleObject struct {
	ModuleID int `json:"moduleid,omitempty"`
	Status   int `json:"status,omitempty"`
}

// RoleActionObject
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/role/object#action
type RoleActionObject struct {
	Name   string `json:"name"`
	Status int    `json:"status,omitempty"`
}

// RoleGetParams
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/role/get#parameters
type RoleGetParams struct {
	GetParameters

	RoleIDs []int `json:"roleids,omitempty"`

	SelectRules SelectQuery `json:"selectRules,omitempty"`
	SelectUsers SelectQuery `json:"selectUsers,omitempty"`
	SortField   []string    `json:"sortfield,omitempty"` // roleid and/or name
}

// Structure to store creation result
type roleCreateResult struct {
	RoleIDs []int `json:"roleids"`
}

// Structure to store deletion result
type roleDeleteResult struct {
	RoleIDs []int `json:"roleids"`
}

// Role gets users
func (z *Context) RoleGet(params RoleGetParams) ([]RoleObject, int, error) {

	var result []RoleObject

	status, err := z.request("role.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// RoleCreate creates roles
func (z *Context) RoleCreate(params []RoleObject) ([]int, int, error) {

	var result roleCreateResult

	status, err := z.request("role.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.RoleIDs, status, nil
}

// RoleDelete deletes roles
func (z *Context) RoleDelete(roleIDs []int) ([]int, int, error) {

	var result roleDeleteResult

	status, err := z.request("role.delete", roleIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.RoleIDs, status, nil
}
