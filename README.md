# Go API Template

Cloud production ready template repository for different golang projects based on MVC Layered Architecture and HTTP RESTful API.

Built on top of [Gin-Gonic](https://github.com/gin-gonic/gin) framework.

The application is built using Cloud Native Application best practices.

## Includes

- MVC Layered Architecture
- Middleware CORS & CSRF, Logger, Recovery
- Example API (Users API)
- Swagger Docs
- Configuration Management
- Database Connections
    - MySQL
    - PostgreSQL
- Container Orchestration
    - Docker
    - Docker Compose
- Mappers
- Request & Response Models
- Hot Reload for Development (Air)
- Migrations

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Local Environment

1. Clone the repository:

```bash
git clone git@github.com:KostLinux/go-template.git
```

2. Run the application

```
docker-compose up
```

3. Enjoy the development!

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE.md) file for details.