FROM golang:1.22 

WORKDIR /app

COPY . .

EXPOSE 8080
RUN go mod tidy
RUN make build

CMD make run