package chunk

import (
	"log"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

type Chunk struct {
	ID        string    `json:"id"`
	FileID    string    `json:"file_id"`
	Size      int       `json:"size"`
	Hash      string    `json:"hash"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const findByIDSQL = `
	SELECT
	file_id, size, hash, position, created_at, updated_at
	FROM chunks
	WHERE id = $1`

const findByFileIDSQL = `
	SELECT
	id, size, hash, position
	FROM chunks
	WHERE file_id = $1
	ORDER BY position asc`

const insertSQL = `
	INSERT INTO chunks
	(file_id, size, hash, position)
	VALUES
	($1, $2, $3, $4)
	RETURNING id`

const updateChunkSQL = `
	UPDATE chunks
	SET file_id = $2, size = $3, hash = $4, position = $5
	WHERE id = $1;
`

// FindByID returns a chunk with the specified id.
func FindByID(tx *sql.Tx, id string) (*Chunk, error) {
	var c Chunk
	c.ID = id
	err := tx.QueryRow(findByIDSQL, id).Scan(
		&c.FileID, &c.Size, &c.Hash, &c.Position, &c.CreatedAt, &c.UpdatedAt,
	)
	return &c, err
}

// FindByFileID return all chunks with the specified FileID.
// TODO: cleanup
func FindByFileID(tx *sql.Tx, id string) (*[]Chunk, error) {
	var chunks []Chunk
	rows, err := tx.Query(findByFileIDSQL, id)
	if err != nil {
		return &chunks, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Chunk

		if err := rows.Scan(&c.ID, &c.Size, &c.Hash, &c.Position); err != nil {
			log.Fatal(err)
		}

		chunks = append(chunks, c)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &chunks, err
}

// Create a chunk inside of a transaction.
func (c *Chunk) Create(tx *sql.Tx) error {
	err := tx.
		QueryRow(insertSQL, c.FileID, c.Size, c.Hash, c.Position).
		Scan(&c.ID)

	return err
}