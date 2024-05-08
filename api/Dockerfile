FROM golang:1.22 

WORKDIR /app

COPY . .

EXPOSE 8080
RUN apt-get update && apt-get install build-essential -y --no-install-recommends
RUN apt install build-essential -y --no-install-recommends
RUN go mod tidy
RUN make build

# FROM golang:1.22 
# COPY --from=builder /app/api /api

# COPY wait-for-it.sh /app/wait-for-it.sh
# RUN chmod +x /app/wait-for-it.sh
# CMD ["/app/wait-for-it.sh", "kafka:9092", "--", "/api-gateway"]
CMD make run