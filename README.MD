# Real-Time Chat Application

Welcome to the **Real-Time Chat Application**! This project showcases a real-time chat system implemented using Go, WebSockets, and Redis. It is designed to demonstrate scalable and efficient real-time messaging and chat history management. 

## Overview

This application consists of three main components:

1. **WebSocket Service**: Handles real-time messaging between users. It also fetches chat history when a new user joins.
2. **History Storage Service**: Stores chat messages in Redis and provides an endpoint to fetch chat history.
3. **Redis**: Acts as a message broker and storage backend for chat messages.

## Features

- **Real-Time Messaging**: Enables instant communication between users using WebSockets.
- **Chat History Management**: Stores and retrieves chat messages in chronological order.
- **Dockerized Services**: Each component runs in its own Docker container for easy deployment and scalability.

## Architecture

1. **WebSocket Service**:
   - Handles WebSocket connections.
   - Publishes new messages to the Redis `chat_channel`.
   - Fetches chat history from the History Storage Service and sends it to clients on connection.

2. **History Storage Service**:
   - Subscribes to Redis `chat_channel`.
   - Saves incoming messages to the `chat_history` list in Redis.
   - Provides an HTTP endpoint to fetch chat history.

3. **Redis**:
   - Acts as a message broker with Pub/Sub functionality.
   - Stores chat history with no expiration.

## Getting Started

### Prerequisites

- Docker and Docker Compose installed on your machine.
- Basic understanding of Go, WebSockets, and Redis.

### Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/mahdi-eth/real-time-chat-app.git
   cd real-time-chat-app
   ```

2. **Build and Start Services**:
   ```bash
   docker-compose up --build
   ```
   This command will build Docker images for the WebSocket and History Storage services, and start all services defined in the docker-compose.yml file.

3. **Testing the Application**:
- Open your WebSocket client (such as WebSocket King Client or postman) and connect to ws://localhost:8080/ws.

- Send messages and see real-time communication in action.

- New connections will receive chat history automatically.

<hr>

Thank you for checking out the Real-Time Chat Application. We hope you find it useful and informative!

Happy coding!
