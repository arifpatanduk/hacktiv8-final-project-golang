# MYGRAM API

MYGRAM API is a RESTful API for create a post with photo and comment on it. It is built using Golang, Gin framework, and Gorm for database interactions.

## Table of Contents
- [Introduction](#introduction)
  - [Main Packages Used](#main-packages-used)
  - [Folder Structure](#folder-structure)
  - [Database Schema](#database-schema)
- [Installation](#installation)
  - [Install Dependencies](#install-dependencies)
  - [Environment Variables](#environment-variables)
  - [Setup Cloudinary](#setup-cloudinary)
  - [Create Database](#create-database)
  - [Run the Project](#run-the-project)
- [API Documentation](#api-documentation)
- [Author](#author)

## Introduction
### Main Packages Used
| Package Name | Description | Documentation Link |
|--------------|-------------|---------------------|
| `Gin` | Gin is a web framework written in Go. | [Documentation](https://github.com/gin-gonic/gin) |
| `Gorm` | ORM library for Golang, aims to be developer friendly. | [Documentation](https://github.com/go-gorm/gorm) |
| `cloundinary-go` | The Cloudinary Go SDK allows you to quickly and easily integrate your application with Cloudinary. | [Documentation](https://github.com/cloudinary/cloudinary-go) |
| `go-plyaground/ validator` | Package validator implements value validations for structs and individual fields based on tags. | [Documentation](https://github.com/go-playground/validator) |
| `jwt-go` | Implementation of JSON Web Tokens | [Documentation](https://github.com/dgrijalva/jwt-go) |

### Folder Structure
```
├── config
    └── database.go
├── controllers
    ├── commentController.go
    ├── photoController.go
    ├── socialMediaController.go
    └── userController.go
├── middlewares
    ├── authentication.go
    └── authorization.go
├── models
    ├── comment.go
    ├── gormModel.go
    ├── photo.go
    ├── socialMedia.go
    └── user.go
├── routers
    └── route.go
├── utils
    ├── apiResponse.go
    ├── bcrypt.go
    ├── headerValue.go
    ├── jwt.go
    ├── uploadCloudinary.go
    └── validationErrors.go
├── .env
├── go.mod
├── go.sum
└── main.go
```
### Database Schema
![image](https://github.com/arifpatanduk/hacktiv8-mygram/assets/57590616/8b4796a3-266f-4223-9fa5-446e6bf372d5)


## Installation

### Install Dependencies

Make sure you have [Go](https://golang.org/dl/) installed on your machine.

1. Clone the repository

```bash
git  clone  https://github.com/arifpatanduk/hacktiv8-mygram
```

2. Change directory to the project folder

```bash
cd  hacktiv8-mygram
```

3. Install project dependencies

```bash
go  mod  tidy
```

### Environment Variables
1. Copy the `.env.example` and named it `.env`
```
cp .env.example .env
```

2. Setup your own `APP_PORT` and `API_SECRET_KEY` 

### Setup Cloudinary
Cloudinary is used to store uploaded photo files.
1. Register and create an account on https://cloudinary.com/
2. Go to Dashboard, copy the `CLOUDINARY_URL` and place it to `CLOUDINARY_URL` variable inside `.env` file

### Create Database

1. Install a PostgreSQL database.
2. Create a new database.
3. Place all the database configuration to `.env` file (host, user, password, port, name).


### Run the Project

```bash
go  run  main.go
```

The API will be accessible at http://localhost:8080.

## API Documentation
  Explore the API using the published Postman collection available at https://documenter.getpostman.com/view/11542567/2sA2r6ZQi3

## Author

[Arif Patanduk](https://github.com/arifpatanduk)
