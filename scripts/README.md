## Table of contents

- [Table of contents](#table-of-contents)
- [Environment variables](#environment-variables)
  - [Local](#local)
  - [Prod](#prod)
- [Scripts](#scripts)
  - [Auth scripts](#auth-scripts)
  - [User scripts](#user-scripts)

## Environment variables

### Local

```
export AWS_REGION=us-east-1
export DYNAMODB_TABLE_NAME=Depploy
export DYNAMODB_ENDPOINT=http://localhost:8000
```

### Prod

```
export AWS_REGION=us-east-1
export DYNAMODB_TABLE_NAME=depploy-users-prod
export DYNAMODB_ENDPOINT=dynamodb.us-east-1.amazonaws.com
```

## Scripts

### Auth scripts

- Create user: `go run ./scripts/auth/createUser/main.go -u=username -e=email -p=password`

- Create verification code: `go run ./scripts/auth/createVerificationCode/main.go -c=verification_code -u=username -e=email`

- Fetch verification code: `go run ./scripts/auth/fetchVerificationCode/main.go -c=verification_code`

- Update user verification: `go run ./scripts/auth/updateUserVerification/main.go -u=username`

### User scripts

- Fetch user: `go run ./scripts/user/fetchUser/main.go -u=username`

- Fetch users: `go run ./scripts/user/fetchUsers/main.go -l=limit -k=next_key`
