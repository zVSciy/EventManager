# Project Writeup: FastAPI-Based Application

---

### Overview
This project is a web application built using FastAPI, providing user authentication and registration functionalities. The API is secured with JWT tokens and supports a MySQL database backend. It uses Docker to simplify deployment and integrates a frontend via a Svelte application.

- Requirements
    - Ubuntu 24.04 Server (recommended)
    - Docker Runtime (more security with Docker Rootless)
    - Source code of UserManagement service

- Ubuntu Server
    - Processors: 4
    - Memory (RAM): 4 GB
    - Storage: 40 GB

---

### Key Components
- Backend
    - Built using FastAPI.
    - Endpoints include:
        - `/register`: Registers a new user.
        - `/token`: Authenticates a user and returns a JWT token.

- Database
    - Backend uses SQLAlchemy as the ORM.
    - MySQL is used as the primary database, managed via Docker.
    - A `User` model is defined in `/src/model/models.py`.

- Authentication
    - JWT tokens are generated using the python-jose library.
    - Passwords are hashed using Passlib.

- Environment Management
    - Configuration variables (e.g., database credentials) are stored in a `.env` file and loaded using python-dotenv.

- Docker
    - Services are containerized:
        - MySQL database.
        - Backend running the FastAPI app.
        - Frontend running Svelte.
    - Managed via docker-compose.

---

### Installing Docker
* Uninstall all conflicting packages:
```bash
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done
```
* Update repository:
```bash
sudo apt-get update
```
* Add Docker's official gpg-key:
```bash
sudo apt-get install ca-certificates curl && \
sudo install -m 0755 -d /etc/apt/keyrings && \
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc && \
sudo chmod a+r /etc/apt/keyrings/docker.asc
```
* Add the repository to apt-sources:
```bash
echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
    $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```
* Update again:
```bash
sudo apt-get update
```
* Install Docker packages:
```bash
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

---

### Setup rootless Docker
* Install uidmap:
```bash
sudo apt-get install -y uidmap
``` 
* Shut down Docker deamon:
```bash
sudo systemctl disable --now docker.service docker.socket
```
* Install Dockerd-Rootless-Setup:
```bash
dockerd-rootless-setuptool.sh install
```
* Define variable DOCKER_HOST:
```bash
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock
```
* Verify that Docker is running rootless:
```bash
docker run hello-world
```

---

### Setup Instructions
- Prerequisites
    - Install Docker and Docker Compose.
    - Install Python 3.10 or higher.
    - Ensure a `.env` file is present with the following variables:
        ```sh
        MYSQL_USER=root
        MYSQL_PASSWORD=yourpassword
        MYSQL_DATABASE=yourdatabase
        MYSQL_DATABASE_HOST=db
        MYSQL_DATABASE_PORT=3306
        SECRET_KEY=yourjwtsecretkey
        ALGORITHM=HS256
        ALLOWED_HOSTS=*
        ```

- Database Configuration
    - The MySQL database is managed via Docker Compose.
    - The database connection string is constructed in `/src/model/database.py` using SQLAlchemy.

- Backend Configuration
    - Backend code is located in `/src`.
    - Key files:
        - `main.py`: Entry point for the FastAPI application.
        - `routes.py`: Defines the /register and /token endpoints.
        - `database.py`: Contains the database engine and get_db() dependency.
        - `models.py`: Defines the User table schema.
        - `schema.py`: Defines Pydantic models for request validation.
        - `security/`: Contains password hashing and JWT token generation logic.

- Frontend Configuration
    - Located in the `/frontend` directory.
    - Built using Svelte.
    - Communicates with the backend over HTTP.

- Running Locally
    - Clone the repository.
    - Navigate to the project directory.
    - Build and start services:
        ```bash
        docker-compose up --build
        ```
    - Access the services:
        - Backend: https://localhost:8000
        - Frontend: http://localhost:3000
    - Verify database connectivity by inspecting the logs.

- Running Tests
    - Tests are located in `/tests`.
    - Use unittest to run test cases.
    - Tests include:
        - User registration success and failure cases.
        - Login with valid and invalid credentials.

---

### How It Works
- Database: `/src/model/database.py`:
    - Loads database connection info from the `.env` file.
    - Sets up a SQLAlchemy engine and session factory.

- Models: `/src/model/models.py`:
            - Defines the `User` table with attributes such as `email`, `hashed_password`, `first_name`, `last_name`, and `role`.

- Routes: `/src/routes.py`:
    - `/register`:
        - Checks if a user with the given email already exists.
        - Hashes the password and saves the new user in the database.
    - `/token`:
        - Verifies the user’s credentials.
        - Generates a JWT token using the `create_access_token` function.

- Authentication: `/src/security/authentication.py`:
    - Generates JWT tokens with a 15-minute expiration.#
    - Uses the `SECRET_KEY` and `ALGORITHM` from the `.env` file.

- Hashing: `/src/security/hashing.py`:
    - Passwords are hashed and verified using Passlib with the bcrypt algorithm.

- Frontend
    - Connects to the backend using fetch APIs to handle user registration and login.

- Docker
    - `/docker-compose.yml`:
        - Defines three services: db, backend, and frontend.
        - Uses environment variables to configure the database and backend.
    - `/docker/FastAPI/Dockerfile`:
        - Installs Python dependencies.
        - Copies the backend code to the container.

---

### Deployment
- SSL Support:
    - Backend is configured to use SSL certificates located in `/ssl`.
    - Modify `/docker-compose.yml` to provide valid `domain.key` and `domain.cert` files.

- Production Considerations:
    - Use a production-grade server (e.g., Gunicorn) for the backend.
    - Use a reverse proxy (e.g., NGINX) to serve the frontend and backend.

---

### Troubleshooting

- Database Connection Issues:
    - Ensure `.env` file values are correct.
    - Verify that the db service is running.
    - Ensure that the host os has enough storage left.

- Authentication Errors:
    - Check the `SECRET_KEY` and `ALGORITHM` values in the `.env` file.
    - Ensure passwords are hashed correctly.

- Testing Failures:
    - Verify the test database setup in `test_api.py`.
    - Ensure all dependencies are installed.