# Setup

## A brief note
I am doing this on my Ubuntu computer, so I will be writing these instructions assuming a unix-like environment. Any specific installation notes will be for Ubuntu. I expect this would work on MacOS as well, but am unable to check any specific changes. Fun enough, I have a new mini pc coming in the mail today so I will be able to really test from zero and insure I am not missing any requirements. 

## Requirements
- Go
  - https://go.dev/doc/install 
- Docker
  - https://docs.docker.com/engine/install/ubuntu/
- Make
  - `sudo apt-get -y install make` for ubuntu, `brew install make` on mac

## Enhancements
These are tools that I used while developing this solution that will make it easier to interact with, but aren't necessarily mandatory. 

- VS Code - IDE I used in development. I also used the following extension, which can be installed by hitting `ctrl` + `p` and pasting in the install string.
  - [Markdown Preview Mermaid Support](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-mermaid)
    - Enables mermaid diagrams to render in VS Code
    - install string: `ext install bierner.markdown-mermaid`
  - [ThunderClient](https://www.thunderclient.com/)
    - Postman, but for VS Code.
    - install string: `ext install rangav.vscode-thunder-client`
    - Note: I am using the free version 
  - [Markdown All in One](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one)
    - Needed in order for the checkboxes in tickets.md to render properly
    - install string `ext install yzhang.markdown-all-in-one`

## Instructions