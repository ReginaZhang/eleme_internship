// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testApps(t *testing.T) {
	t.Parallel()

	query := Apps(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAppsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = app.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Apps(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AppSlice{app}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAppsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AppExists(tx, app.ID)
	if err != nil {
		t.Errorf("Unable to check if App exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AppExistsG to return true, but got false.")
	}
}
func testAppsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	appFound, err := FindApp(tx, app.ID)
	if err != nil {
		t.Error(err)
	}

	if appFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAppsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Apps(tx).Bind(app); err != nil {
		t.Error(err)
	}
}

func testAppsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Apps(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAppsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appOne := &App{}
	appTwo := &App{}
	if err = randomize.Struct(seed, appOne, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}
	if err = randomize.Struct(seed, appTwo, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = appTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Apps(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAppsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	appOne := &App{}
	appTwo := &App{}
	if err = randomize.Struct(seed, appOne, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}
	if err = randomize.Struct(seed, appTwo, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = appTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func appBeforeInsertHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appAfterInsertHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appAfterSelectHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appBeforeUpdateHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appAfterUpdateHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appBeforeDeleteHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appAfterDeleteHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appBeforeUpsertHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func appAfterUpsertHook(e boil.Executor, o *App) error {
	*o = App{}
	return nil
}

func testAppsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &App{}
	o := &App{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, appDBTypes, false); err != nil {
		t.Errorf("Unable to randomize App object: %s", err)
	}

	AddAppHook(boil.BeforeInsertHook, appBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	appBeforeInsertHooks = []AppHook{}

	AddAppHook(boil.AfterInsertHook, appAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	appAfterInsertHooks = []AppHook{}

	AddAppHook(boil.AfterSelectHook, appAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	appAfterSelectHooks = []AppHook{}

	AddAppHook(boil.BeforeUpdateHook, appBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	appBeforeUpdateHooks = []AppHook{}

	AddAppHook(boil.AfterUpdateHook, appAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	appAfterUpdateHooks = []AppHook{}

	AddAppHook(boil.BeforeDeleteHook, appBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	appBeforeDeleteHooks = []AppHook{}

	AddAppHook(boil.AfterDeleteHook, appAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	appAfterDeleteHooks = []AppHook{}

	AddAppHook(boil.BeforeUpsertHook, appBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	appBeforeUpsertHooks = []AppHook{}

	AddAppHook(boil.AfterUpsertHook, appAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	appAfterUpsertHooks = []AppHook{}
}
func testAppsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx, appColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = app.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAppsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AppSlice{app}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAppsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Apps(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	appDBTypes = map[string]string{`Appid`: `varchar`, `CreatedAt`: `timestamp`, `DRCCheckTime`: `timestamp`, `ID`: `bigint`, `UpdatedAt`: `timestamp`}
	_          = bytes.MinRead
)

func testAppsUpdate(t *testing.T) {
	t.Parallel()

	if len(appColumns) == len(appPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	if err = app.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAppsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(appColumns) == len(appPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	app := &App{}
	if err = randomize.Struct(seed, app, appDBTypes, true, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, app, appDBTypes, true, appPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(appColumns, appPrimaryKeyColumns) {
		fields = appColumns
	} else {
		fields = strmangle.SetComplement(
			appColumns,
			appPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(app))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AppSlice{app}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAppsUpsert(t *testing.T) {
	t.Parallel()

	if len(appColumns) == len(appPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	app := App{}
	if err = randomize.Struct(seed, &app, appDBTypes, true); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = app.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert App: %s", err)
	}

	count, err := Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &app, appDBTypes, false, appPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	if err = app.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert App: %s", err)
	}

	count, err = Apps(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
