package dto

import "github.com/CostaFelipe/task-api/internal/entity"

type TaskFilter struct {
	Completed *bool
	Priority  *entity.Priority
	Page      int
	Limit     int
}
