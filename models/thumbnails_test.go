package models

import (
	"testing"
	"reflect"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testThumbnails(t *testing.T) {
	t.Parallel()

	query := Thumbnails(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testThumbnailsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = thumbnail.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testThumbnailsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Thumbnails(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testThumbnailsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ThumbnailSlice{thumbnail}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testThumbnailsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ThumbnailExists(tx, thumbnail.ID)
	if err != nil {
		t.Errorf("Unable to check if Thumbnail exists: %s", err)
	}
	if e != true {
		t.Errorf("Expected ThumbnailExistsG to return true, but got false.")
	}
}

func testThumbnailsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	thumbnailFound, err := FindThumbnail(tx, thumbnail.ID)
	if err != nil {
		t.Error(err)
	}

	if thumbnailFound == nil {
		t.Error("want a record, got nil")
	}
}

func testThumbnailsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Thumbnails(tx).Bind(thumbnail); err != nil {
		t.Error(err)
	}
}

func testThumbnailsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Thumbnails(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testThumbnailsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnailOne := &Thumbnail{}
	thumbnailTwo := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnailOne, thumbnailDBTypes, false, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}
	if err = randomize.Struct(seed, thumbnailTwo, thumbnailDBTypes, false, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnailOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = thumbnailTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Thumbnails(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testThumbnailsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	thumbnailOne := &Thumbnail{}
	thumbnailTwo := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnailOne, thumbnailDBTypes, false, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}
	if err = randomize.Struct(seed, thumbnailTwo, thumbnailDBTypes, false, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnailOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = thumbnailTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

var thumbnailDBTypes = map[string]string{"ID": "uuid", "FileID": "uuid", "Size": "integer", "Hash": "text", "CreatedAt": "timestamp without time zone", "UpdatedAt": "timestamp without time zone"}

func testThumbnailsInPrimaryKeyArgs(t *testing.T) {
	t.Parallel()

	var err error
	var o Thumbnail
	o = Thumbnail{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &o, thumbnailDBTypes, true); err != nil {
		t.Errorf("Could not randomize struct: %s", err)
	}

	args := o.inPrimaryKeyArgs()

	if len(args) != len(thumbnailPrimaryKeyColumns) {
		t.Errorf("Expected args to be len %d, but got %d", len(thumbnailPrimaryKeyColumns), len(args))
	}

	if o.ID != args[0] {
		t.Errorf("Expected args[0] to be value of o.ID, but got %#v", args[0])
	}
}

func testThumbnailsSliceInPrimaryKeyArgs(t *testing.T) {
	t.Parallel()

	var err error
	o := make(ThumbnailSlice, 3)

	seed := randomize.NewSeed()
	for i := range o {
		o[i] = &Thumbnail{}
		if err = randomize.Struct(seed, o[i], thumbnailDBTypes, true); err != nil {
			t.Errorf("Could not randomize struct: %s", err)
		}
	}

	args := o.inPrimaryKeyArgs()

	if len(args) != len(thumbnailPrimaryKeyColumns)*3 {
		t.Errorf("Expected args to be len %d, but got %d", len(thumbnailPrimaryKeyColumns)*3, len(args))
	}

	argC := 0
	for i := 0; i < 3; i++ {

		if o[i].ID != args[argC] {
			t.Errorf("Expected args[%d] to be value of o.ID, but got %#v", i, args[i])
		}
		argC++
	}
}

func thumbnailBeforeInsertHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailAfterInsertHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailAfterSelectHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailBeforeUpdateHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailAfterUpdateHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailBeforeDeleteHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailAfterDeleteHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailBeforeUpsertHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func thumbnailAfterUpsertHook(e boil.Executor, o *Thumbnail) error {
	*o = Thumbnail{}
	return nil
}

func testThumbnailsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Thumbnail{}
	o := &Thumbnail{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, thumbnailDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Thumbnail object: %s", err)
	}

	AddThumbnailHook(boil.BeforeInsertHook, thumbnailBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	thumbnailBeforeInsertHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.AfterInsertHook, thumbnailAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	thumbnailAfterInsertHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.AfterSelectHook, thumbnailAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	thumbnailAfterSelectHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.BeforeUpdateHook, thumbnailBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	thumbnailBeforeUpdateHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.AfterUpdateHook, thumbnailAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	thumbnailAfterUpdateHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.BeforeDeleteHook, thumbnailBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	thumbnailBeforeDeleteHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.AfterDeleteHook, thumbnailAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	thumbnailAfterDeleteHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.BeforeUpsertHook, thumbnailBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	thumbnailBeforeUpsertHooks = []ThumbnailHook{}

	AddThumbnailHook(boil.AfterUpsertHook, thumbnailAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	thumbnailAfterUpsertHooks = []ThumbnailHook{}
}

func testThumbnailsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testThumbnailsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx, thumbnailColumns...); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}




func testThumbnailToOneFile_File(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var foreign File
	var local Thumbnail

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.FileID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}
	check, err := local.File(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ThumbnailSlice{&local}
	if err = local.L.LoadFile(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.File == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.File = nil
	if err = local.L.LoadFile(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.File == nil {
		t.Error("struct should have been eager loaded")
	}
}




func testThumbnailToOneSetOpFile_File(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Thumbnail
	var b, c File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, thumbnailDBTypes, false, thumbnailPrimaryKeyColumns...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, fileDBTypes, false, filePrimaryKeyColumns...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, fileDBTypes, false, filePrimaryKeyColumns...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*File{&b, &c} {
		err = a.SetFile(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.FileID != x.ID {
			t.Error("foreign key was wrong value", a.FileID)
		}
		if a.R.File != x {
			t.Error("relationship struct not set to correct value")
		}

		zero := reflect.Zero(reflect.TypeOf(a.FileID))
		reflect.Indirect(reflect.ValueOf(&a.FileID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FileID != x.ID {
			t.Error("foreign key was wrong value", a.FileID, x.ID)
		}

		if x.R.Thumbnails[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
	}
}
func testThumbnailsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = thumbnail.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testThumbnailsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ThumbnailSlice{thumbnail}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testThumbnailsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Thumbnails(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

func testThumbnailsUpdate(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	// If table only contains primary key columns, we need to pass
	// them into a whitelist to get a valid test result,
	// otherwise the Update method will error because it will not be able to
	// generate a whitelist (due to it excluding primary key columns).
	if strmangle.StringSliceMatch(thumbnailColumns, thumbnailPrimaryKeyColumns) {
		if err = thumbnail.Update(tx, thumbnailPrimaryKeyColumns...); err != nil {
			t.Error(err)
		}
	} else {
		if err = thumbnail.Update(tx); err != nil {
			t.Error(err)
		}
	}
}

func testThumbnailsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thumbnail := &Thumbnail{}
	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, thumbnail, thumbnailDBTypes, true, thumbnailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(thumbnailColumns, thumbnailPrimaryKeyColumns) {
		fields = thumbnailColumns
	} else {
		fields = strmangle.SetComplement(
			thumbnailColumns,
			thumbnailPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(thumbnail))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ThumbnailSlice{thumbnail}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testThumbnailsUpsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	thumbnail := Thumbnail{}
	if err = randomize.Struct(seed, &thumbnail, thumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = thumbnail.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Thumbnail: %s", err)
	}

	count, err := Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &thumbnail, thumbnailDBTypes, false, thumbnailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Thumbnail struct: %s", err)
	}

	if err = thumbnail.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Thumbnail: %s", err)
	}

	count, err = Thumbnails(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

