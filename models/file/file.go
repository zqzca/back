package file

import (
	"database/sql"
	"log"
	"time"
)

type File struct {
	ID        string    `json:"id"`
	Size      int       `json:"size" valid:"required,numeric"`
	Hash      string    `json:"hash" valid:"required,alphanum"`
	Chunks    int       `json:"chunks" valid:"required,numeric"`
	Name      string    `json:"name" valid:"required"`
	Type      string    `json:"type" valid:"required"`
	State     int       `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	Incomplete = iota
	Processing
	Finished
)

const paginationSQL = `
	SELECT
	id, size, hash, chunks, name, type, state, created_at, updated_at
	FROM files
	ORDER BY created_at desc
	OFFSET $1
	LIMIT $2
`

const findByIDSQL = `
	SELECT
	size, hash, chunks, name, type, state, created_at, updated_at
	FROM files
	WHERE id = $1`

const findByHashSQL = `
	SELECT
	id, size, chunks, name, type, state, created_at, updated_at
	FROM files
	WHERE hash = $1`

const insertSQL = `
	INSERT INTO files
	(size, hash, chunks, name, type, state)
	VALUES
	($1, $2, $3, $4, $5, $6)
	RETURNING id
`

const setStateSQL = `
	UPDATE files
	SET state = $2
	WHERE id = $1
	RETURNING state
`

// Create a File inside of a transaction.
func (f *File) Create(tx *sql.Tx) error {
	err := tx.
		QueryRow(insertSQL, f.Size, f.Hash, f.Chunks, f.Name, f.Type, f.State).
		Scan(&f.ID)

	return err
}

// FindByHash returns a File with the specified hash.
func FindByHash(tx *sql.Tx, hash string) (*File, error) {
	var f File
	f.Hash = hash
	err := tx.QueryRow(findByHashSQL, hash).Scan(
		&f.ID, &f.Size, &f.Chunks, &f.Name, &f.Type, &f.State,
		&f.CreatedAt, &f.UpdatedAt,
	)
	return &f, err
}

// FindByID returns a File with the specified id.
func FindByID(tx *sql.Tx, id string) (*File, error) {
	var f File
	f.ID = id
	err := tx.QueryRow(findByIDSQL, id).Scan(
		&f.Size, &f.Hash, &f.Chunks, &f.Name, &f.Type, &f.State,
		&f.CreatedAt, &f.UpdatedAt,
	)
	return &f, err
}

func Pagination(tx *sql.Tx, page int, perPage int) (*[]File, error) {
	var files []File
	var err error
	var rows *sql.Rows

	if perPage == 0 {
		perPage = 20
	}

	page = page - 1
	if page < 0 {
		page = 0
	}

	offset := perPage * page

	if rows, err = tx.Query(paginationSQL, offset, perPage); err != nil {
		return &files, err
	}
	defer rows.Close()

	for rows.Next() {
		var f File

		err = rows.Scan(
			&f.ID, &f.Size, &f.Hash, &f.Chunks, &f.Name, &f.Type, &f.State,
			&f.CreatedAt, &f.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		files = append(files, f)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &files, err
}

// SetState sets the state of the File.
func (f *File) SetState(tx *sql.Tx, state int) error {
	err := tx.QueryRow(setStateSQL, f.ID, state).Scan(&f.State)
	return err
}

func (f *File) Process(tx *sql.Tx) error {
	return f.SetState(tx, Finished)
}
