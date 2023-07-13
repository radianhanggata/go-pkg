package idberrhdl

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/radianhanggata/go-pkg/iecho"
	"gorm.io/gorm"
)

func HandleError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return iecho.ErrorRecordNotFound
	}

	pgError := err.(*pgconn.PgError)

	data := make(map[string]string, 0)
	data["code"] = pgError.Code
	data["message"] = pgError.Detail

	switch pgError.Code {
	case "23505":
		out := iecho.ErrorDuplicate
		out.Data = data
		return out
	}

	return iecho.ErrorInternalServer
}
