# Notex

Backend Assignment for Speer Technologies (https://speer.io)

Table of Contents

- [What is Notex?](#what-is-notex)
- [How to run the project?](#how-to-run-the-project)
  - [Prerequisites](#prerequisites)
  - [Project structure](#project-structure)
  - [Running application](#running-application)
- [How to run the tests?](#how-to-run-the-tests)
- [How to use the API?](#how-to-use-the-api)
- [Design Decisions](#design-decisions)

# What is Notex?

Notex (i know very !creative name) is a backend for a simple, multi user, note taking application. It is created as an assignment for Speer Technologies (https://speer.io).

It is created using GoLang and MongoDB. Using golang's [GoFiber](https://gofiber.io/) framework and MongoDB Go Driver.

# How to run the project?

## Prerequisites:

- [GoLang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)

## Project structure

The `main` entry point to this project is `main.go` file.

- `config` folder: is used to store api and app config.
- `data` folder: is used to store database related operations.
- `docs` folder: is used to store `Swagger 2.0` documentation.
- `password` folder: is used to store password encryption and decryption related operations.
- `tests` folder: is used to store tests.
- `user` folder: is used to store user related operations.

## Running application

Clone the project

```bash
git clone https://github.com/KunalSin9h/notex
```

Download dependencies

```bash
go mod download
```

Run the project

> [!CAUTION]
> Make sure you have `mongodb` running on `localhost:27017`.

```bash
go run main.go
````

The application will start on port `7000` by default.

And uses `mongodb` database at `localhost:27017` by default.

Open `http://localhost:7000/swagger` to view the swagger documentation and test api with ease.

> [!IMPORTANT]
> This application comes with `Swagger 2.0` api documentation. So we can test the APIs without needing any external tools.

# How to run the tests?

To run the tests, run the following command

```bash
go test -v ./...
```

Test need a mongodb database running on `localhost:27017` by default.

# How to use the API?

API Documentation can be found at `http://localhost:7000/swagger`.

To use the application, one need to `signup` first.

After signing up, one can `login` with username and password. To get the `Access Token`. The `Access Token` is then stored in the `Session` which will expire in 1 hour (default).

`Login` endpoint is protected with `Basic Auth` authorization scheme.

And `Access Token` is used to authenticate all other endpoints.

After login successfully, `Access Token` is setup in the `Cookie` header, as well as return in the response body.

For browser clients, Since `Access Token` is setup in the `Cookie` header, it will be automatically sent with every request. Thats why when using `swagger` documentation, after login, every api is accessible without needing to setup `Access Token` manually.

# Design Decisions

1. Why Golang?

In simple words, go is simple and fast. It is simple as Javascript / Python and 2x faster then Java.
Thats why Go is becoming popular day by day for backend development.

> [!NOTE]
> Apart from that, I also have experience with Javascript / Typescript for backend development with Express.js framework.

1. Why GoFiber?

GoFiber is the fastest HTTP framework for Go. It is inspired by Express.js and make development straight forward and easy. Especially for testing.

2. Why MongoDB?

The Entities are modeled like this:

![Entities Diagram](https://i.imgur.com/KNTesiB.png)

Here `Notes` are independent of `Users`, with only `AuthorID` as a reference to `User`.

`Users` has access to `Notes`. Its just an array of `NotesID`.

With this simple and efficient design, we dont need the complex modeling of `SQL` databases.

Also, for text searching mongodb is better for fast and efficient searching.

## Security & Instrumentation

The application is secure by Encrypted Cookies, Argon2ID password hashing, and Timing Safe Password Comparison.

Application is Instrumented with logs and metrics(future).

---

Feel free to contact me at `kunal@kunalsin9h.com` or discord at @kunalsin9h.
