package files

import (
	"net/http"

	"github.com/zqzca/back/lib"
	"github.com/zqzca/back/models"
	"github.com/zqzca/echo"

	"github.com/vattle/sqlboiler/boil"
	. "github.com/vattle/sqlboiler/queries/qm"
)

func fileExistsWithHash(ex boil.Executor, hash string) (bool, error) {
	// Todo write an exists? for this
	count, err := models.Files(ex, Where("hash=$1", hash)).Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Create creates a file container in the database.
func (f FileController) Create(e echo.Context) error {
	file := &models.File{}

	if err := e.Bind(file); err != nil {
		return err
	}

	if file.NumChunks < 1 {
		return e.NoContent(http.StatusUnprocessableEntity)
	}

	if ok, err := fileExistsWithHash(f.DB, file.Hash); err != nil {
		return err
	} else if ok {
		f.Debug("file exists with hash", "hash", file.Hash)
		return e.NoContent(http.StatusConflict)
	}

	f.Debug("file doesnt exist with hash", "hash", file.Hash)
	file.State = lib.FileIncomplete

	if err := file.Insert(f.DB); err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, file)
}
