FROM golang:1.22 

WORKDIR /app

COPY . .

RUN go mod tidy
RUN make build

CMD make run