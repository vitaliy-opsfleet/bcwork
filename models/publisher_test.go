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

func testPublishers(t *testing.T) {
	t.Parallel()

	query := Publishers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPublishersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
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

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublishersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Publishers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublishersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PublisherSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPublishersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PublisherExists(ctx, tx, o.PublisherID)
	if err != nil {
		t.Errorf("Unable to check if Publisher exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PublisherExists to return true, but got false.")
	}
}

func testPublishersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	publisherFound, err := FindPublisher(ctx, tx, o.PublisherID)
	if err != nil {
		t.Error(err)
	}

	if publisherFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPublishersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Publishers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPublishersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Publishers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPublishersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	publisherOne := &Publisher{}
	publisherTwo := &Publisher{}
	if err = randomize.Struct(seed, publisherOne, publisherDBTypes, false, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}
	if err = randomize.Struct(seed, publisherTwo, publisherDBTypes, false, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = publisherOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = publisherTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Publishers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPublishersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	publisherOne := &Publisher{}
	publisherTwo := &Publisher{}
	if err = randomize.Struct(seed, publisherOne, publisherDBTypes, false, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}
	if err = randomize.Struct(seed, publisherTwo, publisherDBTypes, false, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = publisherOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = publisherTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func publisherBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func publisherAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
	*o = Publisher{}
	return nil
}

func testPublishersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Publisher{}
	o := &Publisher{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, publisherDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Publisher object: %s", err)
	}

	AddPublisherHook(boil.BeforeInsertHook, publisherBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	publisherBeforeInsertHooks = []PublisherHook{}

	AddPublisherHook(boil.AfterInsertHook, publisherAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	publisherAfterInsertHooks = []PublisherHook{}

	AddPublisherHook(boil.AfterSelectHook, publisherAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	publisherAfterSelectHooks = []PublisherHook{}

	AddPublisherHook(boil.BeforeUpdateHook, publisherBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	publisherBeforeUpdateHooks = []PublisherHook{}

	AddPublisherHook(boil.AfterUpdateHook, publisherAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	publisherAfterUpdateHooks = []PublisherHook{}

	AddPublisherHook(boil.BeforeDeleteHook, publisherBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	publisherBeforeDeleteHooks = []PublisherHook{}

	AddPublisherHook(boil.AfterDeleteHook, publisherAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	publisherAfterDeleteHooks = []PublisherHook{}

	AddPublisherHook(boil.BeforeUpsertHook, publisherBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	publisherBeforeUpsertHooks = []PublisherHook{}

	AddPublisherHook(boil.AfterUpsertHook, publisherAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	publisherAfterUpsertHooks = []PublisherHook{}
}

func testPublishersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPublishersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(publisherColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPublisherToManyConfiants(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Publisher
	var b, c Confiant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, confiantDBTypes, false, confiantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, confiantDBTypes, false, confiantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PublisherID = a.PublisherID
	c.PublisherID = a.PublisherID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Confiants().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PublisherID == b.PublisherID {
			bFound = true
		}
		if v.PublisherID == c.PublisherID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PublisherSlice{&a}
	if err = a.L.LoadConfiants(ctx, tx, false, (*[]*Publisher)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Confiants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Confiants = nil
	if err = a.L.LoadConfiants(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Confiants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPublisherToManyPublisherDomains(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Publisher
	var b, c PublisherDomain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, publisherDomainDBTypes, false, publisherDomainColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, publisherDomainDBTypes, false, publisherDomainColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PublisherID = a.PublisherID
	c.PublisherID = a.PublisherID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PublisherDomains().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PublisherID == b.PublisherID {
			bFound = true
		}
		if v.PublisherID == c.PublisherID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PublisherSlice{&a}
	if err = a.L.LoadPublisherDomains(ctx, tx, false, (*[]*Publisher)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PublisherDomains); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PublisherDomains = nil
	if err = a.L.LoadPublisherDomains(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PublisherDomains); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPublisherToManyAddOpConfiants(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Publisher
	var b, c, d, e Confiant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Confiant{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, confiantDBTypes, false, strmangle.SetComplement(confiantPrimaryKeyColumns, confiantColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Confiant{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddConfiants(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.PublisherID != first.PublisherID {
			t.Error("foreign key was wrong value", a.PublisherID, first.PublisherID)
		}
		if a.PublisherID != second.PublisherID {
			t.Error("foreign key was wrong value", a.PublisherID, second.PublisherID)
		}

		if first.R.Publisher != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Publisher != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Confiants[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Confiants[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Confiants().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPublisherToManyAddOpPublisherDomains(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Publisher
	var b, c, d, e PublisherDomain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PublisherDomain{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, publisherDomainDBTypes, false, strmangle.SetComplement(publisherDomainPrimaryKeyColumns, publisherDomainColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PublisherDomain{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPublisherDomains(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.PublisherID != first.PublisherID {
			t.Error("foreign key was wrong value", a.PublisherID, first.PublisherID)
		}
		if a.PublisherID != second.PublisherID {
			t.Error("foreign key was wrong value", a.PublisherID, second.PublisherID)
		}

		if first.R.Publisher != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Publisher != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PublisherDomains[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PublisherDomains[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PublisherDomains().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPublishersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
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

func testPublishersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PublisherSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPublishersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Publishers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	publisherDBTypes = map[string]string{`PublisherID`: `character varying`, `CreatedAt`: `timestamp without time zone`, `Name`: `character varying`, `AccountManagerID`: `character varying`, `MediaBuyerID`: `character varying`, `CampaignManagerID`: `character varying`, `OfficeLocation`: `character varying`, `PauseTimestamp`: `bigint`, `StartTimestamp`: `bigint`, `ReactivateTimestamp`: `bigint`, `Status`: `character varying`}
	_                = bytes.MinRead
)

func testPublishersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(publisherPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(publisherAllColumns) == len(publisherPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPublishersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(publisherAllColumns) == len(publisherPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Publisher{}
	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, publisherDBTypes, true, publisherPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(publisherAllColumns, publisherPrimaryKeyColumns) {
		fields = publisherAllColumns
	} else {
		fields = strmangle.SetComplement(
			publisherAllColumns,
			publisherPrimaryKeyColumns,
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

	slice := PublisherSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPublishersUpsert(t *testing.T) {
	t.Parallel()

	if len(publisherAllColumns) == len(publisherPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Publisher{}
	if err = randomize.Struct(seed, &o, publisherDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Publisher: %s", err)
	}

	count, err := Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, publisherDBTypes, false, publisherPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Publisher: %s", err)
	}

	count, err = Publishers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
