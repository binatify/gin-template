package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// BaseModel is model interface implements
type BaseModel struct {
	ID bson.ObjectId `bson:"_id" json:"id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	isNewRecord bool `bson:"-"`
}

func NewBaseModel() BaseModel {
	return BaseModel{
		ID:          bson.NewObjectId(),
		isNewRecord: true,
	}
}

func (m *BaseModel) id() bson.ObjectId {
	return m.ID
}

func (m *BaseModel) IsNewRecord() bool {
	return m.isNewRecord
}

func (m *BaseModel) setIsNewRecord(isNew bool) {
	m.isNewRecord = isNew
}

func (m *BaseModel) setCreatedAt(t time.Time) {
	m.CreatedAt = t
}

func (m *BaseModel) setUpdatedAt(t time.Time) {
	m.UpdatedAt = t
}
