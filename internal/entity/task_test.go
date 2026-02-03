package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	task, err := NewTask("Aprender Go", "Estudar a linguagem Go.", "low", dueDate)

	assert.NoError(t, err)
	assert.Equal(t, "Aprender Go", task.Title)
	assert.Equal(t, "Estudar a linguagem Go.", task.Description)
	assert.Equal(t, PriorityLow, task.Priority)
	assert.NotZero(t, task.DueDate)
}
