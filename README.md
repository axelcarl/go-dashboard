# Go Payment Dashboard

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

Example Dashboard written in Go with best practices in mind. Still under construction!

Visit the dashboard here: [Go Payment Dashboard](https://go-dashboard.fly.dev)

## Getting Started

1. Make sure you have `docker` installed.

2. Create a `.env` in the `root` and `frontend` directories. These should mimic the `.env.example` and `.env.production` respectively.

3. Run the project with live reloading using:
   ```bash
   docker compose up
   ```

- Run the tests using

  ```bash
  go test ./... -v
  ```

- Run the integration tests using
  ```bash
  go test ./internal/database -v
  ```
