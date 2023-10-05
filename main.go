package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	DueDate string `json:"due_date"`
}

var tasks []Task
var currentID int

type App struct {
	Router *mux.Router
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/tasks", app.getTasks).Methods("GET")
	app.Router.HandleFunc("/task/{id}", app.readTask).Methods("GET")
	app.Router.HandleFunc("/task", app.createTask).Methods("POST")
	app.Router.HandleFunc("/task/{id}", app.updateTask).Methods("PUT")
	app.Router.HandleFunc("/task/{id}", app.deleteTask).Methods("DELETE")
}

func (app *App) Initialise(initialTasks []Task, id int) {
	tasks = initialTasks
	currentID = id
	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
}
func main() {
	app := App{}
	tasks, id := CreateInitialTasks()
	app.Initialise(tasks, id)
	app.Run("localhost:10000")
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	errorMessage := map[string]string{"error": err}
	sendResponse(w, statusCode, errorMessage)
}

func (app *App) getTasks(writer http.ResponseWriter, request *http.Request) {
	// Get all tasks from the model
	tasks, err := getTasks()
	if err != nil {
		// It returned an error - send it to caller with a 500 status code
		sendError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	// Send the task list with a 200 (OK) status.
	sendResponse(writer, http.StatusOK, tasks)
}

func (app *App) createTask(writer http.ResponseWriter, r *http.Request) {
	// Declare an instance of Task which we will
	// attempt to read the data sent by the caller's request.
	var p Task

	// Use a json.Decoder to decode the []byte data in the response and
	// store the result in `p` - the instance of a task struct
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		// The payload did not contain something
		// that looks like the correct JSON for a task.
		// That's a user error sending the wrong data.
		sendError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// Create the task, which will se p.ID (it's a pointer receiver)
	// with a new, unused ID.
	err = p.createTask()
	if err != nil {
		// Task could not be created - that's a server error
		sendError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	// Send the updated task info back to the caller
	// with a status of Created, because that's what we did!
	sendResponse(writer, http.StatusCreated, p)
}

func (app *App) readTask(writer http.ResponseWriter, request *http.Request) {
	// Get the URL parameters
	// vars will be map[string]string
	vars := mux.Vars(request)
	// Try to get an integer from the "id" argument as "key"
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If that fails (like it's eg. "abc"), then send 400 (BadRequest)
		// and message "invalid task ID" to the caller
		sendError(writer, http.StatusBadRequest, "invalid task ID")
		return
	}

	// Create a Task struct with only the ID we received
	t := Task{ID:key}
	// Call the getTask receiver that will either fill in the task
	// details or return an error because the task with that ID isn't found
	err = t.getTask()
	if err != nil {
		// We got an error.
		// Send a 404 (NotFound) status and the text of the
		// error returned by getTask
		sendError(writer, http.StatusNotFound, err.Error())
		return
	}
	// All good. Send the filled-in task to the caller.
	sendResponse(writer, http.StatusOK, t)
}

func (app *App) updateTask(writer http.ResponseWriter, request *http.Request) {
	// Get the URL parameters
	// vars will be map[string]string
	vars := mux.Vars(request)
	// Try to get an integer from the "id" argument as "key"
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If that fails (like it's eg. "abc"), then send 400 (BadRequest)
		// and message "invalid task ID" to the caller
		sendError(writer, http.StatusBadRequest, "invalid task ID")
		return
	}
	// Declare instance of Task struct
	var t Task
	// Try to decode the caller's JSON into the task struct
	err = json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		// Failed - not a correct Task payload,
		// send expected error details
		sendError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// Set task's ID from URL parameter
	t.ID = key
	// Call the model's implementation to do the work
	err = t.updateTask()
	if err != nil {
		// Model returned an error - send it to caller with 500 (InternalServerError)
		sendError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	// Return edited task to caller
	sendResponse(writer, http.StatusOK, t)
}

func (app *App) deleteTask(writer http.ResponseWriter, request *http.Request) {
	// Get the URL parameters
	// vars will be map[string]string
	vars := mux.Vars(request)
	// Try to get an integer from the "id" argument as "key"
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If that fails (like it's eg. "abc"), then send 400 (BadRequest)
		// and message "invalid task ID" to the caller
		sendError(writer, http.StatusBadRequest, "invalid task ID")
		return
	}
	// Initialise a Task struct with the given ID
	t := Task{ID: key}
	// Call the deleteTask receiver
	err = t.deleteTask()
	if err != nil {
		// Failed - not not found,
		// Send expected error details
		sendError(writer, http.StatusNotFound, err.Error())
		return
	}
	// Send a map with the expected key and value to the caller along with 200 (OK) status.
	// Recall that sendResponse takes interface{} as its payload argument,
	// which means we can send almost anything and it will be converted to JSON.
	// We are calling it here in the same way that sendError calls it.
	sendResponse(writer, http.StatusOK, map[string]string{"result": "successful deletion"})
}
