package healthCheck

import (
	"time"
)

type Response struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}
