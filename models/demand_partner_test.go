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

func testDemandPartners(t *testing.T) {
	t.Parallel()

	query := DemandPartners()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDemandPartnersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
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

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDemandPartnersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DemandPartners().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDemandPartnersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DemandPartnerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDemandPartnersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DemandPartnerExists(ctx, tx, o.DemandPartnerID)
	if err != nil {
		t.Errorf("Unable to check if DemandPartner exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DemandPartnerExists to return true, but got false.")
	}
}

func testDemandPartnersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	demandPartnerFound, err := FindDemandPartner(ctx, tx, o.DemandPartnerID)
	if err != nil {
		t.Error(err)
	}

	if demandPartnerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDemandPartnersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DemandPartners().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDemandPartnersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DemandPartners().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDemandPartnersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	demandPartnerOne := &DemandPartner{}
	demandPartnerTwo := &DemandPartner{}
	if err = randomize.Struct(seed, demandPartnerOne, demandPartnerDBTypes, false, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}
	if err = randomize.Struct(seed, demandPartnerTwo, demandPartnerDBTypes, false, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = demandPartnerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = demandPartnerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DemandPartners().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDemandPartnersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	demandPartnerOne := &DemandPartner{}
	demandPartnerTwo := &DemandPartner{}
	if err = randomize.Struct(seed, demandPartnerOne, demandPartnerDBTypes, false, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}
	if err = randomize.Struct(seed, demandPartnerTwo, demandPartnerDBTypes, false, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = demandPartnerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = demandPartnerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func demandPartnerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func demandPartnerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DemandPartner) error {
	*o = DemandPartner{}
	return nil
}

func testDemandPartnersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DemandPartner{}
	o := &DemandPartner{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DemandPartner object: %s", err)
	}

	AddDemandPartnerHook(boil.BeforeInsertHook, demandPartnerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	demandPartnerBeforeInsertHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.AfterInsertHook, demandPartnerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	demandPartnerAfterInsertHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.AfterSelectHook, demandPartnerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	demandPartnerAfterSelectHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.BeforeUpdateHook, demandPartnerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	demandPartnerBeforeUpdateHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.AfterUpdateHook, demandPartnerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	demandPartnerAfterUpdateHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.BeforeDeleteHook, demandPartnerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	demandPartnerBeforeDeleteHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.AfterDeleteHook, demandPartnerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	demandPartnerAfterDeleteHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.BeforeUpsertHook, demandPartnerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	demandPartnerBeforeUpsertHooks = []DemandPartnerHook{}

	AddDemandPartnerHook(boil.AfterUpsertHook, demandPartnerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	demandPartnerAfterUpsertHooks = []DemandPartnerHook{}
}

func testDemandPartnersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDemandPartnersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(demandPartnerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDemandPartnerToManyDemandParnterPlacements(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DemandPartner
	var b, c DemandParnterPlacement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, demandParnterPlacementDBTypes, false, demandParnterPlacementColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, demandParnterPlacementDBTypes, false, demandParnterPlacementColumnsWithDefault...); err != nil {
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

	check, err := a.DemandParnterPlacements().All(ctx, tx)
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

	slice := DemandPartnerSlice{&a}
	if err = a.L.LoadDemandParnterPlacements(ctx, tx, false, (*[]*DemandPartner)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DemandParnterPlacements); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.DemandParnterPlacements = nil
	if err = a.L.LoadDemandParnterPlacements(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DemandParnterPlacements); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDemandPartnerToManyAddOpDemandParnterPlacements(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DemandPartner
	var b, c, d, e DemandParnterPlacement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, demandPartnerDBTypes, false, strmangle.SetComplement(demandPartnerPrimaryKeyColumns, demandPartnerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DemandParnterPlacement{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, demandParnterPlacementDBTypes, false, strmangle.SetComplement(demandParnterPlacementPrimaryKeyColumns, demandParnterPlacementColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DemandParnterPlacement{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDemandParnterPlacements(ctx, tx, i != 0, x...)
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

		if a.R.DemandParnterPlacements[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.DemandParnterPlacements[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.DemandParnterPlacements().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDemandPartnersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
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

func testDemandPartnersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DemandPartnerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDemandPartnersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DemandPartners().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	demandPartnerDBTypes = map[string]string{`DemandPartnerID`: `character varying`, `Name`: `character varying`, `IntegrationType`: `character varying`}
	_                    = bytes.MinRead
)

func testDemandPartnersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(demandPartnerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(demandPartnerAllColumns) == len(demandPartnerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDemandPartnersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(demandPartnerAllColumns) == len(demandPartnerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DemandPartner{}
	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, demandPartnerDBTypes, true, demandPartnerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(demandPartnerAllColumns, demandPartnerPrimaryKeyColumns) {
		fields = demandPartnerAllColumns
	} else {
		fields = strmangle.SetComplement(
			demandPartnerAllColumns,
			demandPartnerPrimaryKeyColumns,
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

	slice := DemandPartnerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDemandPartnersUpsert(t *testing.T) {
	t.Parallel()

	if len(demandPartnerAllColumns) == len(demandPartnerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DemandPartner{}
	if err = randomize.Struct(seed, &o, demandPartnerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DemandPartner: %s", err)
	}

	count, err := DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, demandPartnerDBTypes, false, demandPartnerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DemandPartner struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DemandPartner: %s", err)
	}

	count, err = DemandPartners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
