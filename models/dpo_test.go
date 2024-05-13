// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testDpos(t *testing.T) {
	t.Parallel()

	query := Dpos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDposDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
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

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDposQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Dpos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDposSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DpoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDposExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DpoExists(ctx, tx, o.DemandPartnerID)
	if err != nil {
		t.Errorf("Unable to check if Dpo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DpoExists to return true, but got false.")
	}
}

func testDposFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dpoFound, err := FindDpo(ctx, tx, o.DemandPartnerID)
	if err != nil {
		t.Error(err)
	}

	if dpoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDposBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Dpos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDposOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Dpos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDposAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dpoOne := &Dpo{}
	dpoTwo := &Dpo{}
	if err = randomize.Struct(seed, dpoOne, dpoDBTypes, false, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}
	if err = randomize.Struct(seed, dpoTwo, dpoDBTypes, false, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dpoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dpoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Dpos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDposCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dpoOne := &Dpo{}
	dpoTwo := &Dpo{}
	if err = randomize.Struct(seed, dpoOne, dpoDBTypes, false, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}
	if err = randomize.Struct(seed, dpoTwo, dpoDBTypes, false, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dpoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dpoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dpoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func dpoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
	*o = Dpo{}
	return nil
}

func testDposHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Dpo{}
	o := &Dpo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dpoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Dpo object: %s", err)
	}

	AddDpoHook(boil.BeforeInsertHook, dpoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dpoBeforeInsertHooks = []DpoHook{}

	AddDpoHook(boil.AfterInsertHook, dpoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dpoAfterInsertHooks = []DpoHook{}

	AddDpoHook(boil.AfterSelectHook, dpoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dpoAfterSelectHooks = []DpoHook{}

	AddDpoHook(boil.BeforeUpdateHook, dpoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dpoBeforeUpdateHooks = []DpoHook{}

	AddDpoHook(boil.AfterUpdateHook, dpoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dpoAfterUpdateHooks = []DpoHook{}

	AddDpoHook(boil.BeforeDeleteHook, dpoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dpoBeforeDeleteHooks = []DpoHook{}

	AddDpoHook(boil.AfterDeleteHook, dpoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dpoAfterDeleteHooks = []DpoHook{}

	AddDpoHook(boil.BeforeUpsertHook, dpoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dpoBeforeUpsertHooks = []DpoHook{}

	AddDpoHook(boil.AfterUpsertHook, dpoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dpoAfterUpsertHooks = []DpoHook{}
}

func testDposInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDposInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dpoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDpoToManyDemandPartnerDpoRules(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Dpo
	var b, c DpoRule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dpoRuleDBTypes, false, dpoRuleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dpoRuleDBTypes, false, dpoRuleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.DemandPartnerID = a.DemandPartnerID
	c.DemandPartnerID = a.DemandPartnerID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.DemandPartnerDpoRules().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.DemandPartnerID == b.DemandPartnerID {
			bFound = true
		}
		if v.DemandPartnerID == c.DemandPartnerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DpoSlice{&a}
	if err = a.L.LoadDemandPartnerDpoRules(ctx, tx, false, (*[]*Dpo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DemandPartnerDpoRules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.DemandPartnerDpoRules = nil
	if err = a.L.LoadDemandPartnerDpoRules(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DemandPartnerDpoRules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDpoToManyAddOpDemandPartnerDpoRules(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Dpo
	var b, c, d, e DpoRule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dpoDBTypes, false, strmangle.SetComplement(dpoPrimaryKeyColumns, dpoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DpoRule{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dpoRuleDBTypes, false, strmangle.SetComplement(dpoRulePrimaryKeyColumns, dpoRuleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DpoRule{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDemandPartnerDpoRules(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.DemandPartnerID != first.DemandPartnerID {
			t.Error("foreign key was wrong value", a.DemandPartnerID, first.DemandPartnerID)
		}
		if a.DemandPartnerID != second.DemandPartnerID {
			t.Error("foreign key was wrong value", a.DemandPartnerID, second.DemandPartnerID)
		}

		if first.R.DemandPartner != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.DemandPartner != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.DemandPartnerDpoRules[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.DemandPartnerDpoRules[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.DemandPartnerDpoRules().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDposReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
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

func testDposReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DpoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDposSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Dpos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dpoDBTypes = map[string]string{`DemandPartnerID`: `character varying`, `IsInclude`: `boolean`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_          = bytes.MinRead
)

func testDposUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dpoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dpoAllColumns) == len(dpoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDposSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dpoAllColumns) == len(dpoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Dpo{}
	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dpoDBTypes, true, dpoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dpoAllColumns, dpoPrimaryKeyColumns) {
		fields = dpoAllColumns
	} else {
		fields = strmangle.SetComplement(
			dpoAllColumns,
			dpoPrimaryKeyColumns,
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

	slice := DpoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDposUpsert(t *testing.T) {
	t.Parallel()

	if len(dpoAllColumns) == len(dpoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Dpo{}
	if err = randomize.Struct(seed, &o, dpoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Dpo: %s", err)
	}

	count, err := Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dpoDBTypes, false, dpoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Dpo: %s", err)
	}

	count, err = Dpos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}