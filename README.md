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
export DYNAMODB_TABLE_NAME=Depploy
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
`docker compose --env-file .env up` or `make docker-up`  
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

```
aws dynamodb get-item \
  --table-name Depploy \
  --endpoint-url http://localhost:8000 \
  --key '{ "PK": {"S": "VU2LHOHWD7LXRBLDDNCUOVYACPAIFMGP"}, "SK": {"S": "VU2LHOHWD7LXRBLDDNCUOVYACPAIFMGP"}}' \
  --region us-east-1
```

## Create Super Admin User

See scripts (updateUserSuperAdmin)

## Lambda deployment

The executable must be in the root of the zip file — not in a folder within the zip file.  
Use the -j flag to junk directory names, otherwise lambda won't work.

`GOOS=linux go build -o bin/main ./cmd/lambda && zip -j bin/main.zip bin/main`

## Testing

`go test ./...`

## Postman collection

[Confluence link](https://mattcullenmeyer.atlassian.net/wiki/spaces/~701217c77864a7a5f4f69b9b38d1f152ff014/pages/196609/Depploy+Wiki)

## Helpful Go commands

https://go.dev/ref/mod#go-mod-tidy  
`go mod tidy -v`

## Google OAuth

Google OAuth is managed from the [Google API Console](https://console.cloud.google.com/apis/credentials/consent?project=depploy-prod).
There are "projects" for each environment (eg Depploy - PROD).
Toggle through the different projects / environments using the dropdown in the top nav.
Use the "Credentials" and "OAuth consent screen" menus in the side nav to edit the settings.
Credentials are saved in an AWS System Manager Parameter Store.

If you want to change the user support email, you'll first need to sign up for a Google account using that email.
An existing account owner must then invite the new account using the following steps.

1. Navigate to IAM & Admin >> IAM.
2. Click the "Grant Access" button.
3. Add a new "principal", using the new email account address
4. Assign the role of Owner
5. Have the new email account accept the invite

You should now be able to change the user support email to the new email account.
