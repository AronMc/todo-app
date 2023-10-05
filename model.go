package main

import "errors"

func CreateInitialTasks() ([]Task, int) {
	tasks := []Task{
		{ID: 1, Name: "Create project proposal", Description: "Write a proposal for the new project", DueDate: "2022-02-01"},
		{ID: 2, Name: "Design website layout", Description: "Create a layout for the company website", DueDate: "2022-03-01"},
		{ID: 3, Name: "Implement payment system", Description: "Integrate a payment system into the website", DueDate: "2022-04-01"},
		{ID: 4, Name: "Conduct user testing", Description: "Gather feedback from user testing to improve the website", DueDate: "2022-05-01"},
		{ID: 5, Name: "Launch website", Description: "Make website live for the public", DueDate: "2022-06-01"},
		{ID: 6, Name: "Evaluate website performance", Description: "Collect data and analyse websit performance", DueDate: "2022-07-01"},
	}
	return tasks, 6
}

func getTasks() ([]Task, error) {
	return tasks, nil
}

func (t *Task) createTask() error {
	currentID = currentID + 1
	t.ID = currentID
	tasks = append(tasks, *t)
	return nil
}

func (t *Task) getTask() error {
	// Get the task ID from the receiver argument
	id := t.ID

	// Range over existing tasks
	for _, task := range tasks {
		if task.ID == id {
			// Matching ID - fill in other fields of *t
			t.DueDate = task.DueDate
			t.Name = task.Name
			t.Description = task.Description
			// return no error
			return nil
		}
	}
	
	// If we get here, the task isn't found
	// so return the correct error text as an error object.
	return errors.New("task not found")
}

func (t *Task) updateTask() error {
	// Get the task ID from the reciver argument
	id := t.ID

	// Range over existing tasks, capturing the index as well as the task content
	for index, task := range tasks {
		if task.ID == id {
			// Found task - update details
			task.DueDate = t.DueDate
			task.Name = t.Name
			task.Description = t.Description
			// Set updated task back into the list at the index position.
			tasks[index] = task

			// no error
			return nil
		}
	}

	// If we get here, the task isn't found
	// so return the correct error text as an error object.
	return errors.New("task not found")
}

func (t *Task) deleteTask() error {
	// Get the task ID from the receiver argument
	id := t.ID

	// Initialise this to -1 to indicate not found yet
	indexToBeDeleted := -1

	// Range across the task list checking task ID
	for index, task := range tasks {
		if task.ID == id {
			// Found the task, store the index
			indexToBeDeleted = index
			break
		}
	}
	if indexToBeDeleted == -1 {
		// if task.ID == id was never true, then task not found
		return errors.New("task not found")
	}
	// Shuffle the task slice to remove the one to delete by appending
	// a slice of everything up to but not including the index to delete
	// to a slice of everything _after_ the index to delete.
	tasks = append(tasks[:indexToBeDeleted], tasks[indexToBeDeleted+1:]...)

	// No error
	return nil
}
