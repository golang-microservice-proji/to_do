# API Docs

Send a post request with the following JSON data to make a new task:

```json
// endpoint: /tasks
// method: post
{
  "Title": "Ungibungi",
  "Deadline": "2024-04-07T23:00:00Z"
}
```

send a get request to get a list of all tasks in the following format:

````json
// endpoint /tasks
// method: get
[
    {
        "ID": "809cc993-16fc-4b1d-bcdb-e97d9796d422",
        "Title": "Ungibungi",
        "Completed": true,
        "Deadline": "2024-04-08T04:30:00+05:30"
    }
]
```


send a put request to /tasks/[taskID] to update the `Completed` status

```json
// endpoint: /tasks/[taskID]
// method: put
// example: http://localhost:8080/tasks/04ee1ba1-869e-406d-a9c4-95053572faf5
// the request body must contain the following:
{
    "Completed": true
}

// whatever you set `Completed to in the above json, the task will get updated to that Completed value. This way, you can mark an already completed task uncompleted as well
// The response you'll get is the updated json object of that task:
{
    "ID": "809cc993-16fc-4b1d-bcdb-e97d9796d422",
    "Title": "Ungibungi",
    "Completed": true,
    "Deadline": "2024-04-08T04:30:00+05:30"
}
```

send a delete request to /tasks/[ID] to delete a task. You don't have to put anything in the request body

```json
// endpoint: /tasks/[ID]
// method: delete
// example: http://localhost:8080/tasks/04ee1ba1-869e-406d-a9c4-95053572faf5

you'll receive a 200 OK as a response
```
````

