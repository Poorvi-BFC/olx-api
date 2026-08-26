# OLX API

Backend foundation for a marketplace-style listing platform, built in Go.

> This repository is currently in the project setup stage. The executable is a
> small Go proof of life that prints `hey!!`; the HTTP API and persistence layer
> will be added in subsequent milestones.

## Day 1: System Design & Setup

- [x] **Learn architecture basics:** Understand monolithic and microservices
	architectures, including their trade-offs.
- [x] **Define requirements and HLD:** Identify functional and
	non-functional requirements and outline a high-level design using PostgreSQL,
	Cloudflare R2, and a CDN.
- [x] **Project initialization:** Create the Go module, application entry point,
	Makefile targets, and build output directory.

## Project Direction

The first version will use a modular monolith. This keeps development and
deployment simple while the domain is still evolving. The code can be split
into independently deployable services later if scale or team boundaries make
that worthwhile.

### Planned architecture

```text
Client
	|
	v
CDN and API edge
	|
	v
Go API
	|--------------------|
	v                    v
PostgreSQL        Cloudflare R2
(users, listings,   (listing images)
 transactions)
```

- **Go API:** Handles authentication, listings, search, images, and marketplace
	workflows.
- **PostgreSQL:** Stores transactional and relational data such as users,
	listings, categories, and orders.
- **Cloudflare R2:** Stores uploaded images and other media without placing
	large binary objects in PostgreSQL.
- **CDN:** Delivers public media close to users and reduces load on the API and
	object storage.

## Functional Requirements

The planned API will support:

- User registration, authentication, and profile management
- Creating, updating, viewing, and deleting listings
- Listing images uploaded to object storage
- Searching and filtering listings by category, location, price, and keywords
- Favorites, messaging, and listing status management

## Non-Functional Requirements

- Clear package boundaries and maintainable Go code
- Secure authentication and authorization
- Input validation and consistent error responses
- Fast read paths for listing discovery
- Reliable media storage with CDN caching
- Observability through structured logs, metrics, and health checks
- Automated tests and repeatable local development commands

## Getting Started

### Prerequisites

- Go `1.27` or newer
- GNU Make, or an equivalent environment for running the Makefile commands

### Run locally

```bash
make run
```

### Build

```bash
make build
```

The binary is written to `bin/api` (`bin/api.exe` on Windows).

### Clean build output

```bash
make clean
```

## Repository Structure

```text
.
├── cmd/
│   └── api/
│       └── main.go    # Application entry point
├── bin/               # Local build output
├── Makefile           # Build, run, and clean commands
├── go.mod             # Go module definition
└── README.md
```

## Roadmap

1. Add an HTTP server with health-check and version endpoints.
2. Introduce configuration loading and structured logging.
3. Add PostgreSQL migrations and repository interfaces.
4. Implement listing CRUD and search.
5. Integrate Cloudflare R2 for image uploads.
6. Add authentication, tests, API documentation, and deployment automation.

## Development Principle

Learn, build, and push in small, verifiable increments. Each milestone should
leave the project runnable and easy to understand.
