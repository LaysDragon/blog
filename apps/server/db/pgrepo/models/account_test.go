// Code generated by SQLBoiler 4.19.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testAccounts(t *testing.T) {
	t.Parallel()

	query := Accounts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAccountsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
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

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Accounts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AccountExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Account exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AccountExists to return true, but got false.")
	}
}

func testAccountsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	accountFound, err := FindAccount(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if accountFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAccountsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Accounts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAccountsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Accounts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAccountsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAccountsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func accountBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func testAccountsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Account{}
	o := &Account{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, accountDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Account object: %s", err)
	}

	AddAccountHook(boil.BeforeInsertHook, accountBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	accountBeforeInsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterInsertHook, accountAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	accountAfterInsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterSelectHook, accountAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	accountAfterSelectHooks = []AccountHook{}

	AddAccountHook(boil.BeforeUpdateHook, accountBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	accountBeforeUpdateHooks = []AccountHook{}

	AddAccountHook(boil.AfterUpdateHook, accountAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	accountAfterUpdateHooks = []AccountHook{}

	AddAccountHook(boil.BeforeDeleteHook, accountBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	accountBeforeDeleteHooks = []AccountHook{}

	AddAccountHook(boil.AfterDeleteHook, accountAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	accountAfterDeleteHooks = []AccountHook{}

	AddAccountHook(boil.BeforeUpsertHook, accountBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	accountBeforeUpsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterUpsertHook, accountAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	accountAfterUpsertHooks = []AccountHook{}
}

func testAccountsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(strmangle.SetMerge(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...)); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountToManyUserAccessLogs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c AccessLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, accessLogDBTypes, false, accessLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, accessLogDBTypes, false, accessLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.UserID, a.ID)
	queries.Assign(&c.UserID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserAccessLogs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.UserID, b.UserID) {
			bFound = true
		}
		if queries.Equal(v.UserID, c.UserID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadUserAccessLogs(ctx, tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAccessLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserAccessLogs = nil
	if err = a.L.LoadUserAccessLogs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAccessLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAccountToManyOwnerAttachtments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c Attachtment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, attachtmentDBTypes, false, attachtmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, attachtmentDBTypes, false, attachtmentColumnsWithDefault...); err != nil {
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

	check, err := a.OwnerAttachtments().All(ctx, tx)
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

	slice := AccountSlice{&a}
	if err = a.L.LoadOwnerAttachtments(ctx, tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OwnerAttachtments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OwnerAttachtments = nil
	if err = a.L.LoadOwnerAttachtments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OwnerAttachtments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAccountToManyOwnerPosts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c Post

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, postDBTypes, false, postColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, postDBTypes, false, postColumnsWithDefault...); err != nil {
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

	check, err := a.OwnerPosts().All(ctx, tx)
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

	slice := AccountSlice{&a}
	if err = a.L.LoadOwnerPosts(ctx, tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OwnerPosts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OwnerPosts = nil
	if err = a.L.LoadOwnerPosts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OwnerPosts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAccountToManyAddOpUserAccessLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e AccessLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AccessLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, accessLogDBTypes, false, strmangle.SetComplement(accessLogPrimaryKeyColumns, accessLogColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AccessLog{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserAccessLogs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.UserID) {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if !queries.Equal(a.ID, second.UserID) {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserAccessLogs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserAccessLogs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserAccessLogs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAccountToManySetOpUserAccessLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e AccessLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AccessLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, accessLogDBTypes, false, strmangle.SetComplement(accessLogPrimaryKeyColumns, accessLogColumnsWithoutDefault)...); err != nil {
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

	err = a.SetUserAccessLogs(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserAccessLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUserAccessLogs(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserAccessLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.UserID) {
		t.Error("foreign key was wrong value", a.ID, d.UserID)
	}
	if !queries.Equal(a.ID, e.UserID) {
		t.Error("foreign key was wrong value", a.ID, e.UserID)
	}

	if b.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.User != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.User != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.UserAccessLogs[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.UserAccessLogs[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testAccountToManyRemoveOpUserAccessLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e AccessLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AccessLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, accessLogDBTypes, false, strmangle.SetComplement(accessLogPrimaryKeyColumns, accessLogColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUserAccessLogs(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserAccessLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUserAccessLogs(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserAccessLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.User != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.User != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.User != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.UserAccessLogs) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.UserAccessLogs[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.UserAccessLogs[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testAccountToManyAddOpOwnerAttachtments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e Attachtment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Attachtment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, attachtmentDBTypes, false, strmangle.SetComplement(attachtmentPrimaryKeyColumns, attachtmentColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Attachtment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOwnerAttachtments(ctx, tx, i != 0, x...)
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

		if a.R.OwnerAttachtments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OwnerAttachtments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OwnerAttachtments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAccountToManyAddOpOwnerPosts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e Post

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Post{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, postDBTypes, false, strmangle.SetComplement(postPrimaryKeyColumns, postColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Post{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOwnerPosts(ctx, tx, i != 0, x...)
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

		if a.R.OwnerPosts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OwnerPosts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OwnerPosts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAccountsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
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

func testAccountsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAccountsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	accountDBTypes = map[string]string{`ID`: `integer`, `CreatedDate`: `timestamp with time zone`, `UpdatedDate`: `timestamp with time zone`, `Username`: `character varying`, `Role`: `character varying`, `Email`: `character varying`, `PassHash`: `character varying`}
	_              = bytes.MinRead
)

func testAccountsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(accountAllColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAccountsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(accountAllColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(accountAllColumns, accountPrimaryKeyColumns) {
		fields = accountAllColumns
	} else {
		fields = strmangle.SetComplement(
			accountAllColumns,
			accountPrimaryKeyColumns,
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

	slice := AccountSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAccountsUpsert(t *testing.T) {
	t.Parallel()

	if len(accountAllColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Account{}
	if err = randomize.Struct(seed, &o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, accountDBTypes, false, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err = Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
