FROM golang:1.22 

WORKDIR /app

COPY . .

EXPOSE 3000
RUN go mod tidy
RUN make build

CMD make run