// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Permissions", testPermissions)
	t.Run("RolePermissions", testRolePermissions)
	t.Run("Roles", testRoles)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Permissions", testPermissionsDelete)
	t.Run("RolePermissions", testRolePermissionsDelete)
	t.Run("Roles", testRolesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Permissions", testPermissionsQueryDeleteAll)
	t.Run("RolePermissions", testRolePermissionsQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Permissions", testPermissionsSliceDeleteAll)
	t.Run("RolePermissions", testRolePermissionsSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Permissions", testPermissionsExists)
	t.Run("RolePermissions", testRolePermissionsExists)
	t.Run("Roles", testRolesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Permissions", testPermissionsFind)
	t.Run("RolePermissions", testRolePermissionsFind)
	t.Run("Roles", testRolesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Permissions", testPermissionsBind)
	t.Run("RolePermissions", testRolePermissionsBind)
	t.Run("Roles", testRolesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Permissions", testPermissionsOne)
	t.Run("RolePermissions", testRolePermissionsOne)
	t.Run("Roles", testRolesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Permissions", testPermissionsAll)
	t.Run("RolePermissions", testRolePermissionsAll)
	t.Run("Roles", testRolesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Permissions", testPermissionsCount)
	t.Run("RolePermissions", testRolePermissionsCount)
	t.Run("Roles", testRolesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Permissions", testPermissionsHooks)
	t.Run("RolePermissions", testRolePermissionsHooks)
	t.Run("Roles", testRolesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Permissions", testPermissionsInsert)
	t.Run("Permissions", testPermissionsInsertWhitelist)
	t.Run("RolePermissions", testRolePermissionsInsert)
	t.Run("RolePermissions", testRolePermissionsInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Permissions", testPermissionsReload)
	t.Run("RolePermissions", testRolePermissionsReload)
	t.Run("Roles", testRolesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Permissions", testPermissionsReloadAll)
	t.Run("RolePermissions", testRolePermissionsReloadAll)
	t.Run("Roles", testRolesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Permissions", testPermissionsSelect)
	t.Run("RolePermissions", testRolePermissionsSelect)
	t.Run("Roles", testRolesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Permissions", testPermissionsUpdate)
	t.Run("RolePermissions", testRolePermissionsUpdate)
	t.Run("Roles", testRolesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Permissions", testPermissionsSliceUpdateAll)
	t.Run("RolePermissions", testRolePermissionsSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
