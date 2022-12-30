## Go With PostgreSql
This is a simple project to test connection between Golang and PostgreSql database nothing else.

## Usages
If you want to use this repo in your local machine floww the below instruction.
- I am using [Go](https://go.dev/) version 1.19.4 for this repo make sure you install same version or latest version.
- I am using [pgx](https://github.com/jackc/pgx) v5 third party package to connect with local postgresql. So if you want to download the package in your local machine you have to type the below command:
```bash
go get github.com/jackc/pgx/v5
```
or you can download the package by entering the below command
``` bash
go mod tidy
```

- Before compile the repo please make sure you have install postgresql on your local machine.

- To compile the project or see the result in your terminal enter the below command:
```bash
go run main.go
```

## Note:
***Before run the main file please change the database connection uri to your own values. The username, password, host, port, and mydatabase parameters should be set to the correct values for your database.***

## support
if you like my repo and you want more project like this don't forget to give me star.

*If you have any suggestion for this repository feel free to contribute.*
