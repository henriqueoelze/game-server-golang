# Game Server (Golang)

A game server implementation in Go that provides player management with secure authentication.

This repo is a personal project to exemplify how to organize a project and to cover some common cases in the game industry.

## Table of Contents
- [Game Server (Golang)](#game-server-golang)
  - [Table of Contents](#table-of-contents)
  - [Architecture](#architecture)
    - [Layer Overview](#layer-overview)
  - [Project Structure](#project-structure)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Configuration](#configuration)
  - [API Documentation](#api-documentation)
    - [Endpoints](#endpoints)
  - [Development](#development)
    - [Available Make Commands](#available-make-commands)
    - [Code Quality Tools](#code-quality-tools)
  - [Future Improvements](#future-improvements)
  - [Contributing](#contributing)
  - [License](#license)

## Architecture

This project follows a Clean Architecture approach with clear separation of concerns between layers. The architecture is designed to be modular, testable, and maintainable.

### Layer Overview

1. **Domain Layer** (`internal/domain/`)
   - Contains the core business entities
   - No dependencies on external packages
   - Pure Go structs representing the business models
   - Example: `Player` entity

2. **Use Case Layer** (`internal/usecase/`)
   - Implements business logic
   - Depends only on domain layer and gateway interfaces
   - Examples: `PlayerUsecase`, `SecurityUsecase`
   - Each use case has its own package for better organization

3. **Gateway Layer** (`internal/gateway/`)
   - Defines interfaces for external dependencies
   - Contains implementations of these interfaces
   - Examples:
     - `PlayerRepository` interface for database operations
     - `Logger` interface for logging
     - Implementations in `sql_lite/` and `logger/` packages

4. **HTTP Layer** (`internal/http/`)
   - Handles HTTP requests and responses
   - Routes definition and handlers
   - Request/Response mapping
   - Middleware implementation
   - No direct business logic

5. **Config Layer** (`internal/config/`)
   - Configuration management
   - Environment variables handling
   - YAML configuration parsing

## Project Structure

```
.
├── cmd/                    # Application entry points
├── config/                 # Configuration files
├── docs/                   # Documentation and examples
├── internal/               # Internal packages
│   ├── config/            # Configuration structures
│   ├── constant/          # Constants and enums
│   ├── core/              # Core utilities
│   ├── domain/            # Business entities
│   ├── gateway/           # Interface definitions and implementations
│   ├── http/              # HTTP handlers and middleware
│   └── usecase/           # Business logic
├── tools/                  # Development tools
└── bin/                   # Compiled binaries
```

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Make

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/henriqueoelze/game-server-golang.git
   ```

2. Install dependencies:
   ```bash
   make setup
   ```

3. Run database migrations:
   ```bash
   make db-migrate
   ```

4. Start the server:
   ```bash
   make run
   ```

## Configuration

Configuration is managed through:
- Environment variables (prefixed with `GAMESERVER_`)
- YAML configuration file (`config/config.yaml`)
- Default values in code

## API Documentation

### Endpoints

1. Health Check
   ```http
   GET /v1/health
   ```

2. Create Player
   ```http
   POST /v1/player
   ```
   - Returns: Player token in `Player-Key` header

3. Get Player
   ```http
   GET /v1/player
   Authorization: <encrypted-player-id>
   ```

## Development

### Available Make Commands

Run make to get the list of all commands and description

### Code Quality Tools

1. **Linting**
   - Uses golangci-lint
   - Configuration in `.golangci.yaml`
   - Strict linting rules enforced

2. **Mocking**
   - Uses mockery
   - Configuration in `.mockery.yml`
   - Automatic mock generation for interfaces

## Future Improvements

1. **Testing**
   - Add unit tests
   - Add integration tests
   - Add performance tests

2. **Documentation**
   - Add code documentation
   - Add architecture diagrams

3. **Features**
   - Add metrics collection
   - Add tracing
   - Add rate limiting

4. **Operations**
   - Add Docker support

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

[MIT License](LICENSE)