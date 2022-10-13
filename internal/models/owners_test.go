// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testOwnersUpsert(t *testing.T) {
	t.Parallel()

	if len(ownerAllColumns) == len(ownerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Owner{}
	if err = randomize.Struct(seed, &o, ownerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Owner: %s", err)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, ownerDBTypes, false, ownerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Owner: %s", err)
	}

	count, err = Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOwners(t *testing.T) {
	t.Parallel()

	query := Owners()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOwnersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
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

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOwnersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Owners().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOwnersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OwnerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOwnersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OwnerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Owner exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OwnerExists to return true, but got false.")
	}
}

func testOwnersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	ownerFound, err := FindOwner(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if ownerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOwnersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Owners().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOwnersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Owners().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOwnersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ownerOne := &Owner{}
	ownerTwo := &Owner{}
	if err = randomize.Struct(seed, ownerOne, ownerDBTypes, false, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}
	if err = randomize.Struct(seed, ownerTwo, ownerDBTypes, false, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ownerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ownerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Owners().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOwnersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ownerOne := &Owner{}
	ownerTwo := &Owner{}
	if err = randomize.Struct(seed, ownerOne, ownerDBTypes, false, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}
	if err = randomize.Struct(seed, ownerTwo, ownerDBTypes, false, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ownerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ownerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func ownerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func ownerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Owner) error {
	*o = Owner{}
	return nil
}

func testOwnersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Owner{}
	o := &Owner{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ownerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Owner object: %s", err)
	}

	AddOwnerHook(boil.BeforeInsertHook, ownerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ownerBeforeInsertHooks = []OwnerHook{}

	AddOwnerHook(boil.AfterInsertHook, ownerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ownerAfterInsertHooks = []OwnerHook{}

	AddOwnerHook(boil.AfterSelectHook, ownerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ownerAfterSelectHooks = []OwnerHook{}

	AddOwnerHook(boil.BeforeUpdateHook, ownerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ownerBeforeUpdateHooks = []OwnerHook{}

	AddOwnerHook(boil.AfterUpdateHook, ownerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ownerAfterUpdateHooks = []OwnerHook{}

	AddOwnerHook(boil.BeforeDeleteHook, ownerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ownerBeforeDeleteHooks = []OwnerHook{}

	AddOwnerHook(boil.AfterDeleteHook, ownerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ownerAfterDeleteHooks = []OwnerHook{}

	AddOwnerHook(boil.BeforeUpsertHook, ownerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ownerBeforeUpsertHooks = []OwnerHook{}

	AddOwnerHook(boil.AfterUpsertHook, ownerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ownerAfterUpsertHooks = []OwnerHook{}
}

func testOwnersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOwnersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(ownerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOwnerToManyAnswers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Owner
	var b, c Answer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, answerDBTypes, false, answerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, answerDBTypes, false, answerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.OwnerID = a.ID
	c.OwnerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Answers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.OwnerID == b.OwnerID {
			bFound = true
		}
		if v.OwnerID == c.OwnerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OwnerSlice{&a}
	if err = a.L.LoadAnswers(ctx, tx, false, (*[]*Owner)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Answers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Answers = nil
	if err = a.L.LoadAnswers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Answers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOwnerToManyAddOpAnswers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Owner
	var b, c, d, e Answer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ownerDBTypes, false, strmangle.SetComplement(ownerPrimaryKeyColumns, ownerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Answer{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, answerDBTypes, false, strmangle.SetComplement(answerPrimaryKeyColumns, answerColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Answer{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAnswers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.OwnerID {
			t.Error("foreign key was wrong value", a.ID, first.OwnerID)
		}
		if a.ID != second.OwnerID {
			t.Error("foreign key was wrong value", a.ID, second.OwnerID)
		}

		if first.R.Owner != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Owner != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Answers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Answers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Answers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOwnersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
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

func testOwnersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OwnerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOwnersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Owners().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ownerDBTypes = map[string]string{`ID`: `uuid`, `Name`: `string`, `Origin`: `string`, `Service`: `string`, `CreatedAt`: `timestamptz`, `UpdatedAt`: `timestamptz`}
	_            = bytes.MinRead
)

func testOwnersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(ownerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(ownerAllColumns) == len(ownerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOwnersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ownerAllColumns) == len(ownerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Owner{}
	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Owners().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ownerDBTypes, true, ownerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Owner struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ownerAllColumns, ownerPrimaryKeyColumns) {
		fields = ownerAllColumns
	} else {
		fields = strmangle.SetComplement(
			ownerAllColumns,
			ownerPrimaryKeyColumns,
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

	slice := OwnerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
