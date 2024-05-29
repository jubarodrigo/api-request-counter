# api-request-counter
Simple API for request counting and local storage

## Features

- [x] Rest API
- [x] Unit Tests

## Runing Local

---

1. Install the following dependencies:
- Choose one of the following free container management tool:
    - [Rancher Desktop](https://rancherdesktop.io/)
    - [Podman](https://podman.io/)
    - [Orbstack](https://orbstack.dev/)
- [Docker Compose](https://docs.docker.com/compose/install/)


2. Run command on terminal to up service:
 ```bash
   $ docker-compose up
   ```
3. Routes:
- Host: http://localhost:8090
- GET Requests Counter: http://localhost:8090/api/v1/counter
---

## Runing Local

To run all tests, run the following command:
 ```bash
   $ go test ./...
   ```