# liveops-tool

A simple project with Gin framework which contains an endpoint to return a list of users based on their score.

### PostgreSQL

Create a new table for users:

`CREATE TABLE public.users
(
    name character(100) COLLATE pg_catalog."default",
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    country character(50) COLLATE pg_catalog."default",
    total_score integer,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)`

### Run

run `BD_NAME=XXX BD_PWD=XXX go run main.go`

##### Docker

`docker build -t liveopsdocker .`

`docker run -p 8080:8080 liveopsdocker`

### Notes

If you don't want to install PostgreSQL, you can replace _UserPostgrestDirectory_ with _UserDirectory_.

`var userDirectoryService = user.NewService(user.UserPostgrestDirectory{Database: db})`

`var userDirectoryService = user.NewService(user.UserDirectory{})`
