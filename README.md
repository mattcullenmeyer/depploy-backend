## Initialize

`go mod init github.com/mattcullenmeyer/depploy-backend`

TODO: Put these into a Makefile or bash script

`touch .git/hooks/pre-push`

```sh
#!/bin/sh

go test ./...
```

`chmod +x .git/hooks/pre-push`

```
sudo apt update
sudo apt install build-essential
sudo snap install golangci-lint
golangci-lint version
```

## Environment Variables

Set environment variables

```
export AWS_REGION=us-east-1
export DYNAMODB_TABLE_NAME=user
export DYNAMODB_ENDPOINT=http://localhost:8000
```

Output environment variables

```
echo $AWS_REGION
echo $DYNAMODB_TABLE_NAME
echo $DYNAMODB_ENDPOINT
```

## Run Locally

`go run main.go`  
OR  
`go run .` need to do this one if more than one file (package main) in the same directory (eg only running main.go will ignore routes.go)

http://localhost:8080/

## Run Dockerfile

`docker build -t depploy-image .`  
`docker image ls` list all images  
`docker run -p 8080:8080 --name depploy-container depploy-image`  
`docker ps` list all containers  
`docker rm depploy-container -f`
`docker image rm {image-id}`

## Run Docker Compose (this one!!!)

`docker compose build`  
`docker compose up`  
`docker compose down`

Ensure docker has permissions to access dynamodb volume  
https://stackoverflow.com/questions/45850688/unable-to-open-local-dynamodb-database-file-after-power-outage  
`sudo chmod 777 ./docker/dynamodb`

## DynamoDB Table

```
aws dynamodb create-table \
  --table-name user \
  --endpoint-url http://localhost:8000 \
  --attribute-definitions \
    AttributeName=Username,AttributeType=S \
  --key-schema \
    AttributeName=Username,KeyType=HASH \
  --region us-east-1 \
  --provisioned-throughput \
    ReadCapacityUnits=1,WriteCapacityUnits=1
```

```
aws dynamodb delete-table \
  --table-name user
  --endpoint-url http://localhost:8000
  --region us-east-1
```

```
aws dynamodb put-item \
  --table-name user \
  --endpoint-url http://localhost:8000 \
  --item '{"Username": {"S": "matt"}}' \
  --region us-east-1
```

```
aws dynamodb get-item \
  --table-name Depploy \
  --endpoint-url http://localhost:8000 \
  --key '{ "PK": {"S": "ACCOUNT#test20"}, "SK": {"S": "ACCOUNT#test20"}}' \
  --region us-east-1
```

## Lambda Deployment

The executable must be in the root of the zip file — not in a folder within the zip file.  
Use the -j flag to junk directory names, otherwise lambda won't work.

`GOOS=linux go build -o bin/main ./cmd/lambda && zip -j bin/main.zip bin/main`

## Testing

`go test ./...`
