Authentication models and controllers relied heavily on https://codevoweb.com/golang-and-gorm-user-registration-email-verification/

## DynamoDB docs & notes

https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#section-documentation

- You cannot update or delete items from a global secondary index

## Entity chart

| Entity / Type | PK                        | SK                                | Use case                  |
| ------------- | ------------------------- | --------------------------------- | ------------------------- |
| Account       | `ID#<UserAccountId>`      | `ID#<UserAccountId>`              | Fetch account details     |
| Project       | `PROJ#<ProjectId>`        | `PROJ#<ProjectId>`                | Fetch project details     |
| Member        | `PROJ#<ProjectId>`        | `TEAM#<UserAccountId>`            | Fetch members of project  |
| App           | `PROJ#<ProjectId>`        | `APP#<ProjectId>#<AppId>`         | Fetch apps of project     |
| Environment   | `APP#<ProjectId>#<AppId>` | `ENV#<ProjectId>#<AppId>#<EnvId>` | Fetch environments of app |
| OTP           | `OTP#<OTP>`               | `OTP#<OTP>`                       | Fetch OTP details         |

Should projects be renamed to stacks or apps?

## Global secondary index #1 (GSI1) // Sparse index

| Entity  | GSI1PK                 | GSI1SK                 | Use case               |
| ------- | ---------------------- | ---------------------- | ---------------------- |
| Account | `EMAIL#<AccountEmail>` | `EMAIL#<AccountEmail>` | Fetch account by email |

## Global secondary index #2 (GSI2)

| Entity      | GSI2PK                   | GSI2SK                    | Use case                    |
| ----------- | ------------------------ | ------------------------- | --------------------------- |
| Member      | `ID#<UserAccountId>`     | `PROJ#<ProjectId>`        | Fetch projects of account   |
| Project     | `PROJNAME#<ProjectName>` | `PROJNAME#<ProjectName>`  | Fetch project by name       |
| Environment | `PROJ#<ProjectId>`       | `ENV#<ProjectId>#<EnvId>` | Fetch apps of environment   |
| OTP         | `OTP#<AccountEmail>`     | `OTP#<AccountEmail>`      | Delete all OTPs for account |

## Access patterns

- Fetch account by account id: Account entity on PK
- Fetch account by email: Account entity on GSI1
- Fetch all accounts: Scan Account entity on GSI1
- Fetch all projects of which a user is a member: Member entity on GSI2
- Fetch all members of a project: Member entity on PK (SK >= "P")
- Fetch all apps of a project: App entity on PK (SK <= "P")
- Fetch all environments of an app: Environment entity on PK
- Fetch all environments of a project: Listed in the project entity
- Fetch all apps of an environment: Environment entity on GSI1

## Accounts & projects

- Accounts must be unique and projects must be unique
- A project is a collection of apps and team members
- The Depploy console will include the project name in the URL to signal to users which project they're viewing
- When a user creates a project, they become an admin member of the project
- Project admins can invite other team members to join their project
- A user can create and belong to multiple projects
