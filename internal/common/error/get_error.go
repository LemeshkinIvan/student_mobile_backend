package custom_error

import (
	"fmt"
	"time"
)

// default
func GetError(msg string, where string) error {
	// time - msg: context (where)
	t := time.Now().Local().String()
	return fmt.Errorf("%s - %s. where: %s", t, msg, where)
}
