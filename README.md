# Hello
Golang Echo for templates.


### Visual Studio Code

#### Run and Debug: [launch.json](https://github.com/ockibagusp/hello/blob/master/.vscode/launch.json).


## Getting Started
First, clone the repo:
```bash
$ git clone https://github.com/ockibagusp/hello.git
```

### Setting MySQL database

#### Database 
file: hello.sql -> new database: hello

#### Testing
file: hello.sql -> new database: hello_test

## User Tables

| Username | Password |
| --- | --- |
| ockibagusp | user123 |
| sugriwa | user123 |
| subali | user123 |


## Router
This using [router](https://github.com/ockibagusp/hello/blob/master/router/router.go).


### Running app

#### Compile and run Go program
```
$ go run main.go
```

or,

#### Build compiles the packages

```
$ go build
```

- On Linux or Mac:

```
$ ./hello
```

- On  Windows:

```
$ hello.exe
```


## TODO List
- test
- CSRF
- session.GetUser() to session.GetAuth()
- session: IsAdmin, IsUser and IsAuth
- list pagination with next, previous, first and last
- Mutex: BankAccount
- too much


Der Schlaganfall 03.10.2018-heute. Dirilis 7 Januari 2020. Coding ini agak lupa. Bertahap.

---

Copyright © 2020 by Ocki Bagus Pratama
