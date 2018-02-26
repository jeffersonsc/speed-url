package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// MyURL struct
type MyURL struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id"`
	LongURL     string        `json:"long_url,omitempty" bson:"long_url"`
	ShortURL    string        `json:"short_url,omitempty" bson:"short_url"`
	Key         string        `json:"key,omitempty" bson:"key"`
	Count       int           `josn:"count,omitempty" bson:"count"`
	CreatedAt   time.Time     `json:"created_at,omitempty" bson:"created_at"`
	AccessCalls []Access      `json:"access" bson:"access"`
}
