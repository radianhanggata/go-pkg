package iutils

import (
	"fmt"
	"time"
)

func ConvertTimeToIso(t time.Time) string {
	return fmt.Sprintf("%v", t.Format(time.RFC3339))
}
