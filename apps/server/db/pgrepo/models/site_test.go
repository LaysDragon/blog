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

func testSites(t *testing.T) {
	t.Parallel()

	query := Sites()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSitesSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSitesQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Sites().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSitesSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SiteSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSitesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSitesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Sites().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSitesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SiteSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSitesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SiteExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Site exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SiteExists to return true, but got false.")
	}
}

func testSitesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	siteFound, err := FindSite(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if siteFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSitesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Sites().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSitesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Sites().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSitesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	siteOne := &Site{}
	siteTwo := &Site{}
	if err = randomize.Struct(seed, siteOne, siteDBTypes, false, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}
	if err = randomize.Struct(seed, siteTwo, siteDBTypes, false, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = siteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = siteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sites().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSitesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	siteOne := &Site{}
	siteTwo := &Site{}
	if err = randomize.Struct(seed, siteOne, siteDBTypes, false, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}
	if err = randomize.Struct(seed, siteTwo, siteDBTypes, false, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = siteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = siteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func siteBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func siteAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Site) error {
	*o = Site{}
	return nil
}

func testSitesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Site{}
	o := &Site{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, siteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Site object: %s", err)
	}

	AddSiteHook(boil.BeforeInsertHook, siteBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	siteBeforeInsertHooks = []SiteHook{}

	AddSiteHook(boil.AfterInsertHook, siteAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	siteAfterInsertHooks = []SiteHook{}

	AddSiteHook(boil.AfterSelectHook, siteAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	siteAfterSelectHooks = []SiteHook{}

	AddSiteHook(boil.BeforeUpdateHook, siteBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	siteBeforeUpdateHooks = []SiteHook{}

	AddSiteHook(boil.AfterUpdateHook, siteAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	siteAfterUpdateHooks = []SiteHook{}

	AddSiteHook(boil.BeforeDeleteHook, siteBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	siteBeforeDeleteHooks = []SiteHook{}

	AddSiteHook(boil.AfterDeleteHook, siteAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	siteAfterDeleteHooks = []SiteHook{}

	AddSiteHook(boil.BeforeUpsertHook, siteBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	siteBeforeUpsertHooks = []SiteHook{}

	AddSiteHook(boil.AfterUpsertHook, siteAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	siteAfterUpsertHooks = []SiteHook{}
}

func testSitesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSitesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(strmangle.SetMerge(sitePrimaryKeyColumns, siteColumnsWithoutDefault)...)); err != nil {
		t.Error(err)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSiteToManyAttachtments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Site
	var b, c Attachtment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
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

	b.SiteID = a.ID
	c.SiteID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Attachtments().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SiteID == b.SiteID {
			bFound = true
		}
		if v.SiteID == c.SiteID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SiteSlice{&a}
	if err = a.L.LoadAttachtments(ctx, tx, false, (*[]*Site)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Attachtments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Attachtments = nil
	if err = a.L.LoadAttachtments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Attachtments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSiteToManyPosts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Site
	var b, c Post

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
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

	b.SiteID = a.ID
	c.SiteID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Posts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SiteID == b.SiteID {
			bFound = true
		}
		if v.SiteID == c.SiteID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SiteSlice{&a}
	if err = a.L.LoadPosts(ctx, tx, false, (*[]*Site)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Posts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Posts = nil
	if err = a.L.LoadPosts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Posts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSiteToManySiteRoles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Site
	var b, c SiteRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, siteRoleDBTypes, false, siteRoleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, siteRoleDBTypes, false, siteRoleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SiteID = a.ID
	c.SiteID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SiteRoles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SiteID == b.SiteID {
			bFound = true
		}
		if v.SiteID == c.SiteID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SiteSlice{&a}
	if err = a.L.LoadSiteRoles(ctx, tx, false, (*[]*Site)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SiteRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SiteRoles = nil
	if err = a.L.LoadSiteRoles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SiteRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSiteToManyAddOpAttachtments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Site
	var b, c, d, e Attachtment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, siteDBTypes, false, strmangle.SetComplement(sitePrimaryKeyColumns, siteColumnsWithoutDefault)...); err != nil {
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
		err = a.AddAttachtments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SiteID {
			t.Error("foreign key was wrong value", a.ID, first.SiteID)
		}
		if a.ID != second.SiteID {
			t.Error("foreign key was wrong value", a.ID, second.SiteID)
		}

		if first.R.Site != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Site != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Attachtments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Attachtments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Attachtments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSiteToManyAddOpPosts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Site
	var b, c, d, e Post

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, siteDBTypes, false, strmangle.SetComplement(sitePrimaryKeyColumns, siteColumnsWithoutDefault)...); err != nil {
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
		err = a.AddPosts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SiteID {
			t.Error("foreign key was wrong value", a.ID, first.SiteID)
		}
		if a.ID != second.SiteID {
			t.Error("foreign key was wrong value", a.ID, second.SiteID)
		}

		if first.R.Site != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Site != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Posts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Posts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Posts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSiteToManyAddOpSiteRoles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Site
	var b, c, d, e SiteRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, siteDBTypes, false, strmangle.SetComplement(sitePrimaryKeyColumns, siteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*SiteRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, siteRoleDBTypes, false, strmangle.SetComplement(siteRolePrimaryKeyColumns, siteRoleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*SiteRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSiteRoles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SiteID {
			t.Error("foreign key was wrong value", a.ID, first.SiteID)
		}
		if a.ID != second.SiteID {
			t.Error("foreign key was wrong value", a.ID, second.SiteID)
		}

		if first.R.Site != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Site != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SiteRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SiteRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SiteRoles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testSitesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
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

func testSitesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SiteSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSitesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sites().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	siteDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `Name`: `character varying`}
	_           = bytes.MinRead
)

func testSitesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(sitePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(siteAllColumns) == len(sitePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, siteDBTypes, true, sitePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSitesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(siteAllColumns) == len(sitePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Site{}
	if err = randomize.Struct(seed, o, siteDBTypes, true, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, siteDBTypes, true, sitePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(siteAllColumns, sitePrimaryKeyColumns) {
		fields = siteAllColumns
	} else {
		fields = strmangle.SetComplement(
			siteAllColumns,
			sitePrimaryKeyColumns,
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

	slice := SiteSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSitesUpsert(t *testing.T) {
	t.Parallel()

	if len(siteAllColumns) == len(sitePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Site{}
	if err = randomize.Struct(seed, &o, siteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Site: %s", err)
	}

	count, err := Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, siteDBTypes, false, sitePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Site: %s", err)
	}

	count, err = Sites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
