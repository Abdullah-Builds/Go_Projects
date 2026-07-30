package cache

<<<<<<< HEAD
type Item struct {
	Key   string
	Value any
=======
import "time"

type Item struct {
	Value       string
	ExpiresAt time.Time
>>>>>>> 43439c3 (feat: cache version 1)
}
