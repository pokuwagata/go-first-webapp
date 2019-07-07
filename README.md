# README

## local

build docker image

`docker build -t first_webapp .`

local test

`docker run -d -e "PORT=8000" -p 80:8000 first_webapp`

heroku deploy from local

`heroku container:push web`

`heroku container:release web`

## circleci

```bash
$HEROKU_APP_NAME : heroku app name
$HEROKU_API_KEY : `$ heroku auth:token`
$HEROKU_LOGIN : heroku email address
```

## heroku

heroku define `$PORT`
