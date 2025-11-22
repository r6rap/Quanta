# Quanta

Quanta is a lightweight platform for project management and issue tracking. It helps teams plan, organize, and monitor work through structured project spaces, clear workflows, and consistent collaboration tools.

## Core Features

### Project Management
- Create and manage projects with defined ownership.
- Add and manage project members with specific roles.

### Issue Tracking
- Track work items using a structured workflow (Todo, In Progress, Review, Done).
- Assign issues, set priorities, and apply labels.

### Comments & Activity Logs
- Discuss work directly on issues.
- Maintain an immutable activity history for transparency and auditing.

### Role-Based Permissions
- Global and project-level role enforcement to control access and actions.
- Clear separation between admin, project owner, member, and read-only roles.

### Search & Filtering
- Query issues by status, labels, assignee, and priority.
- Optimized indexing for efficient retrieval.

### Caching (Redis)
- Cache frequently accessed data to reduce load on the primary database.
- Support for token blacklisting and rate limiting where needed.

### Authentication
- JWT-based access tokens and rotating refresh tokens.
- Secure session lifecycle and token invalidation support.

## Technical Goals

Quanta is designed as a practical environment for reinforcing backend engineering fundamentals:

- Relational database modeling with proper indexing.
- Clean service-layer architecture and modular code organization.
- Authentication and authorization best practices.
- Redis caching strategies for performance and scalability.
- Domain modeling and workflow implementation.
- Secure, maintainable, and scalable backend design suitable for real-world use.
