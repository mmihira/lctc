# lctc
A toy project to explore using golang as an API backend.
Display tweets from the twitter stream API.

We use :
- Postgres (jmoiron/sqlx go adapter)
- React/Redux/ReduxSaga
- gorilla/websocket
- Docker-compose

## Table of Contents

* [Requiremets](#install)
* [Usage](#usage)

## Requirements

You will go, docker and node.

## Config

A yml config file is required with twitter credentials. Example
`
CONSUMER_KEY: "1234"
CONSUMER_SECRET: "1234"
ACCESS_TOKEN: "1234"
ACCESS_TOKEN_SECRET: "1234"
`

## Usage

Instructions are given for running locally

Start up postgres               ` cd infra; docker-compose up -d; cd ..` <br>
Build lctc dockerfile           ` cd server; docker build -t lctc:1.0 ./; cd ..` <br>
Run lctc_server image. Replace <PATH_TO_CONFIG> with the path to the config file
`
	docker stop lctc_server;
	docker rm lctc_server;
	docker run \
  	--name lctc_server \
		--restart always \
  	--log-driver syslog \
		-e PG_HOST_URL=lctc-postgres \
		--network=infra_lctc_net \
    --mount type=bind,source=<PATH_TO_CONFIG>,target=/config.yml \
		-p 8080:8080 \
		-d \
    lctc:1.0;
` <br>

Build the front end            ` cd app; yarn install; ` <br>
Start the front end            ` yarn run start ` <br>
Open  localhost:8071 in your browser
