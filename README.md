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
    "name": "Create project proposal",
    "description": "Write a proposal for the new project",
    "due_date": "2022-02-01"
  },
  {
    "id": 2,
    "name": "Design website layout",
    "description": "Create a layout for the company website",
    "due_date": "2022-03-01"
  },
  {
    "id": 3,
    "name": "Implement payment system",
    "description": "Integrate a payment system into the website",
    "due_date": "2022-04-01"
  },
  {
    "id": 4,
    "name": "Conduct user testing",
    "description": "Gather feedback from user testing to improve the website",
    "due_date": "2022-05-01"
  },
  {
    "id": 5,
    "name": "Launch website",
    "description": "Make the website live for the public",
    "due_date": "2022-06-01"
  },
  {
    "id": 6,
    "name": "Evaluate website performance",
    "description": "Collect data and analyse websit performance",
    "due_date": "2022-07-01"
  }
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
  "name": "Create project proposal",
  "description": "Write a proposal for the new project",
  "due_date": "2022-02-01"
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
  "name": "Test task",
  "description": "This is a test task",
  "due_date": "2022-01-01"
}
```

**Response:**

- Status code: `201 Created`

Example: `localhost:10000/task`

Body:

```
{
  "id": 7,
  "name": "Test task",
  "description": "This is a test task",
  "due_date": "2022-01-01"
}
```

**Request Body:**

```

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

Example: `localhost:10000/task/2`

**Request Body:**

```
{
  "name": "Collect requirements",
  "description": "Conduct research and gather input from stakeholders",
  "due_date": "2022-02-01"
}
```

**Response:**

- Status code:`200 OK`

Example: `localhost:10000/task/2`

Body:

```
{
"id": 2,
"name": "Collect requirements",
"description": "Conduct research and gather input from stakeholders",
"due_date": "2022-02-01"
}
```

**Error:**

- Status code: `400 Bad Request`

Example: `localhost:10000/task/abc`

Body:

```
{
  "error": "invalid task ID"
}
```

- Status code: `400 Bad Request`

Example: `localhost:10000/task/1`

**Request Body:**

```

```

Body:

```
{
  "error": "Invalid task payload"
}
```

## Delete a specific task

URL: `/task/{id}`

Method: `DELETE`

**Response:**

- Status code: `200 OK`

Example: `localhost:10000/task/3`

Body:

```
{
  "result": "successful deletion"
}
```

- Status code: `404 Not Found`

Example: `localhost:10000/task/10`

Body:

```
{
  "error": "task not found"
}
```
