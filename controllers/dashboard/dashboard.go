package dashboard

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/zqzca/echo"
	"github.com/zqzca/back/db"
	"gopkg.in/nullbio/null.v4"
)

type dashboardEntry struct {
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	ThumbnailID null.String `json:"thumb_id"`
	CreatedAt   time.Time   `json:"created_at"`
}

type dashboardData struct {
	Entries *[]dashboardEntry `json:"data"`
	Page    int               `json:"current_page"`
	Total   int               `json:"total_pages"`
}

const paginationSQL = `
	SELECT
	f.name, t.id, f.slug, f.created_at
	FROM files AS f
	LEFT JOIN thumbnails as t
	ON t.file_id = f.id
	ORDER BY f.created_at DESC
	OFFSET $1
	LIMIT $2
`

const totalPagesSQL = `
	SELECT
	count(*)
	FROM
	files
`

//Index returns a list of files
func Index(c echo.Context) error {
	page, perPage := paginationOptions(c)

	entries, err := pagination(db.Connection, page, perPage)

	if err != nil {
		fmt.Println("failed to fetch page:", err)
		return err
	}

	total := totalPages(db.Connection, perPage)
	var d dashboardData

	d.Entries = entries
	d.Total = total
	d.Page = page

	return c.JSON(http.StatusOK, d)
}

func totalPages(ex db.Executor, perPage int) int {
	var count int

	err := ex.QueryRow(totalPagesSQL).Scan(&count)

	if err != nil {
		return 0
	}

	return int(math.Ceil(float64(count) / float64(perPage)))
}

func pagination(ex db.Executor, page int, perPage int) (*[]dashboardEntry, error) {
	var entries []dashboardEntry
	var err error
	var rows *sql.Rows

	offset := perPage * page

	if rows, err = ex.Query(paginationSQL, offset, perPage); err != nil {
		return &entries, err
	}
	defer rows.Close()

	for rows.Next() {
		var e dashboardEntry

		err = rows.Scan(
			&e.Name, &e.ThumbnailID, &e.Slug, &e.CreatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		entries = append(entries, e)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &entries, err
}

func paginationOptions(c echo.Context) (int, int) {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 0
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		perPage = 10
	}

	if perPage == 0 {
		perPage = 20
	}

	if page < 0 {
		page = 0
	}

	return page, perPage
}