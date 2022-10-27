FROM golang:1 as build

RUN go env -w GOPROXY=direct

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download -x && go mod verify

COPY . .
RUN go build -v -o /main cmd/pa-data-collector/main.go

FROM scratch
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]
