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

func testAppPlaybooks(t *testing.T) {
	t.Parallel()

	query := AppPlaybooks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAppPlaybooksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = appPlaybook.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppPlaybooksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AppPlaybooks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppPlaybooksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AppPlaybookSlice{appPlaybook}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAppPlaybooksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AppPlaybookExists(tx, appPlaybook.ID)
	if err != nil {
		t.Errorf("Unable to check if AppPlaybook exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AppPlaybookExistsG to return true, but got false.")
	}
}
func testAppPlaybooksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	appPlaybookFound, err := FindAppPlaybook(tx, appPlaybook.ID)
	if err != nil {
		t.Error(err)
	}

	if appPlaybookFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAppPlaybooksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AppPlaybooks(tx).Bind(appPlaybook); err != nil {
		t.Error(err)
	}
}

func testAppPlaybooksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AppPlaybooks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAppPlaybooksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybookOne := &AppPlaybook{}
	appPlaybookTwo := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybookOne, appPlaybookDBTypes, false, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}
	if err = randomize.Struct(seed, appPlaybookTwo, appPlaybookDBTypes, false, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybookOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = appPlaybookTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AppPlaybooks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAppPlaybooksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	appPlaybookOne := &AppPlaybook{}
	appPlaybookTwo := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybookOne, appPlaybookDBTypes, false, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}
	if err = randomize.Struct(seed, appPlaybookTwo, appPlaybookDBTypes, false, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybookOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = appPlaybookTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func appPlaybookBeforeInsertHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookAfterInsertHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookAfterSelectHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookBeforeUpdateHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookAfterUpdateHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookBeforeDeleteHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookAfterDeleteHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookBeforeUpsertHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func appPlaybookAfterUpsertHook(e boil.Executor, o *AppPlaybook) error {
	*o = AppPlaybook{}
	return nil
}

func testAppPlaybooksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AppPlaybook{}
	o := &AppPlaybook{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, appPlaybookDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AppPlaybook object: %s", err)
	}

	AddAppPlaybookHook(boil.BeforeInsertHook, appPlaybookBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	appPlaybookBeforeInsertHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.AfterInsertHook, appPlaybookAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	appPlaybookAfterInsertHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.AfterSelectHook, appPlaybookAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	appPlaybookAfterSelectHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.BeforeUpdateHook, appPlaybookBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	appPlaybookBeforeUpdateHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.AfterUpdateHook, appPlaybookAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	appPlaybookAfterUpdateHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.BeforeDeleteHook, appPlaybookBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	appPlaybookBeforeDeleteHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.AfterDeleteHook, appPlaybookAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	appPlaybookAfterDeleteHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.BeforeUpsertHook, appPlaybookBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	appPlaybookBeforeUpsertHooks = []AppPlaybookHook{}

	AddAppPlaybookHook(boil.AfterUpsertHook, appPlaybookAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	appPlaybookAfterUpsertHooks = []AppPlaybookHook{}
}
func testAppPlaybooksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppPlaybooksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx, appPlaybookColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppPlaybooksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = appPlaybook.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAppPlaybooksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AppPlaybookSlice{appPlaybook}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAppPlaybooksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AppPlaybooks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	appPlaybookDBTypes = map[string]string{`AppID`: `bigint`, `CreatedAt`: `timestamp`, `DRCCheckTime`: `timestamp`, `ID`: `bigint`, `PlaybookID`: `bigint`, `UpdatedAt`: `timestamp`}
	_                  = bytes.MinRead
)

func testAppPlaybooksUpdate(t *testing.T) {
	t.Parallel()

	if len(appPlaybookColumns) == len(appPlaybookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	if err = appPlaybook.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAppPlaybooksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(appPlaybookColumns) == len(appPlaybookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	appPlaybook := &AppPlaybook{}
	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, appPlaybook, appPlaybookDBTypes, true, appPlaybookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(appPlaybookColumns, appPlaybookPrimaryKeyColumns) {
		fields = appPlaybookColumns
	} else {
		fields = strmangle.SetComplement(
			appPlaybookColumns,
			appPlaybookPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(appPlaybook))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AppPlaybookSlice{appPlaybook}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAppPlaybooksUpsert(t *testing.T) {
	t.Parallel()

	if len(appPlaybookColumns) == len(appPlaybookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	appPlaybook := AppPlaybook{}
	if err = randomize.Struct(seed, &appPlaybook, appPlaybookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = appPlaybook.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert AppPlaybook: %s", err)
	}

	count, err := AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &appPlaybook, appPlaybookDBTypes, false, appPlaybookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppPlaybook struct: %s", err)
	}

	if err = appPlaybook.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert AppPlaybook: %s", err)
	}

	count, err = AppPlaybooks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
