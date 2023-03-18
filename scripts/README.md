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

- Create user: `go run ./scripts/auth/createUser/main.go -a=account_id -e=email`
- Create one-time password: `go run ./scripts/auth/createOtp/main.go -p=password -a=account_id -e=email`
- Fetch one-time password: `go run ./scripts/auth/fetchOtp/main.go -p=password`
- Fetch one-time passwords by email: `go run ./scripts/auth/fetchOtpsByEmail/main.go -e=email`
- Delete one-time password: `go run ./scripts/auth/deleteOtp/main.go -p=password`
- Update user verification: `go run ./scripts/auth/updateUserVerification/main.go -a=account_id -v=boolean`

### User scripts

- Fetch user: `go run ./scripts/user/fetchUserByAccount/main.go -a=account_id`
- Fetch user by email: `go run ./scripts/user/fetchUserByEmail/main.go -e=email`
- Delete user: `go run ./scripts/user/deleteUser/main.go -a=account_id`
- Fetch users: `go run ./scripts/user/fetchUsers/main.go -l=limit -k=next_key`
- Update user email: `go run ./scripts/user/updateUserEmail/main.go -a=account_id -e=email`
- Update user password: `go run ./scripts/user/updateUserPassword/main.go -a=account_id -p=password`
- Update user super admin access: `go run ./scripts/user/updateUserSuperAdmin/main.go -a=account_id -s=boolean`
