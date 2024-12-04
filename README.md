
# Solidithai Assignment

This project is a full-stack application that includes a **Go** backend and a **React** frontend. It implements user authentication, registration, and basic CRUD operations.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Tech Stack](#tech-stack)
3. [Installation Guide](#installation-guide)
4. [Design Decisions](#design-decisions)

## Project Overview

This application consists of two parts:
1. **Backend**: A Go application using the Gin framework to handle user authentication, registration, and CRUD operations.
2. **Frontend**: A React application to interact with the backend, allowing users to sign up, log in, and manage their accounts.

The backend exposes a REST API for authentication and user management, while the frontend provides a simple UI to interact with the API.

## Tech Stack

- **Backend**: Go, Gin, JWT (for authentication)
- **Frontend**: React, axios
- **Database**: MariaDB (MySQL)

## Installation Guide

### Prerequisites

Before running the application, ensure that the following tools are installed on your system:

- **Go**: [Download Go](https://golang.org/dl/)
- **Node.js**: [Download Node.js](https://nodejs.org/) (for frontend development)
- **MariaDB**: Install MariaDB for local database setup.

### Setting Up the Backend

1. **Clone the repository**:

   ```bash
   git clone https://github.com/pump-p/solidithai-assignment-2.git
   cd solidithai-assignment-2/backend
   ```

2. **Install Go dependencies**:

   Make sure you have Go installed, and run the following command in the backend folder to install the required Go modules:

   ```bash
   go mod tidy
   ```

3. **Set up the database**:

   Ensure you have MySQL running. Create a database called `solidithai_assignment_2` and configure the connection string in `backend/config/config.go`.

   Example configuration in `config.go`:
   ```go
   db, err := gorm.Open(mysql.Open("user:password@tcp(localhost:3306)/solidithai_assignment_2"), &gorm.Config{})
   ```

4. **Run the backend server**:

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

## Design Decisions

### Backend

- **Go with Gin Framework**: I chose Go for the backend because of its performance, simplicity, and scalability. The Gin framework was selected due to its lightweight and fast routing capabilities, which are ideal for REST APIs.
- **Structure**:
  - **Controllers** handle the HTTP requests and responses.
  - **Models** define the structure of the data.
  - **Services** contain the business logic of the application.

### Frontend

- **React**: React was chosen for the frontend because of its component-based architecture, which makes it easy to manage UI components and scale the app.

### Database

- **MariaDB** was chosen over MySQL because it is a fully open-source, while maintaining compatibility with MySQL applications.
