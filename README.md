# Hotel Booking System — Backend

A hotel booking REST API backend built with Go, containerized with Docker, and deployed on AWS EC2 via GitHub Actions.

## Tech Stack

- **Backend** — Go
- **Container** — Docker
- **CI/CD** — GitHub Actions
- **Cloud** — AWS (EC2, ECR, Secrets Manager, CloudWatch)

## CI/CD Pipeline

Triggered on every push to `master`:

1. Vet, test, and build the Go app
2. Build and push Docker image to ECR
3. SSH into EC2, pull the latest image, and run the container
4. Poll `GET /api/health` for up to 60s to confirm the service is up

## Required GitHub Secrets

| Secret                  | Description               |
| ----------------------- | ------------------------- |
| `AWS_ACCESS_KEY_ID`     | IAM access key            |
| `AWS_SECRET_ACCESS_KEY` | IAM secret key            |
| `AWS_REGION`            | AWS region                |
| `ECR_REGISTRY`          | ECR registry URL          |
| `ECR_REPOSITORY`        | ECR repository name       |
| `EC2_HOST`              | EC2 public IP or hostname |
| `EC2_USER`              | SSH username              |
| `EC2_SSH_KEY`           | SSH private key           |
