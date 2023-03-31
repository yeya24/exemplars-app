FROM golang:1.20-alpine

RUN apk add git make

WORKDIR /app

ENV GOPROXY=direct
ENV GOSUMDB=off
COPY . ./
RUN go mod download

COPY *.go ./

RUN make build

CMD [ "/app/main" ]
