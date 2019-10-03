package model

import "time"

type Longman struct {
	Id               int64     `mapstructure:"id"`
	VocabularyId     int64     `mapstructure:"vocabulary_id"`
	Description      string    `mapstructure:"description"`
	CreatedTime      time.Time `mapstructure:"created_time"`
	LastModifiedTime time.Time `mapstructure:"last_modified_time"`
}
