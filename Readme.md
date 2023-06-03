# Golang Data management system

#### Run the server

Inside the **cmd** folder run the below command

```go run main.go```

### Folder Structure

```
cmd
|--main.go
|--cmd.exe
pkg
|--config
	|--app.go
|--controllers
	|--controller.go
|--models
	|--data.go
|--routes
	|--data-routes.go
|--utils
	|--utils.go
```

### Paths

GET `/data/` -> Get the All the data

GET `/data/{dataId}` -> Get a particular data

POST `/data/` -> Create new entry into the DB

PUT `/data/{dataId}` -> Update particular data in DB

DELETE `/data/{dataId}` -> Deletes the particular entry in the DB
