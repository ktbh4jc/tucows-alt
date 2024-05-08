# Tickets

One upside of doing this via git is that I am able to break requirements up into "tickets" and map them 1:1 with my PRs. Tickets are not necesarily done in any particular order and are subject to being changed as I work. 

### Key: 
- TC-XX relates to tickets that cross the microservice border
- OM-XX tickets focus on the Order Manager microservice
- PM-XX tickets focus on the Payment Manager microservice
- DB-XX tickets focus on DB setup


### TC tickets
- [x] TC-01: Initial setup of documentation
  - Write up my initial documentation framework
- [x] TC-02: Define initial tickets
  - Convert the list of requirements from the prompt into a series to tickets to work on one by one
- [x] TC-03 Create API gateway and sample microservice
  - sample microservice will just log "Hello World!"
  - dockerize the things
  - api should publish to Kafka
  - sample service should read from it.

### OM Tickets
- [ ] OM-01: Build Dockerfile for Order Manager
  - Include update to root docker-compose
  - Once the API can hit the order manager, this ticket is done
- [ ] OM-02: Define order type
  - POST to API with a JSON order
  - For now, just log the order
- [ ] OM-03: Publish order to async message broker (leaning Kafka, but we will see)

### PM Tickets
- [ ] PM-01 Build Dockerfile for Payment Manager 
- [ ] PM-02 Receive message from async message broker
- [ ] PM-03 Validate payment and push result to DB

### DB Tickets
- [ ] DB-01 Stub out DB services
- [ ] DB-02 Add Postgress docker container 
- [ ] DB-03 Create Database Tables and pre-populate 