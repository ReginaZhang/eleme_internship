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

func testAppInventories(t *testing.T) {
	t.Parallel()

	query := AppInventories(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAppInventoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = appInventory.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppInventoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AppInventories(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppInventoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AppInventorySlice{appInventory}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAppInventoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AppInventoryExists(tx, appInventory.ID)
	if err != nil {
		t.Errorf("Unable to check if AppInventory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AppInventoryExistsG to return true, but got false.")
	}
}
func testAppInventoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	appInventoryFound, err := FindAppInventory(tx, appInventory.ID)
	if err != nil {
		t.Error(err)
	}

	if appInventoryFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAppInventoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AppInventories(tx).Bind(appInventory); err != nil {
		t.Error(err)
	}
}

func testAppInventoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AppInventories(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAppInventoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventoryOne := &AppInventory{}
	appInventoryTwo := &AppInventory{}
	if err = randomize.Struct(seed, appInventoryOne, appInventoryDBTypes, false, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}
	if err = randomize.Struct(seed, appInventoryTwo, appInventoryDBTypes, false, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = appInventoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AppInventories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAppInventoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	appInventoryOne := &AppInventory{}
	appInventoryTwo := &AppInventory{}
	if err = randomize.Struct(seed, appInventoryOne, appInventoryDBTypes, false, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}
	if err = randomize.Struct(seed, appInventoryTwo, appInventoryDBTypes, false, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventoryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = appInventoryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func appInventoryBeforeInsertHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryAfterInsertHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryAfterSelectHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryBeforeUpdateHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryAfterUpdateHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryBeforeDeleteHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryAfterDeleteHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryBeforeUpsertHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func appInventoryAfterUpsertHook(e boil.Executor, o *AppInventory) error {
	*o = AppInventory{}
	return nil
}

func testAppInventoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AppInventory{}
	o := &AppInventory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, appInventoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AppInventory object: %s", err)
	}

	AddAppInventoryHook(boil.BeforeInsertHook, appInventoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	appInventoryBeforeInsertHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.AfterInsertHook, appInventoryAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	appInventoryAfterInsertHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.AfterSelectHook, appInventoryAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	appInventoryAfterSelectHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.BeforeUpdateHook, appInventoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	appInventoryBeforeUpdateHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.AfterUpdateHook, appInventoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	appInventoryAfterUpdateHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.BeforeDeleteHook, appInventoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	appInventoryBeforeDeleteHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.AfterDeleteHook, appInventoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	appInventoryAfterDeleteHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.BeforeUpsertHook, appInventoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	appInventoryBeforeUpsertHooks = []AppInventoryHook{}

	AddAppInventoryHook(boil.AfterUpsertHook, appInventoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	appInventoryAfterUpsertHooks = []AppInventoryHook{}
}
func testAppInventoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppInventoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx, appInventoryColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppInventoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = appInventory.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAppInventoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AppInventorySlice{appInventory}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAppInventoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AppInventories(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	appInventoryDBTypes = map[string]string{`AppID`: `bigint`, `CreatedAt`: `timestamp`, `DRCCheckTime`: `timestamp`, `ID`: `bigint`, `InventoryID`: `bigint`, `UpdatedAt`: `timestamp`}
	_                   = bytes.MinRead
)

func testAppInventoriesUpdate(t *testing.T) {
	t.Parallel()

	if len(appInventoryColumns) == len(appInventoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	if err = appInventory.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAppInventoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(appInventoryColumns) == len(appInventoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	appInventory := &AppInventory{}
	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, appInventory, appInventoryDBTypes, true, appInventoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(appInventoryColumns, appInventoryPrimaryKeyColumns) {
		fields = appInventoryColumns
	} else {
		fields = strmangle.SetComplement(
			appInventoryColumns,
			appInventoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(appInventory))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AppInventorySlice{appInventory}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAppInventoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(appInventoryColumns) == len(appInventoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	appInventory := AppInventory{}
	if err = randomize.Struct(seed, &appInventory, appInventoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appInventory.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert AppInventory: %s", err)
	}

	count, err := AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &appInventory, appInventoryDBTypes, false, appInventoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppInventory struct: %s", err)
	}

	if err = appInventory.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert AppInventory: %s", err)
	}

	count, err = AppInventories(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
