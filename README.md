# README

build docker image

`docker build -t first_webapp .`

local test

`docker run -d -e "PORT=8000" -p 80:8000 first_webapp`

heroku deploy

`heroku container:push web`

`heroku container:release web`
