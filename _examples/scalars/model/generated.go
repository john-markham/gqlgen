// Code generated by github.com/john-markham/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/john-markham/gqlgen/_examples/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location,omitempty"`
}

type Query struct {
}
