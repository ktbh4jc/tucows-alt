# Kyle Submission for Tucows Interview Exercise 

This is where I will walk you through my general approach to the project. 

## Other files
- design.md: where I plan to go over my design as it evolves with the process
- prompt.md: A copy of the readme file on the [interview-exercise-alt](https://github.com/tucows/interview-exercise-alt) repo
- setup.md: A file containing the instructions to run this project on a unix-like environment. 
- tickets.md: Where I plan on defining my tickets and including any proof-it-works screenshots
- mermaidGraphIconsReference.md: A scratchpad for playing with mermaid chart where I can save potentially useful stuff. I plan on removing it before final release, but if I think I can clean it up into something neat I might leave it alone. 

## Technologies used (and why I chose them)
- Mermaid Chart: https://mermaid.js.org/#/
  - I wanted a free diagraming-as-code tool that would allow me to commit my diagrams and render them in VS Code
  - Pros: Free, renders in VS Code and Github
  - Cons: Less fine-tuned control over where exactly things go than in a more traditional drag-and-drop diagraming tool


## Questions (and my assumed answers)
Spot of context: I received this takehome on a Friday where I already had afternoon plans and want to make the 4 day turnaround, meaning I am not going to have a chance to ask my questions before the weekend. As such, I am going to just list the questions I would normally ask here, along with my assumptions.

1) Is there any expectation for how this should be run? IE: Local native, local docker, aws, etc?
    > I am guessing local is fine, but I am going to try to dockerize my services.
2) Do we have any concept of a user or of authentication?
    > I'm gonna guess no as there is nothing in the prompt. 

## A note on the lack of unit tests
Normally I am a very unit-test forward engineer. However for this project I spent a lot of my time on the architecture and learning new technologies. This is my first forray into writing code with Kafka and the first time I have done Docker and Postgres work in about 4 years. Additionally I tend to follow the principal of unit testing logic more than architecture. That said, I am including a collection of REST calls though Thunder Client that demonstrate end-to-end testing. 

For an example of a takehome where I was more familiar with the tech and had time to add unit tests, see https://github.com/ktbh4jc/esusu-domain-challenge 

## References
This is my first time going from 0 to 1 on a multi-microservice go project without a principal or staff engineer to reach out to for help, and it's my first time using Kafka, so I am going to need to read up on some stuff as I go. In order to keep track, I am going to include my references here. 

TC-03: 
  https://docs.confluent.io/kafka-clients/go/current/overview.html 
  https://hub.docker.com/_/golang

DB-01:
  https://hub.docker.com/_/postgres 