# Ascii-Art-Web

## Members
- El Mehdi Belfkih
- Larbi Mergaoui
- Said Oubaaisse

### Description

**ascii-art-web** is a project related to ASCII art! The purpose of this project is to run a basic server from the terminal on localhost. It uses a basic REST API with HTTP GET and POST requests and dynamic HTML page templates.

### How it Works
- On the client side, a user can enter text in a text bar and click "submit." The browser sends an HTTP POST request to the server.
- On the server side, the server handles the request and returns an HTML page containing the text in ASCII art format.
- The server will process the input and respond with a dynamic ASCII art representation of the requested text.

### Project Structure

#### requirements
- **banners**: This folder contains `file.txt`, which includes ASCII characters (32 -> 126) in different formats: `standard`, `shadow`, and `thinkertoy`. This file is used to fetch the alphabet based on the client's choice.
- **error**: This folder contains functions that serve HTTP error responses and status codes. The error responses are generated dynamically using templates.
- **internal**: This folder contains functions for configuring the server, which handles and processes client requests.

#### static
- This folder contains CSS files for styling the web page.

#### templates
- This folder contains dynamic HTML page templates. These templates inject the correct data and serve it to the client.

#### main.go
- The main function handles the core operations and starts the process of running the server.

#### Dockerfile
- The Dockerfile configures the image to run the container.

#### docker-compose.yml
- Docker Compose is used as a controller to help manage and run multiple containers easily.

## Installation

Follow these steps to install and set up the project:

1. Clone the repository:
    ```sh
    git clone https://github.com/elmehdibelfkih/ascii-art-web.git
    ```

2. Navigate into the project directory:
    ```sh
    cd ascii-art-web
    ```
3. Run the Go server:
    ```sh
    go run main.go
    ```
For Docker setup:

## Build and run the container
1. run container
    ```sh
    docker compose up
    ```
## Usage

Once the server is running, open your browser and visit `http://localhost:8080`. You can input text and see it converted into ASCII art.