package models

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Model is basic model struct
// All biz model extends from it
type Model struct {
	ID        string   `json:"id" xorm:"pk notnull uuid"`
	CreatedAt DateTime `json:"created_at" xorm:"created created_at"`
	UpdatedAt DateTime `json:"updated_at" xorm:"updated updated_at"`
	DeletedAt DateTime `json:"deleted_at" xorm:"deleted deleted_at"`
}

// BeforeInsert is a hook function before create a record
func (model *Model) BeforeInsert() {
	model.ID = fmt.Sprintf("%s", uuid.NewV4())
}
