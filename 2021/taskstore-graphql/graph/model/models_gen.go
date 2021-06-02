// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type NewTask struct {
	Text string    `json:"Text"`
	Tags []string  `json:"Tags"`
	Due  time.Time `json:"Due"`
}

type Task struct {
	ID   int       `json:"Id"`
	Text string    `json:"Text"`
	Tags []string  `json:"Tags"`
	Due  time.Time `json:"Due"`
}
