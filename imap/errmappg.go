package imap

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/radianhanggata/go-pkg/iconst"
)

type Pg struct{}

func (pg *Pg) Read(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return iconst.ErrorRecordNotFound
	}

	pgErr := err.(*pgconn.PgError)

	if err, ok := errMap[pgErr.Code]; ok {
		err.Data = iconst.Response{
			SC:      pgErr.Code,
			Message: pgErr.Detail,
		}
		return err
	}

	return iconst.ErrorInternalServer
}

var errMap = map[string]*iconst.Response{
	"23505": iconst.ErrorDuplicate,
}
