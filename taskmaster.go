package tasks

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    "text/tabwriter"
)

// TaskMaster handles adding and removing tasks to a master "list".
type TaskMaster struct {
    Tasks []Task
}

// NewTaskMaster creates a new TaskMaster object.
func NewTaskMaster() (tm TaskMaster) {
    return tm
}

// RegisterTask adds a Task object to the Tasks array and returns the Task's id.
func (s *TaskMaster) RegisterTask(t Task) (id int) {
    id = len(s.Tasks)
    s.Tasks = append(s.Tasks, t)
    return
}

// DeleteTaskId removes a task from TaskMaster Tasks slice associated
// with the provided id.
func (s *TaskMaster) DeleteTaskId(id int) {
    // Make sure the id actually exists before deleting it.
    if len(s.Tasks) > id {
        s.Tasks = append(s.Tasks[:id], s.Tasks[id+1:]...)
    }
}

// TableDisplayTo writes a table to output. This table is formater
// so that it diplays well on the command line.
func (s *TaskMaster) TableDisplayTo(output io.Writer) {
    w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
    fmt.Fprintln(w, tabWriterHelper("=Id=",
        "=Due=",
        "=Priority=",
        "=Name="))
    for id, tsk := range s.Tasks {
        fmt.Fprintln(w, tabWriterHelper(strconv.Itoa(id),
            tsk.Due.Format("2006-01-02 3:04 PM"),
            strconv.Itoa(tsk.Priority),
            tsk.Name))
    }
    w.Flush()
}
func tabWriterHelper(id string, priority string, name string, due string) string{
    // This adds a blank column of once space for the first column.
    // Makes abit easier to look at, not sure if this should stay though.
    return fmt.Sprintf(" \t%s\t%s\t%s\t%s\t", id, priority, name, due)
}

// SaveToJson writes a json encoded TaskMaster object to a file location p.
func (s *TaskMaster) SaveToJson(p string) {
    d, err := json.MarshalIndent(s, "", "   ")
    if err != nil {
        log.Fatal(err)
        return
    }
    err = os.WriteFile(p, d, 0666)
    if err != nil {
        log.Fatal(err)
        return
    }
}

// LoadFromJson reads a json encoded TaskMaster object from file p
// and populates the receiver TaskMaster with that data.
func (s *TaskMaster) LoadFromJson(p string) {
    d, err := os.ReadFile(p)
    if err != nil {
        log.Fatal(err)
        return
    }
    err = json.Unmarshal(d, s)
    if err != nil {
        log.Fatal(err)
        return
    }
}

// Sort
// func (s *TaskMaster) Sort(fields []string, asc bool) {
//     n
//
// }
