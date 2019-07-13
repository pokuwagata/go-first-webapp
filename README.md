# README

[![CircleCI](https://circleci.com/gh/pokuwagata/go-first-webapp/tree/master.svg?style=svg)](https://circleci.com/gh/pokuwagata/go-first-webapp/tree/master)

simple Go + MySQL API Server sample

## Features

- Docker based environment
- Using Docker-Compose for local development
- Testing and Deploying on CircleCI
- Heroku hosting
- Using ClearDB for Production

## Requirements

- Go 1.11 or higher
- ClearDB MySQL Heroku Add-ons

## Getting Started

`$ docker-compose up -d`

access localhost:5000

## CircleCI Settings

set the following enviroment variable

```bash
$HEROKU_APP_NAME : heroku app name
$HEROKU_API_KEY : `$ heroku auth:token`
$HEROKU_LOGIN : heroku email address
```

## Heroku Settings

set the following enviroment variable

```bash
$PORT : heroku defined automatically
```

run `$ heroku config | grep CLEARDB_DATABASE_URL`

```bash
$CLEARDB_USER
$CLEARDB_PASSWORD
$CLEARDB_HOST
$CLEARDB_DBNAME
```
