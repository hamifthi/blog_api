# Personal Blog API

## Overview

The Personal Blog API is a RESTful service designed using the Domain-Driven Design (DDD) approach. The application provides functionalities to manage blog posts, authors, categories, and tags. This project is structured to ensure scalability, maintainability, and clear separation of concerns.

## Project Structure

The project is divided into several packages, each representing a distinct layer of the application. Below is an overview of each folder:


### Detailed Directory Breakdown

- **`cmd/`**: Contains the main application entry point (`main.go`). This is where the application bootstraps and starts running by initializing routes, services, and dependencies.

- **`conf/`**: Contains the file for handling the configuration of project.

- **`deployments/`**: Holds deployment-related scripts and configurations, such as Dockerfiles, Kubernetes manifests, and CI/CD pipeline scripts to deploy the application.

- **`entity/`**: Contains domain entities that define the core business objects, such as `Blog`, `Author`, `Category`, and `Tag`. These entities are the heart of the domain layer, focusing on business rules.

- **`log/`**: Manages the logging setup, using libraries like `zap` or `logrus`. Proper logging helps in debugging and monitoring application performance.

- **`repository/`**: Implements the Repository pattern to abstract and encapsulate data access logic. Each repository handles the CRUD operations for the associated entities (e.g., `BlogRepository`, `AuthorRepository`).

- **`resources/`**: Stores configuration files, or other resources that the application might serve or use.

- **`router/`**: Defines the API routes and sets up middleware. It configures the routing logic using frameworks like Echo, Gin, or Chi, handling the mapping of HTTP requests to the appropriate controller methods.

- **`service/`**: Implements business logic and orchestrates the interaction between the domain layer and repositories. Services are responsible for enforcing business rules and coordinating use cases.

- **`storage/`**: Manage the db connection and migration of database when the project bootstrap, etc.

## Prerequisites

- **Go 1.21+**
- **Docker** (for containerization)
- **PostgreSQL** or any supported database
- **Make** sure all dependencies are installed:


### Summary

This README serves as a comprehensive guide for explaining structure of the Personal Blog API project. It clearly explains the project's structure, and outlines the architecture's adherence to Domain-Driven Design principles.

