##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /todo-list

COPY . ./
RUN go mod download

RUN go build -o /todo-list ./cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /todo-list /todo-list

EXPOSE 8080

USER nonroot:nonroot

ENV DB_PASSWORD=cdd37bad254efce82674faf409cf9d17b74a9bc1a74be76090ee3db79092cc88

ENTRYPOINT ["/todo-list"]