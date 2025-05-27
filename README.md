# SaaStack

## Communication between Two plugins

```mermaid
sequenceDiagram
  autonumber
  participant P as ProductPlugin
  participant C as Core
  participant U as UserInterfaceHandler
  participant E as EmailInterfaceHandler
  participant UP as ClerkPlugin
  participant EP as EmailSESPlugin

  Note left of P: {pluginId: Clerk, config: config, data: data}
  P ->> C : user.getPreference({data})
  C ->> U : {data}
  U ->> UP : {data}
  UP ->> U : {response}
  UP ->> C: {response}
  C ->> P : {response}

  P ->> C : Email.send({data})
  Note left of P: {pluginId: EmailSES, config: config, data: data}
  C ->> E : {data}
  E ->> EP : {data}
  EP ->> E : {response}
  E ->> C : {response}
  C ->> P : {reponse}

```

## Architecture

![SaaStack Architecture](./docs/public/architechture.svg)

## Project Strucuture

The project is organized into the following main directories:

- `core/`: Core grpc Server.<http://localhost:8965/626>
- `docs/`: OpenApi spec file for each interface plugin.
- `gen/`: Generated code, likely from protobuf definitions.
- `http-gateway/`: Handles HTTP requests and acts as a gateway to the core services.
- `interfaces/`: Create interface handler for plugins(Handle Plugin Mapping and Calling the Plugin).
- `plugins/`: Actual logic of plugin(Implement proto file)
- `proto/`: Protobuf definitions for each interface plugins.

## How to start project locally

Main package:

1. core/main.go
2. http-gateway/main.go
3. Plugin

   a. plugins/email/custom/main.go

   b. plugins/payment/custom/main.go

### Steps

1. **Build the project:**
   Open your terminal and navigate to the project's root directory. Then, run the following command:

   ```bash
   make build
   ```

   This command will compile the core server, HTTP proxy, and the example plugins, placing the binaries in the `./bin` directory.

2. **Run the services:**
   After a successful build, you can run the individual services. For example, to run the core server:

   ```bash
   ./bin/core
   ```

   Similarly, you can start the HTTP proxy and any plugins you intend to use:

   ```bash
   ./bin/http-proxy
   ```

3. **Run Custom Plugin:**
   This plugins are deploy independently or at compile time.

   ```bash
   ./bin/plugins/payment
   ./bin/plugins/email
   ```

