<p align="center">
  <a href="https://codecov.io/gh/yavurb/mobility-payments" >
    <img src="https://codecov.io/gh/yavurb/mobility-payments/graph/badge.svg?token=wtqF1EfDG9"/>
  </a>
</p>

# Mobility Payments

The Mobility Payments project is a backend service that provides a RESTful API to manage payments between users.

## Quick Start :rocket:

To start the project, you just need to have docker and docker-compose installed on your machine. Then, you can run the following command:

```bash
docker compose up
```

## Documentation :book:

For a review of the operations and details on the operations available, head to the [OpenAPI Spec](https://github.com/yavurb/mobility-payments/blob/develop/openapi/openapi.yaml).

## Configuration :wrench:

Mobility Payments uses [Pkl](https://pkl-lang.org) as its configuration language. The primary configuration template file is located at `config/ConfigSchema.pkl`. This file should be used to amend the final configuration file.

> [!Note]
> The final config file should be placed in one of the following locations based on the environment:
>
> - `devolopment-config.pkl`
> - `production-config.pkl`

Example configuration:

```pkl
// development-config.pkl

environment = "development"
host = read?("env:HOST") ?? "0.0.0.0"
port = read?("env:PORT")?.toInt() ?? 8910
cors = new Cors {
  allowOrigins = new Listing { "http://localhost:4321" }
  allowMethods = new Listing { "GET" "POST" "PUT" "DELETE" }
}
httpAuth = new HttpAuth {
  JWTSecret = "somesecret"
  HeaderKey = "MobilityPayments-Api-Key"
}
database = new DatabaseConfig {
  URI = read?("env:DATABASE_URI") ?? "postgres://postgres:postgres@localhost:5432/mobility-payments"
  name = read?("env:DATABASE_NAME") ?? "mobility-payments"
}

logLevel = read?("env:LOG_LEVEL")?.trim()?.toLowerCase() ?? "debug"
```

> Once the configuration file is ready, you can execute any command listed below.

> [!warning]
> For convenience purposes, the development configuration file is already created. You can find it at `config/development-config.pkl`.
> However, this practice is discouraged in a production environment.

## Commands :hammer:

Mobility Payments uses a Makefile to manage its commands. You can run the following commands:

> [!Warning]
> The commands must be prefixed with `make`.

| Command | Arguments | Action |
|:--------------:|:---------------------------------------------:|:-----------------------------------------------:|
| `run`          | -                                             | Run the server                                  |
| `build`        | -                                             | Build the server into a binary                  |
| `dev`          | -                                             | Run the server with hot-reloading               |
| `test`         | -                                             | Run tests                                       |
| `docker_build` | [env(development,production)\|pkl_version\]   | Build a Docker image of the server              |
| `docker_run`   | inherited from `docker_build`                 | Build and run a Docker container of the server  |
| `gen_config`   | -                                             | Generate the types from the configuration file  |

> There are more commands available. You can find them in the `Makefile`.

## Project Structure :open_file_folder:

Mobility Payments tries to follow the [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) principles. The project follows the following structure:

```txt
.
├── cmd
│   └── mobility
├── config
│   └── app_config
├── internal
│   ├── app
│   ├── payments
│   │   ├── application
│   │   ├── domain
│   │   └── infrastructure
│   │       ├── adapters
│   │       └── ui
│   └── pkg
│       └── publicid
└── scripts
```

The three main directories are:

- `cmd`: Contains the entry point to the application.
- `internal`: Contains the application's modules.
- `config`: Contains the configuration files.

Inside the `internal` directory, are define Mobility Payments modules. Each module is divided into three subdirectories:

- `application`: Contains the use cases and orchestrates the business logic.
- `domain`: Contains the definition of the business logic.
- `infrastructure`: Contains the implementation details. Here, the repository and the UI are defined.

> [!Note]
> The `app` module is a special module that contains the main application logic.
> Here is defined the server definition and the application's configuration.
>
> The `pkg` module contains the public ID package. This package is used to generate unique public IDs.
