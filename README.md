# todo-app

This is a Create, Read, Update, Delete (CRUD) Application Programming Interface (API) implementation written in Go using the Gorilla HTTP request multiplexer (Mux) library for routing and the encoding/json package for JavaScript Object Notation (JSON) encoding and decoding.

# Endpoints

## Get all tasks
URL: `/tasks`

Method: `GET`

**Response:**

- Status code: `200 OK`

Example:  `localhost:10000/tasks`

Body:

```
[
{
"id": 1,
"name": "Task 1",
"description": "Description of task 1",
"due_date": "2022-01-01"
},
{
"id": 2,
"name": "Task 2",
"description": "Description of task 2",
"due_date": "2022-02-01"
}
...
]
```

## Read a specific task

URL: `/task/{id}`

Method: `GET`

**Response:**

- Status code: `200 OK`

Example: `localhost:10000/task/1`

Body:

```
{
"id": 1,
"name": "Task 1",
"description": "Description of task 1",
"due_date": "2022-01-01"
}
```

**Error:**

- Status code: `404 Not Found`

Example: `localhost:10000/task/10`

Body:

```
{
  "error": "task not found"
}
```

- Status code: `400 Bad Request`

Example: `localhost:10000/task/abc`

Body:

```
{
  "error": "invalid task ID"
}
```

## Create a new task

URL: `/task`

Method: `POST`

**Request Body:**

```
{
"name": "Task 3",
"description": "Description of task 3",
"due_date": "2022-03-01"
}
```

**Response:**

- Status code: `201 Created`

Body:

```
{
"id": 3,
"name": "Task 3",
"description": "Description of task 3",
"due_date": "2022-03-01"
}
```

**Error:**

- Status code: `400 Bad Request`

Body:

```
{
  "error": "Invalid request payload"
}
```

## Update a specific task

URL: `/task/{id}`

Method: `PUT`

Example: `localhost:10000/task/3`

**Request Body:**

```
{
"name": "Updated Task 3",
"description": "Updated description of task 3",
"due_date": "2022-03-15"
}
```

**Response:**

- Status code:`200 OK`

Body:

```
{
"id": 3,
"name": "Updated Task 3",
"description": "Updated description of task 3",
"due_date": "2022-03-15"
}
```

**Error:**

- Status code: `400 Bad Request`

Example: `localhost:10000/task/ids`

Body:

```
{
  "error": "invalid task ID"
}
```

- Status code: `404 Not Found`

Body:

```
{
  "error": "task not found"
}
```

## Delete a specific task

URL: `/task/{id}`

Method: `DELETE`

**Response:**

- Status code: `200 OK`

Body:

```
{
  "result": "successful deletion"
}
```

- Status code:  `400 Bad Request`

Example: `localhost:10000/task/ids`

Body:

```
{
  "error": "invalid task ID"
}
```

- Status code: `404 Not Found`

Body:

```
{
  "error": "task not found"
}
```
