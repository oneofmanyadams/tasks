package tasks

import (
    "time"
)

// Task is a simple type that is something that needs to be done.
type Task struct {
    Name string
    Priority int
    Duration int
    Due time.Time
}

// NewTask creates a new Task object.
func NewTask(name string) (t Task) {
    t.Name = name
    t.Priority = 5
    t.Duration = 10
    t.Due = time.Now()
    return t
}
