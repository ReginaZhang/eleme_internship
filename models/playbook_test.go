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

func testPlaybooks(t *testing.T) {
	t.Parallel()

	query := Playbooks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testPlaybooksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = playbook.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlaybooksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Playbooks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlaybooksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PlaybookSlice{playbook}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testPlaybooksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := PlaybookExists(tx, playbook.ID)
	if err != nil {
		t.Errorf("Unable to check if Playbook exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PlaybookExistsG to return true, but got false.")
	}
}
func testPlaybooksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	playbookFound, err := FindPlaybook(tx, playbook.ID)
	if err != nil {
		t.Error(err)
	}

	if playbookFound == nil {
		t.Error("want a record, got nil")
	}
}
func testPlaybooksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Playbooks(tx).Bind(playbook); err != nil {
		t.Error(err)
	}
}

func testPlaybooksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Playbooks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPlaybooksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbookOne := &Playbook{}
	playbookTwo := &Playbook{}
	if err = randomize.Struct(seed, playbookOne, playbookDBTypes, false, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}
	if err = randomize.Struct(seed, playbookTwo, playbookDBTypes, false, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbookOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = playbookTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Playbooks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPlaybooksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	playbookOne := &Playbook{}
	playbookTwo := &Playbook{}
	if err = randomize.Struct(seed, playbookOne, playbookDBTypes, false, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}
	if err = randomize.Struct(seed, playbookTwo, playbookDBTypes, false, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbookOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = playbookTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func playbookBeforeInsertHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookAfterInsertHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookAfterSelectHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookBeforeUpdateHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookAfterUpdateHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookBeforeDeleteHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookAfterDeleteHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookBeforeUpsertHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func playbookAfterUpsertHook(e boil.Executor, o *Playbook) error {
	*o = Playbook{}
	return nil
}

func testPlaybooksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Playbook{}
	o := &Playbook{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, playbookDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Playbook object: %s", err)
	}

	AddPlaybookHook(boil.BeforeInsertHook, playbookBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	playbookBeforeInsertHooks = []PlaybookHook{}

	AddPlaybookHook(boil.AfterInsertHook, playbookAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	playbookAfterInsertHooks = []PlaybookHook{}

	AddPlaybookHook(boil.AfterSelectHook, playbookAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	playbookAfterSelectHooks = []PlaybookHook{}

	AddPlaybookHook(boil.BeforeUpdateHook, playbookBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	playbookBeforeUpdateHooks = []PlaybookHook{}

	AddPlaybookHook(boil.AfterUpdateHook, playbookAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	playbookAfterUpdateHooks = []PlaybookHook{}

	AddPlaybookHook(boil.BeforeDeleteHook, playbookBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	playbookBeforeDeleteHooks = []PlaybookHook{}

	AddPlaybookHook(boil.AfterDeleteHook, playbookAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	playbookAfterDeleteHooks = []PlaybookHook{}

	AddPlaybookHook(boil.BeforeUpsertHook, playbookBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	playbookBeforeUpsertHooks = []PlaybookHook{}

	AddPlaybookHook(boil.AfterUpsertHook, playbookAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	playbookAfterUpsertHooks = []PlaybookHook{}
}
func testPlaybooksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPlaybooksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx, playbookColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPlaybooksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = playbook.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPlaybooksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := PlaybookSlice{playbook}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testPlaybooksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Playbooks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	playbookDBTypes = map[string]string{`CreatedAt`: `timestamp`, `DRCCheckTime`: `timestamp`, `Entry`: `varchar`, `GitRepo`: `varchar`, `ID`: `bigint`, `Name`: `varchar`, `UpdatedAt`: `timestamp`}
	_               = bytes.MinRead
)

func testPlaybooksUpdate(t *testing.T) {
	t.Parallel()

	if len(playbookColumns) == len(playbookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	if err = playbook.Update(tx); err != nil {
		t.Error(err)
	}
}

func testPlaybooksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(playbookColumns) == len(playbookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	playbook := &Playbook{}
	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, playbook, playbookDBTypes, true, playbookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(playbookColumns, playbookPrimaryKeyColumns) {
		fields = playbookColumns
	} else {
		fields = strmangle.SetComplement(
			playbookColumns,
			playbookPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(playbook))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := PlaybookSlice{playbook}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testPlaybooksUpsert(t *testing.T) {
	t.Parallel()

	if len(playbookColumns) == len(playbookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	playbook := Playbook{}
	if err = randomize.Struct(seed, &playbook, playbookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = playbook.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Playbook: %s", err)
	}

	count, err := Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &playbook, playbookDBTypes, false, playbookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Playbook struct: %s", err)
	}

	if err = playbook.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Playbook: %s", err)
	}

	count, err = Playbooks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
