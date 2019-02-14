# BUILD STAGE
FROM golang:latest as build

# copy
WORKDIR /go/src/github.com/mchmarny/myevents/
COPY . /src/

# dependancies
WORKDIR /src/
ENV GO111MODULE=on
RUN go mod download

# build
WORKDIR /src/cmd/service/
RUN CGO_ENABLED=0 go build -v -o /myevents



# RUN STAGE
FROM alpine as release
RUN apk add --no-cache ca-certificates

# app executable
COPY --from=build /myevents /app/

# copy static dependancies
COPY --from=build /src/templates /app/templates/
COPY --from=build /src/static /app/static/

# run
WORKDIR /app/
ENTRYPOINT ["./myevents"]