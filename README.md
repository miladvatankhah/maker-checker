# Message Approval Service

## Description
This project implements a simplified version of the maker-checker approval process using Go (Golang) with Domain-Driven Design (DDD) and Clean Architecture principles. The service allows users to send messages which must be validated by other users before being delivered to the recipient.

## Features
- User Registration
- Send Message with Approval/Reject Process
- Event Publishing to RabbitMQ on Approval
- Docker Support with Dockerfile and Docker Compose
- Clean and Modular Code using DDD and Clean Architecture

## Installation

### Prerequisites
- Docker
- Docker Compose
- Golang installed on your machine

### Steps
1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Build and run the application using Docker Compose:
    ```sh
    docker-compose up --build
    ```

## API Endpoints

### Register User
Register two users by providing only the `id`.

- **Endpoint:** `POST /api/v1/users`
- **Request Body:**
    ```json
    {
        "id": "{uuid}"
    }
    ```

### Create Message
Create a message to be sent from one user to another.

- **Endpoint:** `POST /api/v1/messages`
- **Request Body:**
    ```json
    {
        "content": "Hello World!",
        "sender_id": "{uuid of one of the created users}",
        "receiver_id": "{uuid of another created user}"
    }
    ```

### Approve Message
Approve the message to be sent to the recipient. This action publishes an event to a RabbitMQ exchange.

- **Endpoint:** `PATCH /api/v1/messages/approve/{id}`

### Reject Message
Reject the message so it will not be sent to the recipient.

- **Endpoint:** `PATCH /api/v1/messages/reject/{id}`

## Design Decisions
- **DDD and Clean Architecture:** Ensures code is modular, scalable, and maintainable.
- **Docker:** Allows easy setup and deployment.
- **Event-Driven:** Publishes events to RabbitMQ on message approval for further processing like email or notifications.
- **No Authentication/Authorization:** Simplified for demonstration purposes only.

## Future Enhancements
- Implement authentication and authorization.
- Develop a separate messaging bounded context (like `message_approval`) to consume events and handle notifications.

