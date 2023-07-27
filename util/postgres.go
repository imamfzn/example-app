package util

import (
	"errors"

	"github.com/lib/pq"
)

func PostgresUniqueViolation(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		// https://github.com/lib/pq/blob/v1.10.5/error.go#L78
		return pqErr.Code.Name() == "unique_violation"
	}
	return false
}
