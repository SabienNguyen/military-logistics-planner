# 🪖 Military Logistics Planner

## 🎯 Goal
The Military Logistics Planner is a backend system built in Go to manage and track military resources (troops, vehicles), their assignments to operational zones, and historical movement logs. The system is designed with production-quality architecture to simulate the core features of military-grade planning platforms like Palantir's defense tools.

Its primary goal is to provide secure, auditable, and role-based command infrastructure for assigning, viewing, and analyzing resource deployment in real time.

## 🧱 Tech Stack
| Area                | Technology                                |
|---------------------|-------------------------------------------|
| Language            | Go (Golang)                               |
| Web Framework       | Gin                                       |
| ORM                 | GORM                                      |
| Database            | PostgreSQL (via Docker) or SQLite         |
| Auth                | JWT-based authentication with roles       |
| Deployment          | Docker & Docker Compose (optional)        |
| CI/CD               | GitHub Actions (planned)                  |

## 📁 Core Features
- ✅ **Zone Management**: Create and view operational zones
- ✅ **Resource Tracking**: Add troops/vehicles, assign to zones
- ✅ **Assignments**: Move resources between zones with logging
- ✅ **Movement Logs**: Audit who moved where, when, and why
- 🔜 **Role-Based Access Control**: Secure endpoints by user role
- 🔜 **CI/CD Pipeline**: Automate tests and builds via GitHub Actions
- 🔜 **Mission Planning**: Assign resources to active missions

## 🔐 User Roles (Planned Examples)
| Role    | Capabilities                              |
|---------|-------------------------------------------|
| Admin   | Full access to all endpoints              |
| Officer | Can assign resources and view zones/logs  |
| Viewer  | Read-only access to zones/resources/logs  |

## 📦 Current Status
- Project is structured in idiomatic Go with `internal/` modules
- Models, DB migrations, and handlers are implemented for:
  - Zones
  - Resources
  - Assignments
  - Logs
- API is testable locally with curl or Postman
