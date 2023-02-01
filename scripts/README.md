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
- Create verification code: `go run ./scripts/auth/createVerificationCode/main.go -c=verification_code -a=account_id -u=username -e=email`
- Fetch verification code: `go run ./scripts/auth/fetchVerificationCode/main.go -c=verification_code`
- Update user verification: `go run ./scripts/auth/updateUserVerification/main.go -a=account_id`

### User scripts

- Fetch user by username: `go run ./scripts/user/fetchUserByUsername/main.go -u=username`
- Fetch user by account: `go run ./scripts/user/fetchUserByAccount/main.go -a=account_id`
- Delete user: `go run ./scripts/user/deleteUser/main.go -a=account_id`
- Fetch users: `go run ./scripts/user/fetchUsers/main.go -l=limit -k=next_key`
- Update account username: `go run ./scripts/user/updateAccountUsername/main.go -u=username -a=account_id`
- Update user superuser access: `go run ./scripts/user/updateUserSuperuser/main.go -a=account_id -s=true`
