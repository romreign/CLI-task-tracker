package model

import (
	"encoding/json"
	"time"
)

type (
	TaskStatus string
	DateTime   time.Time
)

const (
	TODO        = "todo"
	IN_PROGRESS = "in-progress"
	DONE        = "done"
)

type Task struct {
	id          uint
	description string
	status      TaskStatus
	createdAt   DateTime
	updatedAt   DateTime
}

func NewTask(id uint, description string, status TaskStatus, createdAt DateTime, updatedAt DateTime) *Task {
	return &Task{
		id:          id,
		description: description,
		status:      status,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (t *Task) GetId() uint {
	return t.id
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetStatus() TaskStatus {
	return t.status
}

func (t *Task) GetCreatedAt() DateTime {
	return t.createdAt
}

func (t *Task) GetUpdatedAt() DateTime {
	return t.updatedAt
}

func (t *Task) SetId(id uint) {
	t.id = id
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) SetStatus(status TaskStatus) {
	t.status = status
}

func (t *Task) SetCreatedAt(createdAt DateTime) {
	t.createdAt = createdAt
}

func (t *Task) SetUpdatedAt(updatedAt DateTime) {
	t.updatedAt = updatedAt
}

func (t Task) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(&struct {
		ID          uint       `json:"id"`
		Description string     `json:"description"`
		Status      TaskStatus `json:"status"`
		CreatedAt   DateTime   `json:"createdAt"`
		UpdatedAt   DateTime   `json:"updatedAt"`
	}{
		ID:          t.id,
		Description: t.description,
		Status:      t.status,
		CreatedAt:   t.createdAt,
		UpdatedAt:   t.updatedAt,
	}, "", " ")
}

func (t *Task) UnmarshalJSON(data []byte) error {
	var temp struct {
		ID          uint       `json:"id"`
		Description string     `json:"description"`
		Status      TaskStatus `json:"status"`
		CreatedAt   DateTime   `json:"createdAt"`
		UpdatedAt   DateTime   `json:"updatedAt"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	t.id = temp.ID
	t.description = temp.Description
	t.status = temp.Status
	t.createdAt = temp.CreatedAt
	t.updatedAt = temp.UpdatedAt

	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Format(time.RFC3339) + `"`), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*d = DateTime(time.Time{})
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	*d = DateTime(t)
	return nil
}
