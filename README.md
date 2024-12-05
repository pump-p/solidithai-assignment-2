# Solidithai Assignment

This project is a full-stack application that includes a **Go** backend and a **React** frontend. It implements user authentication, registration, basic CRUD operations, and real-time WebSocket communication with Elasticsearch integration for log storage and querying.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Tech Stack](#tech-stack)
3. [Installation Guide](#installation-guide)
   - [Backend Setup](#setting-up-the-backend)
   - [Frontend Setup](#setting-up-the-frontend)
4. [Features](#features)
5. [Design Decisions](#design-decisions)
6. [Postman API Documentation](#postman-api-documentation)

## Project Overview

This project is designed to meet the requirements of both assignments by implementing the specified features.

This application consists of two parts:

1. **Backend**: A Go application using the Gin framework to handle user authentication, registration, CRUD operations, and WebSocket-based real-time communication. Logs from WebSocket messages are stored in Elasticsearch for efficient querying.
2. **Frontend**: A React application to interact with the backend, allowing users to sign up, log in, manage their accounts, and view logs or interact with the WebSocket client.

The backend exposes a REST API for authentication and user management, while the frontend provides a simple UI to interact with the API and WebSocket features.

## Tech Stack

- **Backend**: Go, Gin, JWT (for authentication), WebSocket
- **Frontend**: React, axios, TailwindCSS
- **Database**: MariaDB (MySQL)
- **Logging**: Elasticsearch for WebSocket log storage and querying

## Installation Guide

### Prerequisites

Before running the application, ensure that the following tools are installed on your system:

- **Go**: [Download Go](https://golang.org/dl/)
- **Node.js**: [Download Node.js](https://nodejs.org/) (for frontend development)
- **MariaDB**: Install MariaDB for local database setup.
- **Elasticsearch**: Install Elasticsearch for log storage.

### Setting Up the Backend

1. **Clone the repository**:

   ```bash
   git clone https://github.com/pump-p/solidithai-assignment-2.git
   cd solidithai-assignment-2/backend
   ```

2. **Install Go dependencies**:

   Run the following command in the backend folder to install the required Go modules:

   ```bash
   go mod tidy
   ```

3. **Set up the database**:

   Ensure MariaDB is running. Create a database called `solidithai_assignment_2` and configure the connection string in `backend/config/config.go`.

   Example configuration in `config.go`:
   ```go
   db, err := gorm.Open(mysql.Open("user:password@tcp(localhost:3306)/solidithai_assignment_2"), &gorm.Config{})
   ```

4. **Set up Elasticsearch**:

   Run Elasticsearch locally or using Docker:

   ```bash
   docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.17.0
   ```

   Ensure the Elasticsearch URL is configured in `.env` or `config.go`:
   ```env
   ELASTICSEARCH_URL=http://localhost:9200
   ```

5. **Run the backend server**:

   Start the Go application:

   ```bash
   go run main.go
   ```

   The backend server should now be running on `http://localhost:8080`.

### Setting Up the Frontend

1. **Navigate to the frontend folder**:

   ```bash
   cd ../frontend
   ```

2. **Install frontend dependencies**:

   Run the following command to install the React application dependencies:

   ```bash
   npm install
   ```

3. **Run the frontend**:

   Start the React application:

   ```bash
   npm start
   ```

   The frontend should now be running on `http://localhost:3000`.

## Running the Application

1. **Start the backend**: Follow the steps above to run the Go backend on `http://localhost:8080`.
2. **Start the frontend**: Run `npm start` to launch the React frontend on `http://localhost:3000`.

Now, you can interact with the application in your browser, with the frontend making requests to the backend.

## Features

### Backend

1. **User Authentication**:
   - JWT-based login and signup functionality.
   - Passwords are securely hashed using bcrypt.
2. **CRUD Operations**:
   - RESTful endpoints for managing users.
   - Passwords are hashed during updates if provided.
3. **WebSocket Integration**:
   - Real-time communication using WebSockets.
   - WebSocket messages are logged to Elasticsearch.
4. **Elasticsearch Integration**:
   - Logs from WebSocket communication are stored in Elasticsearch.
   - A REST API endpoint is available for querying these logs.

### Frontend

1. **Authentication**:
   - Login and signup pages for user authentication.
2. **User Management**:
   - Pages for viewing, creating, updating, and deleting users.
3. **Logs Viewer**:
   - A page to query and display WebSocket logs from Elasticsearch.
4. **WebSocket Client**:
   - A client interface for sending and receiving real-time messages.

## Design Decisions

### Backend

- **Go with Gin Framework**: Go was chosen for its performance, simplicity, and scalability. The Gin framework was selected for its lightweight and fast routing capabilities.
- **Structure**:
  - **Controllers** handle HTTP requests and responses.
  - **Models** define the structure of the data.
  - **Services** encapsulate business logic.
  - **Repositories** handle database interactions.
- **Elasticsearch**: Used for storing and querying logs due to its high-performance search capabilities.

### Frontend

- **React**: React’s component-based architecture simplifies UI management and scaling.
- **TailwindCSS**: TailwindCSS was chosen for its utility-first design, allowing rapid development with consistent styles.

### Database

- **MariaDB** was chosen over MySQL for its open-source nature while maintaining MySQL compatibility.

## Postman API Documentation

You can find the detailed API documentation for this project on Postman here:

[Postman API Documentation](https://documenter.getpostman.com/view/26586964/2sAYBaAVdd)


The following image demonstrates that WebSocket functionality is successfully working in Postman:

<img width="720" alt="ws" src="https://github.com/user-attachments/assets/63607dde-c085-44df-b930-af9dc89097a2">

