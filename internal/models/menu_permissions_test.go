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

func testMenuPermissions(t *testing.T) {
	t.Parallel()

	query := MenuPermissions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMenuPermissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
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

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMenuPermissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MenuPermissions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMenuPermissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MenuPermissionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMenuPermissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MenuPermissionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MenuPermission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MenuPermissionExists to return true, but got false.")
	}
}

func testMenuPermissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	menuPermissionFound, err := FindMenuPermission(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if menuPermissionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMenuPermissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MenuPermissions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMenuPermissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MenuPermissions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMenuPermissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	menuPermissionOne := &MenuPermission{}
	menuPermissionTwo := &MenuPermission{}
	if err = randomize.Struct(seed, menuPermissionOne, menuPermissionDBTypes, false, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, menuPermissionTwo, menuPermissionDBTypes, false, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = menuPermissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = menuPermissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MenuPermissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMenuPermissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	menuPermissionOne := &MenuPermission{}
	menuPermissionTwo := &MenuPermission{}
	if err = randomize.Struct(seed, menuPermissionOne, menuPermissionDBTypes, false, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, menuPermissionTwo, menuPermissionDBTypes, false, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = menuPermissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = menuPermissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func menuPermissionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func menuPermissionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MenuPermission) error {
	*o = MenuPermission{}
	return nil
}

func testMenuPermissionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MenuPermission{}
	o := &MenuPermission{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MenuPermission object: %s", err)
	}

	AddMenuPermissionHook(boil.BeforeInsertHook, menuPermissionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	menuPermissionBeforeInsertHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.AfterInsertHook, menuPermissionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	menuPermissionAfterInsertHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.AfterSelectHook, menuPermissionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	menuPermissionAfterSelectHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.BeforeUpdateHook, menuPermissionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	menuPermissionBeforeUpdateHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.AfterUpdateHook, menuPermissionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	menuPermissionAfterUpdateHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.BeforeDeleteHook, menuPermissionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	menuPermissionBeforeDeleteHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.AfterDeleteHook, menuPermissionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	menuPermissionAfterDeleteHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.BeforeUpsertHook, menuPermissionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	menuPermissionBeforeUpsertHooks = []MenuPermissionHook{}

	AddMenuPermissionHook(boil.AfterUpsertHook, menuPermissionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	menuPermissionAfterUpsertHooks = []MenuPermissionHook{}
}

func testMenuPermissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMenuPermissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(menuPermissionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMenuPermissionToManyRoleMenuPermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b, c RoleMenuPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, roleMenuPermissionDBTypes, false, roleMenuPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleMenuPermissionDBTypes, false, roleMenuPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.MenuPermissionID, a.ID)
	queries.Assign(&c.MenuPermissionID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RoleMenuPermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.MenuPermissionID, b.MenuPermissionID) {
			bFound = true
		}
		if queries.Equal(v.MenuPermissionID, c.MenuPermissionID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MenuPermissionSlice{&a}
	if err = a.L.LoadRoleMenuPermissions(ctx, tx, false, (*[]*MenuPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleMenuPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RoleMenuPermissions = nil
	if err = a.L.LoadRoleMenuPermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleMenuPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMenuPermissionToManyAddOpRoleMenuPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b, c, d, e RoleMenuPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, false, strmangle.SetComplement(menuPermissionPrimaryKeyColumns, menuPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RoleMenuPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, roleMenuPermissionDBTypes, false, strmangle.SetComplement(roleMenuPermissionPrimaryKeyColumns, roleMenuPermissionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*RoleMenuPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRoleMenuPermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.MenuPermissionID) {
			t.Error("foreign key was wrong value", a.ID, first.MenuPermissionID)
		}
		if !queries.Equal(a.ID, second.MenuPermissionID) {
			t.Error("foreign key was wrong value", a.ID, second.MenuPermissionID)
		}

		if first.R.MenuPermission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.MenuPermission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RoleMenuPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RoleMenuPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RoleMenuPermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMenuPermissionToManySetOpRoleMenuPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b, c, d, e RoleMenuPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, false, strmangle.SetComplement(menuPermissionPrimaryKeyColumns, menuPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RoleMenuPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, roleMenuPermissionDBTypes, false, strmangle.SetComplement(roleMenuPermissionPrimaryKeyColumns, roleMenuPermissionColumnsWithoutDefault)...); err != nil {
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

	err = a.SetRoleMenuPermissions(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RoleMenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetRoleMenuPermissions(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RoleMenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MenuPermissionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MenuPermissionID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.MenuPermissionID) {
		t.Error("foreign key was wrong value", a.ID, d.MenuPermissionID)
	}
	if !queries.Equal(a.ID, e.MenuPermissionID) {
		t.Error("foreign key was wrong value", a.ID, e.MenuPermissionID)
	}

	if b.R.MenuPermission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.MenuPermission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.MenuPermission != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.MenuPermission != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.RoleMenuPermissions[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.RoleMenuPermissions[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMenuPermissionToManyRemoveOpRoleMenuPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b, c, d, e RoleMenuPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, false, strmangle.SetComplement(menuPermissionPrimaryKeyColumns, menuPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RoleMenuPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, roleMenuPermissionDBTypes, false, strmangle.SetComplement(roleMenuPermissionPrimaryKeyColumns, roleMenuPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddRoleMenuPermissions(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RoleMenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveRoleMenuPermissions(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RoleMenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MenuPermissionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MenuPermissionID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.MenuPermission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.MenuPermission != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.MenuPermission != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.MenuPermission != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.RoleMenuPermissions) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.RoleMenuPermissions[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.RoleMenuPermissions[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMenuPermissionToOneMenuUsingMenu(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MenuPermission
	var foreign Menu

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, menuDBTypes, false, menuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Menu struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.MenuID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Menu().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddMenuHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Menu) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := MenuPermissionSlice{&local}
	if err = local.L.LoadMenu(ctx, tx, false, (*[]*MenuPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Menu == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Menu = nil
	if err = local.L.LoadMenu(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Menu == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testMenuPermissionToOnePermissionUsingPermission(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MenuPermission
	var foreign Permission

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, permissionDBTypes, false, permissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Permission struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.PermissionID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Permission().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddPermissionHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Permission) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := MenuPermissionSlice{&local}
	if err = local.L.LoadPermission(ctx, tx, false, (*[]*MenuPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Permission == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Permission = nil
	if err = local.L.LoadPermission(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Permission == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testMenuPermissionToOneSetOpMenuUsingMenu(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b, c Menu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, false, strmangle.SetComplement(menuPermissionPrimaryKeyColumns, menuPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Menu{&b, &c} {
		err = a.SetMenu(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Menu != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MenuPermissions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.MenuID, x.ID) {
			t.Error("foreign key was wrong value", a.MenuID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MenuID))
		reflect.Indirect(reflect.ValueOf(&a.MenuID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.MenuID, x.ID) {
			t.Error("foreign key was wrong value", a.MenuID, x.ID)
		}
	}
}

func testMenuPermissionToOneRemoveOpMenuUsingMenu(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b Menu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, false, strmangle.SetComplement(menuPermissionPrimaryKeyColumns, menuPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, menuDBTypes, false, strmangle.SetComplement(menuPrimaryKeyColumns, menuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetMenu(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveMenu(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Menu().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Menu != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.MenuID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.MenuPermissions) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testMenuPermissionToOneSetOpPermissionUsingPermission(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b, c Permission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, false, strmangle.SetComplement(menuPermissionPrimaryKeyColumns, menuPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, permissionDBTypes, false, strmangle.SetComplement(permissionPrimaryKeyColumns, permissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, permissionDBTypes, false, strmangle.SetComplement(permissionPrimaryKeyColumns, permissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Permission{&b, &c} {
		err = a.SetPermission(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Permission != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MenuPermissions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.PermissionID, x.ID) {
			t.Error("foreign key was wrong value", a.PermissionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PermissionID))
		reflect.Indirect(reflect.ValueOf(&a.PermissionID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.PermissionID, x.ID) {
			t.Error("foreign key was wrong value", a.PermissionID, x.ID)
		}
	}
}

func testMenuPermissionToOneRemoveOpPermissionUsingPermission(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MenuPermission
	var b Permission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, menuPermissionDBTypes, false, strmangle.SetComplement(menuPermissionPrimaryKeyColumns, menuPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, permissionDBTypes, false, strmangle.SetComplement(permissionPrimaryKeyColumns, permissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPermission(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePermission(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Permission().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Permission != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.PermissionID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.MenuPermissions) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testMenuPermissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
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

func testMenuPermissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MenuPermissionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMenuPermissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MenuPermissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	menuPermissionDBTypes = map[string]string{`ID`: `integer`, `MenuID`: `smallint`, `PermissionID`: `smallint`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_                     = bytes.MinRead
)

func testMenuPermissionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(menuPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(menuPermissionAllColumns) == len(menuPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMenuPermissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(menuPermissionAllColumns) == len(menuPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MenuPermission{}
	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, menuPermissionDBTypes, true, menuPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(menuPermissionAllColumns, menuPermissionPrimaryKeyColumns) {
		fields = menuPermissionAllColumns
	} else {
		fields = strmangle.SetComplement(
			menuPermissionAllColumns,
			menuPermissionPrimaryKeyColumns,
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

	slice := MenuPermissionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMenuPermissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(menuPermissionAllColumns) == len(menuPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MenuPermission{}
	if err = randomize.Struct(seed, &o, menuPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MenuPermission: %s", err)
	}

	count, err := MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, menuPermissionDBTypes, false, menuPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MenuPermission struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MenuPermission: %s", err)
	}

	count, err = MenuPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}