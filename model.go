package gitdb

import (
	"time"
)

//Model interface describes methods GitDB supports
type Model interface {
	GetSchema() *Schema
	//Validate validates a Model
	Validate() error
	//IsLockable informs GitDb if a Model support locking
	IsLockable() bool
	//GetLockFileNames informs GitDb of files a Models using for locking
	GetLockFileNames() []string
	//ShouldEncrypt informs GitDb if a Model support encryption
	ShouldEncrypt() bool
	//BeforeInsert is called by gitdb before insert
	BeforeInsert() error
}

//TimeStampedModel provides time stamp fields
type TimeStampedModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

//BeforeInsert implements Model.BeforeInsert
func (m *TimeStampedModel) BeforeInsert() error {
	stampTime := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = stampTime
	}
	m.UpdatedAt = stampTime

	return nil
}

type model struct {
	Version string
	Indexes map[string]interface{}
	Data    Model
}

func wrap(m Model) *model {
	return &model{
		Version: RecVersion,
		Data:    m,
	}
}

func (m *model) GetSchema() *Schema {
	return m.Data.GetSchema()
}

func (m *model) Validate() error {
	return m.Data.Validate()
}
func (m *model) IsLockable() bool {
	return m.Data.IsLockable()
}
func (m *model) ShouldEncrypt() bool {
	return m.Data.ShouldEncrypt()
}
func (m *model) GetLockFileNames() []string {
	return m.Data.GetLockFileNames()
}

func (m *model) BeforeInsert() error {
	err := m.Data.BeforeInsert()
	m.Indexes = m.GetSchema().indexes
	return err
}
