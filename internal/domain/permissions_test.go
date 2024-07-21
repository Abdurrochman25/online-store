// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPermissions(t *testing.T) {
	t.Parallel()

	query := Permissions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPermissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPermissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Permissions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPermissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PermissionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPermissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PermissionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Permission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PermissionExists to return true, but got false.")
	}
}

func testPermissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	permissionFound, err := FindPermission(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if permissionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPermissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Permissions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPermissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Permissions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPermissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	permissionOne := &Permission{}
	permissionTwo := &Permission{}
	if err = randomize.Struct(seed, permissionOne, permissionDBTypes, false, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}
	if err = randomize.Struct(seed, permissionTwo, permissionDBTypes, false, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = permissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = permissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Permissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPermissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	permissionOne := &Permission{}
	permissionTwo := &Permission{}
	if err = randomize.Struct(seed, permissionOne, permissionDBTypes, false, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}
	if err = randomize.Struct(seed, permissionTwo, permissionDBTypes, false, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = permissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = permissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func permissionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func permissionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
	*o = Permission{}
	return nil
}

func testPermissionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Permission{}
	o := &Permission{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, permissionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Permission object: %s", err)
	}

	AddPermissionHook(boil.BeforeInsertHook, permissionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	permissionBeforeInsertHooks = []PermissionHook{}

	AddPermissionHook(boil.AfterInsertHook, permissionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	permissionAfterInsertHooks = []PermissionHook{}

	AddPermissionHook(boil.AfterSelectHook, permissionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	permissionAfterSelectHooks = []PermissionHook{}

	AddPermissionHook(boil.BeforeUpdateHook, permissionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	permissionBeforeUpdateHooks = []PermissionHook{}

	AddPermissionHook(boil.AfterUpdateHook, permissionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	permissionAfterUpdateHooks = []PermissionHook{}

	AddPermissionHook(boil.BeforeDeleteHook, permissionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	permissionBeforeDeleteHooks = []PermissionHook{}

	AddPermissionHook(boil.AfterDeleteHook, permissionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	permissionAfterDeleteHooks = []PermissionHook{}

	AddPermissionHook(boil.BeforeUpsertHook, permissionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	permissionBeforeUpsertHooks = []PermissionHook{}

	AddPermissionHook(boil.AfterUpsertHook, permissionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	permissionAfterUpsertHooks = []PermissionHook{}
}

func testPermissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPermissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(permissionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPermissionToManyRolePermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Permission
	var b, c RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, rolePermissionDBTypes, false, rolePermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, rolePermissionDBTypes, false, rolePermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.PermissionID, a.ID)
	queries.Assign(&c.PermissionID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RolePermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.PermissionID, b.PermissionID) {
			bFound = true
		}
		if queries.Equal(v.PermissionID, c.PermissionID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PermissionSlice{&a}
	if err = a.L.LoadRolePermissions(ctx, tx, false, (*[]*Permission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RolePermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RolePermissions = nil
	if err = a.L.LoadRolePermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RolePermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPermissionToManyAddOpRolePermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Permission
	var b, c, d, e RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, permissionDBTypes, false, strmangle.SetComplement(permissionPrimaryKeyColumns, permissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RolePermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rolePermissionDBTypes, false, strmangle.SetComplement(rolePermissionPrimaryKeyColumns, rolePermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*RolePermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRolePermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.PermissionID) {
			t.Error("foreign key was wrong value", a.ID, first.PermissionID)
		}
		if !queries.Equal(a.ID, second.PermissionID) {
			t.Error("foreign key was wrong value", a.ID, second.PermissionID)
		}

		if first.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RolePermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RolePermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RolePermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPermissionToManySetOpRolePermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Permission
	var b, c, d, e RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, permissionDBTypes, false, strmangle.SetComplement(permissionPrimaryKeyColumns, permissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RolePermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rolePermissionDBTypes, false, strmangle.SetComplement(rolePermissionPrimaryKeyColumns, rolePermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetRolePermissions(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetRolePermissions(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PermissionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PermissionID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.PermissionID) {
		t.Error("foreign key was wrong value", a.ID, d.PermissionID)
	}
	if !queries.Equal(a.ID, e.PermissionID) {
		t.Error("foreign key was wrong value", a.ID, e.PermissionID)
	}

	if b.R.Permission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Permission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Permission != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Permission != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.RolePermissions[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.RolePermissions[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPermissionToManyRemoveOpRolePermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Permission
	var b, c, d, e RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, permissionDBTypes, false, strmangle.SetComplement(permissionPrimaryKeyColumns, permissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RolePermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rolePermissionDBTypes, false, strmangle.SetComplement(rolePermissionPrimaryKeyColumns, rolePermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddRolePermissions(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveRolePermissions(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.PermissionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.PermissionID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Permission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Permission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Permission != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Permission != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.RolePermissions) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.RolePermissions[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.RolePermissions[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPermissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPermissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PermissionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPermissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Permissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	permissionDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_                 = bytes.MinRead
)

func testPermissionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(permissionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(permissionAllColumns) == len(permissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPermissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(permissionAllColumns) == len(permissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Permission{}
	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, permissionDBTypes, true, permissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(permissionAllColumns, permissionPrimaryKeyColumns) {
		fields = permissionAllColumns
	} else {
		fields = strmangle.SetComplement(
			permissionAllColumns,
			permissionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PermissionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPermissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(permissionAllColumns) == len(permissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Permission{}
	if err = randomize.Struct(seed, &o, permissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Permission: %s", err)
	}

	count, err := Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, permissionDBTypes, false, permissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Permission: %s", err)
	}

	count, err = Permissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
