package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/boil/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// Thumbnail is an object representing the database table.
type Thumbnail struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	FileID    string    `boil:"file_id" json:"file_id" toml:"file_id" yaml:"file_id"`
	Size      int       `boil:"size" json:"size" toml:"size" yaml:"size"`
	Hash      string    `boil:"hash" json:"hash" toml:"hash" yaml:"hash"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *ThumbnailR `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// ThumbnailR is where relationships are stored.
type ThumbnailR struct {
	File *File
}

var (
	thumbnailColumns               = []string{"id", "file_id", "size", "hash", "created_at", "updated_at"}
	thumbnailColumnsWithoutDefault = []string{"file_id", "size", "hash", "created_at", "updated_at"}
	thumbnailColumnsWithDefault    = []string{"id"}
	thumbnailPrimaryKeyColumns     = []string{"id"}
	thumbnailTitleCases            = map[string]string{
		"id":         "ID",
		"file_id":    "FileID",
		"size":       "Size",
		"hash":       "Hash",
		"created_at": "CreatedAt",
		"updated_at": "UpdatedAt",
	}
)

type (
	ThumbnailSlice []*Thumbnail
	ThumbnailHook  func(boil.Executor, *Thumbnail) error

	thumbnailQuery struct {
		*boil.Query
	}
)

// Force time package dependency for automated UpdatedAt/CreatedAt.
var _ = time.Second

var thumbnailBeforeInsertHooks []ThumbnailHook
var thumbnailBeforeUpdateHooks []ThumbnailHook
var thumbnailBeforeDeleteHooks []ThumbnailHook
var thumbnailBeforeUpsertHooks []ThumbnailHook

var thumbnailAfterInsertHooks []ThumbnailHook
var thumbnailAfterSelectHooks []ThumbnailHook
var thumbnailAfterUpdateHooks []ThumbnailHook
var thumbnailAfterDeleteHooks []ThumbnailHook
var thumbnailAfterUpsertHooks []ThumbnailHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Thumbnail) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Thumbnail) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Thumbnail) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Thumbnail) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Thumbnail) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Thumbnail) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Thumbnail) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Thumbnail) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Thumbnail) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range thumbnailAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

func ThumbnailAddHook(hookPoint boil.HookPoint, thumbnailHook ThumbnailHook) {
	switch hookPoint {
	case boil.HookBeforeInsert:
		thumbnailBeforeInsertHooks = append(thumbnailBeforeInsertHooks, thumbnailHook)
	case boil.HookBeforeUpdate:
		thumbnailBeforeUpdateHooks = append(thumbnailBeforeUpdateHooks, thumbnailHook)
	case boil.HookBeforeDelete:
		thumbnailBeforeDeleteHooks = append(thumbnailBeforeDeleteHooks, thumbnailHook)
	case boil.HookBeforeUpsert:
		thumbnailBeforeUpsertHooks = append(thumbnailBeforeUpsertHooks, thumbnailHook)
	case boil.HookAfterInsert:
		thumbnailAfterInsertHooks = append(thumbnailAfterInsertHooks, thumbnailHook)
	case boil.HookAfterSelect:
		thumbnailAfterSelectHooks = append(thumbnailAfterSelectHooks, thumbnailHook)
	case boil.HookAfterUpdate:
		thumbnailAfterUpdateHooks = append(thumbnailAfterUpdateHooks, thumbnailHook)
	case boil.HookAfterDelete:
		thumbnailAfterDeleteHooks = append(thumbnailAfterDeleteHooks, thumbnailHook)
	case boil.HookAfterUpsert:
		thumbnailAfterUpsertHooks = append(thumbnailAfterUpsertHooks, thumbnailHook)
	}
}

// OneP returns a single thumbnail record from the query, and panics on error.
func (q thumbnailQuery) OneP() *Thumbnail {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single thumbnail record from the query.
func (q thumbnailQuery) One() (*Thumbnail, error) {
	o := &Thumbnail{}

	boil.SetLimit(q.Query, 1)

	err := q.BindFast(o, thumbnailTitleCases)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for thumbnails")
	}

	if err := o.doAfterSelectHooks(boil.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Thumbnail records from the query, and panics on error.
func (q thumbnailQuery) AllP() ThumbnailSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Thumbnail records from the query.
func (q thumbnailQuery) All() (ThumbnailSlice, error) {
	var o ThumbnailSlice

	err := q.BindFast(&o, thumbnailTitleCases)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Thumbnail slice")
	}

	if len(thumbnailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(boil.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Thumbnail records in the query, and panics on error.
func (q thumbnailQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Thumbnail records in the query.
func (q thumbnailQuery) Count() (int64, error) {
	var count int64

	boil.SetCount(q.Query)

	err := boil.ExecQueryOne(q.Query).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count thumbnails rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q thumbnailQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q thumbnailQuery) Exists() (bool, error) {
	var count int64

	boil.SetCount(q.Query)
	boil.SetLimit(q.Query, 1)

	err := boil.ExecQueryOne(q.Query).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if thumbnails exists")
	}

	return count > 0, nil
}

// FileG pointed to by the foreign key.
func (t *Thumbnail) FileG(mods ...qm.QueryMod) fileQuery {
	return t.File(boil.GetDB(), mods...)
}

// File pointed to by the foreign key.
func (t *Thumbnail) File(exec boil.Executor, mods ...qm.QueryMod) fileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=$1", t.FileID),
	}

	queryMods = append(queryMods, mods...)

	query := Files(exec, queryMods...)
	boil.SetFrom(query.Query, "files")

	return query
}



// LoadFile allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (r *ThumbnailR) LoadFile(e boil.Executor, singular bool, maybeThumbnail interface{}) error {
	var slice []*Thumbnail
	var object *Thumbnail

	count := 1
	if singular {
		object = maybeThumbnail.(*Thumbnail)
	} else {
		slice = *maybeThumbnail.(*ThumbnailSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		args[0] = object.FileID
	} else {
		for i, obj := range slice {
			args[i] = obj.FileID
		}
	}

	query := fmt.Sprintf(
		`select * from "files" where "id" in (%s)`,
		strmangle.Placeholders(count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load File")
	}
	defer results.Close()

	var resultSlice []*File
	if err = boil.BindFast(results, &resultSlice, fileTitleCases); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice File")
	}

	if len(fileAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		if object.R == nil {
			object.R = &ThumbnailR{}
		}
		object.R.File = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FileID == foreign.ID {
				if local.R == nil {
					local.R = &ThumbnailR{}
				}
				local.R.File = foreign
				break
			}
		}
	}

	return nil
}



// ThumbnailsG retrieves all records.
func ThumbnailsG(mods ...qm.QueryMod) thumbnailQuery {
	return Thumbnails(boil.GetDB(), mods...)
}

// Thumbnails retrieves all the records using an executor.
func Thumbnails(exec boil.Executor, mods ...qm.QueryMod) thumbnailQuery {
	mods = append(mods, qm.From("thumbnails"))
	return thumbnailQuery{NewQuery(exec, mods...)}
}

// ThumbnailFindG retrieves a single record by ID.
func ThumbnailFindG(id string, selectCols ...string) (*Thumbnail, error) {
	return ThumbnailFind(boil.GetDB(), id, selectCols...)
}

// ThumbnailFindGP retrieves a single record by ID, and panics on error.
func ThumbnailFindGP(id string, selectCols ...string) *Thumbnail {
	retobj, err := ThumbnailFind(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// ThumbnailFind retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func ThumbnailFind(exec boil.Executor, id string, selectCols ...string) (*Thumbnail, error) {
	thumbnailObj := &Thumbnail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(selectCols), ",")
	}
	query := fmt.Sprintf(
		`select %s from "thumbnails" where "id"=$1`, sel,
	)

	q := boil.SQL(exec, query, id)

	err := q.BindFast(thumbnailObj, thumbnailTitleCases)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from thumbnails")
	}

	return thumbnailObj, nil
}

// ThumbnailFindP retrieves a single record by ID with an executor, and panics on error.
func ThumbnailFindP(exec boil.Executor, id string, selectCols ...string) *Thumbnail {
	retobj, err := ThumbnailFind(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Thumbnail) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Thumbnail) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Thumbnail) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are inferred (i.e. name, age)
// - All columns with a default, but non-zero are inferred (i.e. health = 75)
func (o *Thumbnail) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no thumbnails provided for insertion")
	}

	var err error
	loc := boil.GetLocation()
	currTime := time.Time{}
	if loc != nil {
		currTime = time.Now().In(boil.GetLocation())
	} else {
		currTime = time.Now()
	}

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	wl, returnColumns := strmangle.InsertColumnSet(
		thumbnailColumns,
		thumbnailColumnsWithDefault,
		thumbnailColumnsWithoutDefault,
		boil.NonZeroDefaultSet(thumbnailColumnsWithDefault, thumbnailTitleCases, o),
		whitelist,
	)

	ins := fmt.Sprintf(`INSERT INTO thumbnails ("%s") VALUES (%s)`, strings.Join(wl, `","`), strmangle.Placeholders(len(wl), 1, 1))

	if len(returnColumns) != 0 {
		ins = ins + fmt.Sprintf(` RETURNING %s`, strings.Join(returnColumns, ","))
		err = exec.QueryRow(ins, boil.GetStructValues(o, thumbnailTitleCases, wl...)...).Scan(boil.GetStructPointers(o, thumbnailTitleCases, returnColumns...)...)
	} else {
		_, err = exec.Exec(ins, boil.GetStructValues(o, thumbnailTitleCases, wl...)...)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, ins)
		fmt.Fprintln(boil.DebugWriter, boil.GetStructValues(o, thumbnailTitleCases, wl...))
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into thumbnails")
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Thumbnail record. See Update for
// whitelist behavior description.
func (o *Thumbnail) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Thumbnail record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Thumbnail) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Thumbnail, and panics on error.
// See Update for whitelist behavior description.
func (o *Thumbnail) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Thumbnail.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Thumbnail) Update(exec boil.Executor, whitelist ...string) error {
	loc := boil.GetLocation()
	currTime := time.Time{}
	if loc != nil {
		currTime = time.Now().In(boil.GetLocation())
	} else {
		currTime = time.Now()
	}

	o.UpdatedAt = currTime

	if err := o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}

	var err error
	var query string
	var values []interface{}

	wl := strmangle.UpdateColumnSet(thumbnailColumns, thumbnailPrimaryKeyColumns, whitelist)
	if len(wl) == 0 {
		return errors.New("models: unable to update thumbnails, could not build whitelist")
	}

	query = fmt.Sprintf(`UPDATE thumbnails SET %s WHERE %s`, strmangle.SetParamNames(wl), strmangle.WhereClause(len(wl)+1, thumbnailPrimaryKeyColumns))
	values = boil.GetStructValues(o, thumbnailTitleCases, wl...)
	values = append(values, o.ID)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	result, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update thumbnails row")
	}

	if r, err := result.RowsAffected(); err == nil && r != 1 {
		return errors.Errorf("failed to update single row, updated %d rows", r)
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q thumbnailQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q thumbnailQuery) UpdateAll(cols M) error {
	boil.SetUpdate(q.Query, cols)

	_, err := boil.ExecQuery(q.Query)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for thumbnails")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ThumbnailSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ThumbnailSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ThumbnailSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ThumbnailSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = strmangle.IdentQuote(name)
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	args = append(args, o.inPrimaryKeyArgs()...)

	sql := fmt.Sprintf(
		`UPDATE thumbnails SET (%s) = (%s) WHERE (%s) IN (%s)`,
		strings.Join(colNames, ", "),
		strmangle.Placeholders(len(colNames), 1, 1),
		strings.Join(strmangle.IdentQuoteSlice(thumbnailPrimaryKeyColumns), ","),
		strmangle.Placeholders(len(o)*len(thumbnailPrimaryKeyColumns), len(colNames)+1, len(thumbnailPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in thumbnail slice")
	}

	if r, err := result.RowsAffected(); err == nil && r != ln {
		return errors.Errorf("failed to update %d rows, only affected %d", ln, r)
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Thumbnail) UpsertG(update bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), update, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Thumbnail) UpsertGP(update bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), update, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Thumbnail) UpsertP(exec boil.Executor, update bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, update, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Thumbnail) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no thumbnails provided for upsert")
	}
	loc := boil.GetLocation()
	currTime := time.Time{}
	if loc != nil {
		currTime = time.Now().In(boil.GetLocation())
	} else {
		currTime = time.Now()
	}

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	var err error
	var ret []string
	whitelist, ret = strmangle.InsertColumnSet(
		thumbnailColumns,
		thumbnailColumnsWithDefault,
		thumbnailColumnsWithoutDefault,
		boil.NonZeroDefaultSet(thumbnailColumnsWithDefault, thumbnailTitleCases, o),
		whitelist,
	)
	update := strmangle.UpdateColumnSet(
		thumbnailColumns,
		thumbnailPrimaryKeyColumns,
		updateColumns,
	)
	conflict := conflictColumns
	if len(conflict) == 0 {
		conflict = make([]string, len(thumbnailPrimaryKeyColumns))
		copy(conflict, thumbnailPrimaryKeyColumns)
	}

	query := generateUpsertQuery("thumbnails", updateOnConflict, ret, update, conflict, whitelist)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, boil.GetStructValues(o, thumbnailTitleCases, whitelist...))
	}
	if len(ret) != 0 {
		err = exec.QueryRow(query, boil.GetStructValues(o, thumbnailTitleCases, whitelist...)...).Scan(boil.GetStructPointers(o, thumbnailTitleCases, ret...)...)
	} else {
		_, err = exec.Exec(query, boil.GetStructValues(o, thumbnailTitleCases, whitelist...)...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for thumbnails")
	}

	if err := o.doAfterUpsertHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteP deletes a single Thumbnail record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Thumbnail) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Thumbnail record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Thumbnail) DeleteG() error {
	if o == nil {
		return errors.New("models: no Thumbnail provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Thumbnail record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Thumbnail) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Thumbnail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Thumbnail) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Thumbnail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := o.inPrimaryKeyArgs()

	sql := `DELETE FROM thumbnails WHERE "id"=$1`

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from thumbnails")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q thumbnailQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q thumbnailQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no thumbnailQuery provided for delete all")
	}

	boil.SetDelete(q.Query)

	_, err := boil.ExecQuery(q.Query)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from thumbnails")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, and panics on error.
func (o ThumbnailSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ThumbnailSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Thumbnail slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ThumbnailSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ThumbnailSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Thumbnail slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(thumbnailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	args := o.inPrimaryKeyArgs()

	sql := fmt.Sprintf(
		`DELETE FROM thumbnails WHERE (%s) IN (%s)`,
		strings.Join(strmangle.IdentQuoteSlice(thumbnailPrimaryKeyColumns), ","),
		strmangle.Placeholders(len(o)*len(thumbnailPrimaryKeyColumns), 1, len(thumbnailPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from thumbnail slice")
	}

	if len(thumbnailAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Thumbnail) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Thumbnail) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Thumbnail) ReloadG() error {
	if o == nil {
		return errors.New("models: no Thumbnail provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Thumbnail) Reload(exec boil.Executor) error {
	ret, err := ThumbnailFind(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

func (o *ThumbnailSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

func (o *ThumbnailSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

func (o *ThumbnailSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ThumbnailSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ThumbnailSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	thumbnails := ThumbnailSlice{}
	args := o.inPrimaryKeyArgs()

	sql := fmt.Sprintf(
		`SELECT thumbnails.* FROM thumbnails WHERE (%s) IN (%s)`,
		strings.Join(strmangle.IdentQuoteSlice(thumbnailPrimaryKeyColumns), ","),
		strmangle.Placeholders(len(*o)*len(thumbnailPrimaryKeyColumns), 1, len(thumbnailPrimaryKeyColumns)),
	)

	q := boil.SQL(exec, sql, args...)

	err := q.BindFast(&thumbnails, thumbnailTitleCases)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ThumbnailSlice")
	}

	*o = thumbnails

	return nil
}

// ThumbnailExists checks if the Thumbnail row exists.
func ThumbnailExists(exec boil.Executor, id string) (bool, error) {
	var exists bool

	sql := `select exists(select 1 from "thumbnails" where "id"=$1 limit 1)`

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if thumbnails exists")
	}

	return exists, nil
}

// ThumbnailExistsG checks if the Thumbnail row exists.
func ThumbnailExistsG(id string) (bool, error) {
	return ThumbnailExists(boil.GetDB(), id)
}

// ThumbnailExistsGP checks if the Thumbnail row exists. Panics on error.
func ThumbnailExistsGP(id string) bool {
	e, err := ThumbnailExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ThumbnailExistsP checks if the Thumbnail row exists. Panics on error.
func ThumbnailExistsP(exec boil.Executor, id string) bool {
	e, err := ThumbnailExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

func (o Thumbnail) inPrimaryKeyArgs() []interface{} {
	var args []interface{}
	args = append(args, o.ID)
	return args
}

func (o ThumbnailSlice) inPrimaryKeyArgs() []interface{} {
	var args []interface{}

	for i := 0; i < len(o); i++ {
		args = append(args, o[i].ID)
	}

	return args
}





// SetFile of the thumbnail to the related item.
// Sets t.R.File to related.
// Adds t to related.R.Thumbnails.
func (t *Thumbnail) SetFile(exec boil.Executor, insert bool, related *File) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	oldVal := t.FileID
	t.FileID = related.ID
	if err = t.Update(exec, "file_id"); err != nil {
		t.FileID = oldVal
		return errors.Wrap(err, "failed to update local table")
	}

	if t.R == nil {
		t.R = &ThumbnailR{
			File: related,
		}
	} else {
		t.R.File = related
	}

	if related.R == nil {
		related.R = &FileR{
			Thumbnails: ThumbnailSlice{t},
		}
	} else {
		related.R.Thumbnails = append(related.R.Thumbnails, t)
	}
	return nil
}
