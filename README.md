# Go API Template

Template repository for different golang projects based on MVC Layered Architecture and HTTP RESTful API.

Built on top of [Gin-Gonic](https://github.com/gin-gonic/gin) framework.

## Features

- **MVC Layered Architecture**: The project is structured in a way that it is easy to understand and maintain.
- **Middlewares**: The project has a set of middlewares used for observability and security purposes.
    - **Logging**: The project has a logging middleware that logs the request and responses.
    - **Recovery**: The project has a recovery middleware that recovers from panics and logs the error.
    - **CORS**: The project has a CORS middleware that handles CORS requests.
    - **CSRF**: The project has a basic CSRF middleware that handles CSRF attacks.
- **Configuration**: The project has a configuration package that reads configuration from [config.yaml](config.yaml) file.
- **Graceful Shutdown**: The project has a graceful shutdown mechanism for HTTP server.