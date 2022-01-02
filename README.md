# Hello
Golang Echo and html template. 


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

#### Test the packages

```
$ go test github.com/ockibagusp/hello/test 
```

or,

```
$ go test github.com/ockibagusp/hello/test -v
```


## TODO List
- mock unit test
- CSRF: no testing
- session.GetUser() to session.GetAuth()
- session: IsAdmin, IsUser and IsAuth
- list pagination with next, previous, first and last
- Mutex: BankAccount
- too much

## Operating System (with me)
### Linux:
- Fedora 35 Workstation

### Go: 
- go1.16.11 linux/amd64

### MySQL: 
- mysql  Ver 8.0.27 for Linux on x86_64 (Source distribution)


### Bahasa Indonesia
Der Schlaganfall 03.10.2018-heute. Dirilis 7 Januari 2020. Coding ini sedikit lupa. Lalu pun, itu Bahasa Inggris lupa lagi. Perlahan-lahan dari stroke. Amin.

### English (translate[.]google[.]co[.]id)
Stroke: 03 10 2018-today. Released January 7, 2020. This coding is a little forgotten. Then again, this English forgot again. Little by little from stroke. Amen.

---

Copyright © 2020 by Ocki Bagus Pratama
