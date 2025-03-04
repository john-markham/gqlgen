// Code generated by github.com/john-markham/gqlgen, DO NOT EDIT.

package chat

import (
	"time"
)

type Message struct {
	ID           string        `json:"id"`
	Text         string        `json:"text"`
	CreatedBy    string        `json:"createdBy"`
	CreatedAt    time.Time     `json:"createdAt"`
	Subscription *Subscription `json:"subscription"`
}

type Mutation struct {
}

type Query struct {
}

type Subscription struct {
}
