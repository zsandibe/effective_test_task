# Effective Mobile Test Task


## Clone the project

```
$ git clone https://github.com/zsandibe/effective_test_task
$ cd effective_test_task
```

## Launch a project

```
$ make run
```

## Execute migrations

```
$ make migrate-up
$ make migrate-down
```

OR 
```
$ make docker-up
$ make docker-down
```

## API server provides the following endpoints:
* `GET /api/v1/{id}` - returns a car details by id
* `GET /api/v1/list` - returns a cars list by filters(regNum,mark,model,year,owner`s credentials)
* `POST /api/v1/add` - adds a car
* `PUT /api/v1/update/{id}` - updates a car details by id
* `DELETE /api/v1/{id}` - deletes a car by id

# .env file
## API configuration

```
API_URL=YOUR_TESTING_URL
API_KEY=YOUR_TESTING_KEY
```

## Server configuration

```
SERVER_HOST=localhost
SERVER_PORT=8888
```

## Postgres configuration

```
DRIVER=
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

