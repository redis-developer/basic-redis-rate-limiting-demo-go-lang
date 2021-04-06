# Redis rate-limiting example

![rate-limiting example screenshot](https://github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang/blob/master/preview.png?raw=true)

## How it works

This app will block connections from a client after surpassing certain amount of requests (default: 10) per time (default: 10 sec)

The application will return after each request the following headers. That will let the user know how many requests they have remaining before the run over the limit.

On the 10th run server should return an HTTP status code of **429 Too Many Requests**

### Cookies

User identification based on cookies, on first request user will receive a cookie if it not exists
`CookieName: user-limiter`
`CookieValue: md5(<current time>)`
`<current time>` - request time in a format: `2006-01-02 15:04:05.999999999 -0700 MST`

### Redis Commands

- Read requests for user by `user-limiter` cookie: `GET requests.<USER_IDENTIFIER>` - get `USER_IDENTIFIER` from request cookie

  - E.g `GET requests.0cbc6611f5540bd0809a388dc95a615b`

- Set request counter with expired 10 sec if not exist in `requests.<USER_IDENTIFIER>`: `SETEX requests.<USER_IDENTIFIER> 10 0`

  - E.g `SETEX requests.0cbc6611f5540bd0809a388dc95a615b 10 0`

- Increment requests counter for each of user request: `INC requests.<USER_IDENTIFIER>`

  - E.g `INC requests.0cbc6611f5540bd0809a388dc95a615b`

- Get requests number for user: `GET requests.<USER_IDENTIFIER>`
  - E.g `GET requests.0cbc6611f5540bd0809a388dc95a615b`

#### Code for rate limiting

```Go
func (c Controller) AcceptedRequest(user string, limit int) (int, bool) {
	key := c.key(user)

	if _, err := c.r.Get(key); err == redis.Nil {
		err := c.r.Set(key, "0", time.Second * time.Duration(limit))
		if err != nil {
			log.Println(err)
			return 0, false
		}
	}

	if err := c.r.Inc(key); err != nil {
		log.Println(err)
		return 0, false
	}

	requests, err := c.r.Get(key)
	if err != nil {
		log.Println(err)
		return 0,false
	}
	requestsNum, err := strconv.Atoi(requests)
	if err != nil {
		log.Println(err)
		return 0, false
	}

	if requestsNum > limit {
		return requestsNum, false
	}

	return requestsNum, true
}
```

Where `c` corresponds to the active controller and `c.r` is a Redis client.

### Response

#### Status codes

`200 - OK` - responded `PONG`

`406 - Not Acceptable` - could not read cookie from request, it may have effect when cookies not allowed on browser side

`429 - To Many Requests` - user send more than 10 requests / 10sec

#### Headers

`X-RateLimit-Limit: 10` - allowed number of limits per 10sec

`X-RateLimit-Remaining: 9` - number of left request in 10sec window

### Available commands

```
"SETEX", "GET", "DEL", "INCR"
```

## Hot to run it locally?

### Prerequisites

- Golang - v1.15

### Local installation

```
git clone https://github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang/

# copy file and set proper data inside
cp .env.example .env

# install dependencies
go get

# run
go run main.go

# open on browser
http://localhost:5000

```

### Env variables:

- API_HOST - host that server is listening on, default is all interfaces
- API_PORT - port that server is listening on, default is 5000
- API_PUBLIC_PATH - frontend location
- REDIS_HOST - host for redis instance
- REDIS_PORT - port for redis instance, default is 6379
- REDIS_PASSWORD - password for redis. Running redis by docker-compose, there's no password. Running by https://redislabs.com/try-free/ you have to pass this variable.

# Deployment

To make deploys work, you need to create free account in https://redislabs.com/try-free/ and get Redis' instance informations - REDIS_HOST, REDIS_PORT and REDIS_PASSWORD. You must pass them as environmental variables (in .env file or by server config, like `Heroku Config Variables`).

### Google Cloud Run

[![Run on Google
Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run/?git_repo=https://github.com/redis-developer/basic-redis-rate-limiting-demo-go-lang.git)

### Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)
