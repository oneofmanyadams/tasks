package tasks

import (
    "reflect"
    "strconv"
    "testing"
)

func TestNewTaskMaster(t *testing.T) {
    expected_type := "tasks.TaskMaster"

    tsk_mstr := NewTaskMaster()
    if reflect.TypeOf(tsk_mstr).String() != expected_type {
        t.Fatalf(`NewTaskMaster() did not return a "%s" type.`, expected_type)
    }
}

func TestRegisterTask(t *testing.T) {
    total_tasks := 10

    tsk_mstr := NewTaskMaster()
    for i := 0; i < total_tasks; i++ {
        tsk_mstr.RegisterTask(NewTask(strconv.Itoa(i)))
    }
    if len(tsk_mstr.Tasks) < total_tasks {
        t.Fatalf(`Did not register all tasks. Want:"%d" Got:"%d"`,
            total_tasks,
            len(tsk_mstr.Tasks))
    }
    if len(tsk_mstr.Tasks) > total_tasks {
        t.Fatalf(`Registered too many tasks. Want:"%d" Got:"%d"`,
            total_tasks,
            len(tsk_mstr.Tasks))
    }
    for n, tsk := range tsk_mstr.Tasks {
        if tsk.Name != strconv.Itoa(n) {
            t.Fatalf(`Tasks not registed in correct order. Want:"%s" Got:"%s".`,
                strconv.Itoa(n),
                tsk.Name)
        }
    }
}
