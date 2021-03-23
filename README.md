# Redis rate-limiting example

![alt text](https://github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang/blob/master/preview.png?raw=true)

## How it works

This app will block connections from a client after surpassing certain amount of requests (default: 10) per time (default: 10 sec)

The application will return after each request the following headers. That will let the user know how many requests they have remaining before the run over the limit.

```
X-RateLimit-Limit: 10
X-RateLimit-Remaining: 9
```

On the 10th run server should return an HTTP status code of **429 Too Many Requests**

### Available commands

```
"SETEX", "GET", "DEL", "INCR"
```

## Hot to run it locally?

### Prerequisites

-   Golang - v1.15
-   Docker - v19.03.13 (optional)
-   Docker Compose - v1.26.2 (optional)

### Local installation

```
git clone https://github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang/

# copy file and set proper data inside
cp .env.example .env

# install dependencies
go get

# run docker compose or install redis manually
docker-compose up -d --build

# open on browser
http://localhost:5000

```

### Env variables:

-   API_HOST - host that server is listening on, default is all interfaces
-   API_PORT - port that server is listening on, default is 5000
-   API_PUBLIC_PATH - frontend location
-   REDIS_HOST - host for redis instance
-   REDIS_PORT - port for redis instance, default is 6379
-   REDIS_PASSWORD - password for redis. Running redis by docker-compose, there's no password. Running by https://redislabs.com/try-free/ you have to pass this variable.

## Deployment

To make deploys work, you need to create free account in https://redislabs.com/try-free/ and get Redis' instance informations - REDIS_HOST, REDIS_PORT and REDIS_PASSWORD. You must pass them as environmental variables (in .env file or by server config, like `Heroku Config Variables`).

### Google Cloud Run

[![Run on Google
Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run/?git_repo=https://github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang.git)

### Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

### Vercel

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/git/external?repository-url=https://github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang&env=API_HOST,API_PORT,API_PUBLIC_PATH,REDIS_HOST,REDIS_PORT,REDIS_PASSWORD)
