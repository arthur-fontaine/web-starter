# Web Starter

A web-starter with e2e type-safety, an Expo application, and a Golang API with hexagonal architecture.

## First Steps

1. **Clone the Repository**
   ```
   git clone https://github.com/arthur-fontaine/web-starter.git
   cd web-starter
   rm -rf .git
   ```

2. **Replace Project Name**
   Globally replace the string `webstarter` with your project name.

3. **Generate Initial Setup**
   ```
   moon :generate
   ```

## What's Included

- **[Moonrepo](https://moonrepo.dev/)**: For running tasks in the monorepo.
- **[Expo App](./apps/app)**: A cross-platform application (web, iOS, and Android).
- **[Go API](./services/api)**: A Golang-based API.
- **[PostgreSQL Database](./services/postgres)**: A PostgreSQL database setup.

## Project Structure

### [./apps/app](./apps/app)

- **Source Directory**: `src`
  - **File-based Routing**: `src/app` (Using Expo Router)
  - **Feature-based Organization**: `src/features`
    - **Components**: `src/features/xxx/components`
    - **Hooks**: `src/features/xxx/hooks`
    - **Main Page**: `src/features/xxx/xxx-page.tsx`
  - **Shared Code**: `src/shared` (e.g., hooks, components, etc.)

### [./services/api](./services/api)

- **Main API Code**: `cmd/server/main.go` (Initializes the API and handles dependency injection)
- **Database Schema**: `db/schema.prisma`
- **Bob Generator Configuration**: `db/bobgen.yaml`
- **Domain Entities**: `internal/domain/xxx/entity.go`
- **Repositories**: `internal/domain/xxx/repository.go`
- **Repository Implementations**: `internal/repository`
- **Service Implementations**: `internal/usecase`
- **Service Definitions**: `proto/xxx/xxx.proto`
- **Docker Configuration**: `docker-compose.yaml`
- **Build Configuration**: `Dockerfile`
- **Task Definitions**: `moon.yml`

### [./services/postgres](./services/postgres)

- **Docker Configuration**: `docker-compose.yaml`

## Recommendations

### [./apps/app](./apps/app)

- Use **TanStack Query**.
- Place every query inside a hook.

### [./services/api](./services/api)

- Implement **Hexagonal Architecture**.