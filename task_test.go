package tasks

import (
    "reflect"
    "testing"
)

func TestNewTask(t *testing.T) {
    task_name := "TestTask"
    tsk := NewTask(task_name)
    if reflect.TypeOf(tsk).String() != "tasks.Task" {
        t.Fatalf(`NewTask("%s") did not return a "Task" type.`, task_name)
    }
    if tsk.Name != task_name {
        t.Fatalf(`NewTask("%s") set task name as "%s" instead of "%s".`,
            task_name,
            tsk.Name,
            task_name)
    }
}
