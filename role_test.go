package zabbix

import (
	"reflect"
	"testing"
)

const (
	testRoleName = "Test Role"
)

func TestRoleCRUD(t *testing.T) {

	var z Context

	// Login
	loginTest(&z, t)
	defer logoutTest(&z, t)

	// Create and delete
	roleIDs := testRoleCreate(t, z)
	defer testRoleDelete(t, z, roleIDs)

	// Get
	testRoleGet(t, z)
}

func testRoleCreate(t *testing.T, z Context) []int {

	roleIDs, _, err := z.RoleCreate([]RoleObject{
		{
			Name: testRoleName,
			Type: RoleTypeUser,

			Rules: RoleRuleObject{
				UIs: []UIElement{
					{
						Name:   "monitoring.hosts",
						Status: 0,
					},
				},
				UIsDefaultAccess:         UIsAccessTypeDeny,
				Modules:                  []ModuleObject{},
				ModulesDefaultAccess:     ModulesAccessTypeDeny,
				APIAccess:                APIAccessTypeDeny,
				APIMode:                  APIModeAccessList,
				APIMethods:               []string{"user.get"},
				RoleActions:              []RoleActionObject{},
				RoleActionsDefaultAccess: RoleActionStatusDisabled,
			},
		},
	})
	if err != nil {
		t.Fatal("Role create error:", err)
	}

	if len(roleIDs) == 0 {
		t.Fatal("Role create error: empty IDs array")
	}

	t.Logf("Role create: success")

	return roleIDs
}

func testRoleDelete(t *testing.T, z Context, roleIDs []int) []int {

	rDeletedIDs, _, err := z.RoleDelete(roleIDs)
	if err != nil {
		t.Fatal("Roles delete error:", err)
	}

	if len(rDeletedIDs) == 0 {
		t.Fatal("Role delete error: empty IDs array")
	}

	if reflect.DeepEqual(rDeletedIDs, roleIDs) == false {
		t.Fatal("Role delete error: IDs arrays for created and deleted roles are mismatch")
	}

	t.Logf("Role delete: success")

	return rDeletedIDs
}

func testRoleGet(t *testing.T, z Context) []RoleObject {

	rObjects, _, err := z.RoleGet(RoleGetParams{
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name": testRoleName,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Role get error:", err)
	} else {
		if len(rObjects) == 0 {
			t.Error("Role get error: unable to find created role")
		} else {
			t.Logf("Role get: success")
		}
	}

	return rObjects
}
