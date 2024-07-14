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

func testDpoRules(t *testing.T) {
	t.Parallel()

	query := DpoRules()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDpoRulesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
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

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDpoRulesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DpoRules().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDpoRulesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DpoRuleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDpoRulesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DpoRuleExists(ctx, tx, o.RuleID)
	if err != nil {
		t.Errorf("Unable to check if DpoRule exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DpoRuleExists to return true, but got false.")
	}
}

func testDpoRulesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dpoRuleFound, err := FindDpoRule(ctx, tx, o.RuleID)
	if err != nil {
		t.Error(err)
	}

	if dpoRuleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDpoRulesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DpoRules().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDpoRulesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DpoRules().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDpoRulesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dpoRuleOne := &DpoRule{}
	dpoRuleTwo := &DpoRule{}
	if err = randomize.Struct(seed, dpoRuleOne, dpoRuleDBTypes, false, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}
	if err = randomize.Struct(seed, dpoRuleTwo, dpoRuleDBTypes, false, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dpoRuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dpoRuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DpoRules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDpoRulesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dpoRuleOne := &DpoRule{}
	dpoRuleTwo := &DpoRule{}
	if err = randomize.Struct(seed, dpoRuleOne, dpoRuleDBTypes, false, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}
	if err = randomize.Struct(seed, dpoRuleTwo, dpoRuleDBTypes, false, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dpoRuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dpoRuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dpoRuleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func dpoRuleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DpoRule) error {
	*o = DpoRule{}
	return nil
}

func testDpoRulesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DpoRule{}
	o := &DpoRule{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DpoRule object: %s", err)
	}

	AddDpoRuleHook(boil.BeforeInsertHook, dpoRuleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dpoRuleBeforeInsertHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.AfterInsertHook, dpoRuleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dpoRuleAfterInsertHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.AfterSelectHook, dpoRuleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dpoRuleAfterSelectHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.BeforeUpdateHook, dpoRuleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dpoRuleBeforeUpdateHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.AfterUpdateHook, dpoRuleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dpoRuleAfterUpdateHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.BeforeDeleteHook, dpoRuleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dpoRuleBeforeDeleteHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.AfterDeleteHook, dpoRuleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dpoRuleAfterDeleteHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.BeforeUpsertHook, dpoRuleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dpoRuleBeforeUpsertHooks = []DpoRuleHook{}

	AddDpoRuleHook(boil.AfterUpsertHook, dpoRuleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dpoRuleAfterUpsertHooks = []DpoRuleHook{}
}

func testDpoRulesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDpoRulesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dpoRuleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDpoRuleToOneDpoUsingDemandPartner(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DpoRule
	var foreign Dpo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dpoRuleDBTypes, false, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dpoDBTypes, false, dpoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Dpo struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DemandPartnerID = foreign.DemandPartnerID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.DemandPartner().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.DemandPartnerID != foreign.DemandPartnerID {
		t.Errorf("want: %v, got %v", foreign.DemandPartnerID, check.DemandPartnerID)
	}

	ranAfterSelectHook := false
	AddDpoHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Dpo) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := DpoRuleSlice{&local}
	if err = local.L.LoadDemandPartner(ctx, tx, false, (*[]*DpoRule)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DemandPartner == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DemandPartner = nil
	if err = local.L.LoadDemandPartner(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DemandPartner == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testDpoRuleToOnePublisherUsingDpoRulePublisher(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DpoRule
	var foreign Publisher

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, publisherDBTypes, false, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.Publisher, foreign.PublisherID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.DpoRulePublisher().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.PublisherID, foreign.PublisherID) {
		t.Errorf("want: %v, got %v", foreign.PublisherID, check.PublisherID)
	}

	ranAfterSelectHook := false
	AddPublisherHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := DpoRuleSlice{&local}
	if err = local.L.LoadDpoRulePublisher(ctx, tx, false, (*[]*DpoRule)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DpoRulePublisher == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DpoRulePublisher = nil
	if err = local.L.LoadDpoRulePublisher(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DpoRulePublisher == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testDpoRuleToOneSetOpDpoUsingDemandPartner(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DpoRule
	var b, c Dpo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dpoRuleDBTypes, false, strmangle.SetComplement(dpoRulePrimaryKeyColumns, dpoRuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dpoDBTypes, false, strmangle.SetComplement(dpoPrimaryKeyColumns, dpoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dpoDBTypes, false, strmangle.SetComplement(dpoPrimaryKeyColumns, dpoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Dpo{&b, &c} {
		err = a.SetDemandPartner(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DemandPartner != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DemandPartnerDpoRules[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DemandPartnerID != x.DemandPartnerID {
			t.Error("foreign key was wrong value", a.DemandPartnerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DemandPartnerID))
		reflect.Indirect(reflect.ValueOf(&a.DemandPartnerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DemandPartnerID != x.DemandPartnerID {
			t.Error("foreign key was wrong value", a.DemandPartnerID, x.DemandPartnerID)
		}
	}
}
func testDpoRuleToOneSetOpPublisherUsingDpoRulePublisher(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DpoRule
	var b, c Publisher

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dpoRuleDBTypes, false, strmangle.SetComplement(dpoRulePrimaryKeyColumns, dpoRuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Publisher{&b, &c} {
		err = a.SetDpoRulePublisher(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DpoRulePublisher != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DpoRules[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.Publisher, x.PublisherID) {
			t.Error("foreign key was wrong value", a.Publisher)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Publisher))
		reflect.Indirect(reflect.ValueOf(&a.Publisher)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.Publisher, x.PublisherID) {
			t.Error("foreign key was wrong value", a.Publisher, x.PublisherID)
		}
	}
}

func testDpoRuleToOneRemoveOpPublisherUsingDpoRulePublisher(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DpoRule
	var b Publisher

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dpoRuleDBTypes, false, strmangle.SetComplement(dpoRulePrimaryKeyColumns, dpoRuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDpoRulePublisher(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDpoRulePublisher(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.DpoRulePublisher().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.DpoRulePublisher != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.Publisher) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.DpoRules) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDpoRulesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
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

func testDpoRulesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DpoRuleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDpoRulesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DpoRules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dpoRuleDBTypes = map[string]string{`RuleID`: `character varying`, `DemandPartnerID`: `character varying`, `Publisher`: `character varying`, `Domain`: `character varying`, `Country`: `character varying`, `Browser`: `character varying`, `Os`: `character varying`, `DeviceType`: `character varying`, `PlacementType`: `character varying`, `Factor`: `double precision`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `Active`: `boolean`}
	_              = bytes.MinRead
)

func testDpoRulesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dpoRulePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dpoRuleAllColumns) == len(dpoRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDpoRulesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dpoRuleAllColumns) == len(dpoRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DpoRule{}
	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dpoRuleDBTypes, true, dpoRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dpoRuleAllColumns, dpoRulePrimaryKeyColumns) {
		fields = dpoRuleAllColumns
	} else {
		fields = strmangle.SetComplement(
			dpoRuleAllColumns,
			dpoRulePrimaryKeyColumns,
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

	slice := DpoRuleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDpoRulesUpsert(t *testing.T) {
	t.Parallel()

	if len(dpoRuleAllColumns) == len(dpoRulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DpoRule{}
	if err = randomize.Struct(seed, &o, dpoRuleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DpoRule: %s", err)
	}

	count, err := DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dpoRuleDBTypes, false, dpoRulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DpoRule struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DpoRule: %s", err)
	}

	count, err = DpoRules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
