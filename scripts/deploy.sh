#!/bin/bash
set -e

aws ecr get-login-password --region "$AWS_REGION" | docker login --username AWS --password-stdin "$ECR_REGISTRY"
TAG="${IMAGE_TAG:-latest}"
docker pull "$ECR_REGISTRY/$ECR_REPOSITORY:$TAG"
docker stop hotel-booking-backend || true
docker rm hotel-booking-backend || true

SECRET=$(aws secretsmanager get-secret-value \
  --secret-id "hotel-booking/db-credentials" \
  --region "$AWS_REGION" \
  --query SecretString \
  --output text)

DB_HOST=$(echo "$SECRET" | jq -r '.DB_HOST')
DB_USER=$(echo "$SECRET" | jq -r '.DB_USER')
DB_PASSWORD=$(echo "$SECRET" | jq -r '.DB_PASSWORD')
DB_NAME=$(echo "$SECRET" | jq -r '.DB_NAME')

docker run -d --restart always --name hotel-booking-backend -p 8080:8080 \
  --log-driver awslogs \
  --log-opt awslogs-region="$AWS_REGION" \
  --log-opt awslogs-group=/hotel-booking/backend \
  --log-opt awslogs-stream=ec2 \
  -e DB_HOST="$DB_HOST" \
  -e DB_USER="$DB_USER" \
  -e DB_PASSWORD="$DB_PASSWORD" \
  -e DB_NAME="$DB_NAME" \
  "$ECR_REGISTRY/$ECR_REPOSITORY:$TAG"

for i in $(seq 1 12); do
  HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/health || echo "000")
  if [ "$HTTP_STATUS" = "200" ]; then
    echo "Health check passed (attempt $i)."
    exit 0
  fi
  echo "Attempt $i: got HTTP $HTTP_STATUS, retrying in 5s..."
  sleep 5
done

echo "Health check failed after 60s."
docker logs hotel-booking-backend
exit 1
