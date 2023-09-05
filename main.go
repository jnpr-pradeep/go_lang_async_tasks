package main

import (
    "fmt"
    "log"
    "time"
    "strconv"
    "net/http"
)

var (
     TASK_STORE = make(map[string]Task)
)


type Task struct {
    ID        string
    Name      string
    Status    string
    UpdatedAt time.Time
}

func generateID() string {
    l := strconv.Itoa(len( TASK_STORE) + 1)
    fmt.Printf("New ID:= %s\n", l)
    return l
}

func insertTask(t Task) Task {
    t.ID = generateID()
     TASK_STORE[t.ID] = t
    return t
}

func updateTaskStatus(taskID string, newStatus string) Task {
    t :=  TASK_STORE[taskID]
    t.Status = newStatus
     TASK_STORE[taskID] = t
    return t
}

func readTaskStatus(taskID string) Task {
    t :=  TASK_STORE[taskID]
    return t
}

func createTask(w http.ResponseWriter, r *http.Request) {
    // Insert a new task
    t := Task{Name: fmt.Sprintf("Task %v", generateID()) , Status: "Pending"}

    t = insertTask(t)

    fmt.Printf("Inserted task with ID %s\n", t.ID)
    go func() {
        // Sleep for 2 seconds
        time.Sleep(60 * time.Second)

        // Update the status of the task
        t := updateTaskStatus(t.ID, "Completed")
        fmt.Printf("Task status updated, status=%s\n", t.Status)
    }()
    fmt.Printf("Returning from ServiceMethod for id=%s\n", t.ID)
    msg := fmt.Sprintf("Task ID=%s", t.ID)
    fmt.Fprintln(w, msg)
}

func getTaskStatus(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/get_task_status/"):]
    fmt.Printf("Entered into getTaskStatus with id=%s\n", id)

    t := readTaskStatus(id)

    fmt.Println("Returning from getTaskStatus for id=", t.ID)
    msg := fmt.Sprintf("Task ID=%v, status=%v", t.ID, t.Status)
    fmt.Fprintln(w, msg)
}

func main() {
    http.HandleFunc("/start_task", createTask)
    http.HandleFunc("/get_task_status/", getTaskStatus)

    log.Println("Server listening on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }

}
