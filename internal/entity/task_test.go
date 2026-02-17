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
		task, err := NewTask("Aprender Go", "Estudar a linguagem Go.", "low", 1)

		assert.Error(t, err)
		assert.Nil(t, task)
		assert.Equal(t, errTitleIsRequired, task.Title)
	})

	t.Run("required description", func(t *testing.T) {
		task, err := NewTask("Aprender Go", "", "low", 1)

		assert.Error(t, err)
		assert.Nil(t, task)
		assert.Equal(t, errDescriptionIsRequired, task.Description)
	})

	t.Run("required id", func(t *testing.T) {
		task, err := NewTask("Aprender Go", "Estudar a linguagem Go.", "low", 0)

		assert.Error(t, err)
		assert.Nil(t, task)
		assert.Equal(t, errIDIsRequired, task.UserID)
	})
}
