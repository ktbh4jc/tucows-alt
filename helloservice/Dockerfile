FROM golang:1.22 

WORKDIR /app

COPY . .

# RUN apt-get update && apt-get install build-essential -y --no-install-recommends
RUN go mod tidy
RUN make build

CMD make run