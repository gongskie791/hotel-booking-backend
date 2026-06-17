# Hotel Booking System — Backend

A hotel booking REST API backend built with Go, containerized with Docker, and deployed on AWS EC2 via GitHub Actions.

## Tech Stack

- **Backend** — Go
- **Container** — Docker
- **CI/CD** — GitHub Actions
- **Cloud** — AWS (EC2, ECR, RDS, ALB, ACM, Route 53, Secrets Manager, CloudWatch, SSM)

## CI/CD Pipeline

Triggered on every push to `master`:

1. Vet, test, and build the Go app
2. Build and push Docker image to ECR
3. Deploy to EC2 via SSM Session Manager (no SSH required)
4. EC2 pulls the latest image, stops the old container, and starts the new one
5. Poll `GET /api/health` for up to 60s to confirm the service is up

## Required GitHub Secrets

| Secret                  | Description                   |
| ----------------------- | ----------------------------- |
| `AWS_ACCESS_KEY_ID`     | IAM access key                |
| `AWS_SECRET_ACCESS_KEY` | IAM secret key                |
| `AWS_REGION`            | AWS region                    |
| `ECR_REGISTRY`          | ECR registry URL              |
| `ECR_REPOSITORY`        | ECR repository name           |
| `INSTANCE_NAME`         | EC2 instance Name tag for SSM |
