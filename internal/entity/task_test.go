package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	task, err := NewTask("Aprender Go", "Estudar a linguagem Go.", "low", 1)

	assert.NoError(t, err)
	assert.Equal(t, "Aprender Go", task.Title)
	assert.Equal(t, "Estudar a linguagem Go.", task.Description)
	assert.Equal(t, PriorityLow, task.Priority)
	assert.Equal(t, 1, task.UserID)
}

func TestTaskValidation(t *testing.T) {
	t.Run("required title", func(t *testing.T) {
		task, err := NewTask("", "Estudar a linguagem Go.", "low", 1)
		assert.Error(t, err)
		assert.Nil(t, task)
		assert.Equal(t, errTitleIsRequired, err)
	})

	t.Run("required description", func(t *testing.T) {
		task, err := NewTask("Aprender Go", "", "low", 1)

		assert.Error(t, err)
		assert.Nil(t, task)
		assert.Equal(t, errDescriptionIsRequired, err)
	})

	t.Run("priotiry empty", func(t *testing.T) {
		task, _ := NewTask("Aprender Go", "Estudar a linguagem Go.", "", 1)
		assert.Equal(t, PriorityMedium, task.Priority)
	})
}
