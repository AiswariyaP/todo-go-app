# Todo APP
Fully-fledged REST API that exposes GET, POST, DELETE and PUT endpoints that will subsequently allow you to perform the full range of CRUD operations

## Getting started

To get a local copy up and running follow these simple steps.


### Clone repository

    git clone https://github.com/AiswariyaP/todo-go-app.git

### Setting up database

**MongoDB**
- Replace `URI` value in db.go with your connection string in the below provided format.
    
        mongodb://hostname:port/

Create DB name: todo

Create collection name: tasks

### APIs


1. /todo/           - POST - Creates todo
2. /todo/           - GET -  Fetches all todos
3. /todo/:id        - GET - Fetches todo by id
4. /todo/:id        - PUT - Updates todo by id
5. /todo/:id        - DELETE - Deletes book by id




