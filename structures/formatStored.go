package structures

import (
	"time"
)

type FormatStored struct {
	Id        int
	Task      string
	Status    string
	CreatedAt time.Time
}
