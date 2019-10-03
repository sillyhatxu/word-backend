package model

import "time"

type Vocabulary struct {
	Id               int64     `mapstructure:"id"`
	Word             string    `mapstructure:"word"`
	CreatedTime      time.Time `mapstructure:"created_time"`
	LastModifiedTime time.Time `mapstructure:"last_modified_time"`
}