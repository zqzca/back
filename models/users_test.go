package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testUsers(t *testing.T) {
	t.Parallel()

	query := Users(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = user.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Users(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserSlice{user}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UserExists(tx, user.ID)
	if err != nil {
		t.Errorf("Unable to check if User exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserExistsG to return true, but got false.")
	}
}
func testUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	userFound, err := FindUser(tx, user.ID)
	if err != nil {
		t.Error(err)
	}

	if userFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Users(tx).Bind(user); err != nil {
		t.Error(err)
	}
}

func testUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Users(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userOne := &User{}
	userTwo := &User{}
	if err = randomize.Struct(seed, userOne, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}
	if err = randomize.Struct(seed, userTwo, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Users(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userOne := &User{}
	userTwo := &User{}
	if err = randomize.Struct(seed, userOne, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}
	if err = randomize.Struct(seed, userTwo, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = userOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = userTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func userBeforeInsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterInsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterSelectHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeUpdateHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterUpdateHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeDeleteHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterDeleteHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userBeforeUpsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func userAfterUpsertHook(e boil.Executor, o *User) error {
	*o = User{}
	return nil
}

func testUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &User{}
	o := &User{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userDBTypes, false); err != nil {
		t.Errorf("Unable to randomize User object: %s", err)
	}

	AddUserHook(boil.BeforeInsertHook, userBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userBeforeInsertHooks = []UserHook{}

	AddUserHook(boil.AfterInsertHook, userAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userAfterInsertHooks = []UserHook{}

	AddUserHook(boil.AfterSelectHook, userAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userAfterSelectHooks = []UserHook{}

	AddUserHook(boil.BeforeUpdateHook, userBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userBeforeUpdateHooks = []UserHook{}

	AddUserHook(boil.AfterUpdateHook, userAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userAfterUpdateHooks = []UserHook{}

	AddUserHook(boil.BeforeDeleteHook, userBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userBeforeDeleteHooks = []UserHook{}

	AddUserHook(boil.AfterDeleteHook, userAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userAfterDeleteHooks = []UserHook{}

	AddUserHook(boil.BeforeUpsertHook, userBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userBeforeUpsertHooks = []UserHook{}

	AddUserHook(boil.AfterUpsertHook, userAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userAfterUpsertHooks = []UserHook{}
}
func testUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx, userColumns...); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = user.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UserSlice{user}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Users(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userDBTypes = map[string]string{"Banned": "boolean", "CreatedAt": "timestamp without time zone", "Email": "character varying", "FirstName": "character varying", "Hash": "character varying", "ID": "uuid", "LastName": "character varying", "Phone": "character varying", "UpdatedAt": "timestamp without time zone", "Username": "character varying"}
	_           = bytes.MinRead
)

func testUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(userColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, user, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err = user.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	user := &User{}
	if err = randomize.Struct(seed, user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, user, userDBTypes, true, userPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userColumns, userPrimaryKeyColumns) {
		fields = userColumns
	} else {
		fields = strmangle.SetComplement(
			userColumns,
			userPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(user))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UserSlice{user}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(userColumns) == len(userPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	user := User{}
	if err = randomize.Struct(seed, &user, userDBTypes, true); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = user.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert User: %s", err)
	}

	count, err := Users(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &user, userDBTypes, false, userPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err = user.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert User: %s", err)
	}

	count, err = Users(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
