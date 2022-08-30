## Run Locally

`$ go run main.go`  
http://localhost:8080/

## Run Docker

`$ docker compose up -d`  
`$ docker compose down`

Create a table

```
$ aws dynamodb \
  --endpoint-url http://localhost:8000 create-table \
  --table-name user \
  --attribute-definitions AttributeName=name,AttributeType=S \
  --key-schema AttributeName=name,KeyType=HASH \
  --region us-east-1 \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
```

`$ aws dynamodb put-item --endpoint-url http://localhost:8000 --table-name user --item '{"name": {"S": "matt"}}' --region us-east-1`  
`$ aws dynamodb get-item --endpoint-url http://localhost:8000 --table-name user --key '{"name": {"S": "matt"}}' --region us-east-1`

## Initialize

`$ go mod init github.com/mattcullenmeyer/depploy-backend`

https://www.youtube.com/playlist?list=PL5dTjWUk_cPbe1PZvwMIufb2fdwzVpM3O
