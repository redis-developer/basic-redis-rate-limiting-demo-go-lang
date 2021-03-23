FROM golang as builder

RUN mkdir /build

COPY . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -o bin .

FROM golang

RUN mkdir /api
RUN addgroup --system redis
RUN adduser --system --disabled-password --no-create-home --home /api --ingroup redis redis
RUN chown redis:redis /api

USER redis

COPY --from=builder /build/bin /api/

WORKDIR /api

LABEL   Name="Redis Rate Limiting Demo"

#Run service
ENTRYPOINT ["./bin"]