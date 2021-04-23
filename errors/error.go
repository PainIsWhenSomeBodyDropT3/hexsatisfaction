package errors

import (
	"fmt"
)

// DatabaseError represents errors related with database.
func DatabaseError(details string, err error) error {
	return fmt.Errorf("%s : %v", details, err) // errors.Wrap(err, details)
}
