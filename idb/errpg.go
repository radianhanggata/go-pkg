package idb

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/radianhanggata/go-pkg/ictx"
)

type Pg struct{}

func (pg *Pg) Read(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ictx.ErrorRecordNotFound
	}

	pgErr := err.(*pgconn.PgError)

	if err, ok := errMap[pgErr.Code]; ok {
		return err.Embed(map[string]string{
			"code":    pgErr.Code,
			"message": pgErr.Detail,
		})
	}

	return ictx.ErrorDB
}

var errMap = map[string]*ictx.Response{
	"23505": ictx.ErrorDBDuplicate,
}
