# Wallester Task

Forked from [https://github.com/ekateryna-tln/wallester-task](https://github.com/ekateryna-tln/wallester-task) and containerized incl. database seeding

Clone the repo and run (docker should be installed on your machine)
```bash
docker-compose up --build
```
Open [http://0.0.0.0:8080](http://0.0.0.0:8080) in browser to see working application

#### Edit conflicts
To prevent record edit conflicts database lock or concurrency control can be used:
- [https://www.postgresql.org/docs/current/mvcc.html](https://www.postgresql.org/docs/current/mvcc.html)
- [https://stackoverflow.com/questions/17768608/what-is-the-practical-use-of-timestamp-column-in-sql-server-with-example](https://stackoverflow.com/questions/17768608/what-is-the-practical-use-of-timestamp-column-in-sql-server-with-example)
---
![preview](https://github.com/svirmi/wallester-task/blob/master/wallester.png)
---
This repository contains a web application example to work with a customer object.

- Build in GO version 1.17.1
## Live running app example could be found [here](http://www.golang.studyfield.org/en)

# Project structure

## cmd
Contains entry point, main setting, middleware and routers.
### Please fill main settings in the main.go file to run the application:
- const portNumber = ":8080" //should be changed if it needs
- const dbName = ""
- const dbUser = ""
- const dbPass = ""
- const dbHost = ""
- const dbPort = ""

## internal
Contains all app functionality

## migrations
Contains database migrations.
The app uses [Soda CLI](https://gobuffalo.io/en/docs/db/toolbox).
Soda CLI should be installed additional.
Also the app could be used without running migrations.
Please find the database script in the wallester_db_query.sql file

## static
Contains static element of the app as css, javascript and images

## templates
Contain app templates

# Tests
Unit test examples could be found in the form package:
internal/forms/forms_test.go
Database functional tests could be found in the repository package:
internal/repository/dbrepo/postgres_test.go
### For database tests separated database should be used. In this case: wallester_tests. It must have the same structure.
### Tests will remove data from the table. In case of using the main database data will be gone.

# Comments
Since this application is an example please note:
- implemented basic search functionality. In real projects it should be done according to business requirements.
  In case strict search will be used to improve database speed appropriate fields should be indexed.
- implemented UI pagination. This kind of implementation fits small amounts of data.
  In other cases the pagination should be done from the backend side (for example with offset and limit or auto incremented field)
- settings are in the main.go file, but they should be separated for live and dev environments and moved to the additional files for live projects.
- implemented not all tests.