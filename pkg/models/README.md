Authentication models and controllers relied heavily on https://codevoweb.com/golang-and-gorm-user-registration-email-verification/

## Entity chart

| Entity       | PK                            | SK                                      |
| ------------ | ----------------------------- | --------------------------------------- |
| Account      | `ACCOUNT#<Account>`           | `ACCOUNT#<Account>`                     |
| Organization | `ACCOUNT#<Organization>`      | `ACCOUNT#<Organization>`                |
| Membership   | `ACCOUNT#<Organization>`      | `MEMBERSHIP#<Username>`                 |
| Project      | `ACCOUNT#<Account>`           | `PROJECT#<Account>#<Project>`           |
| Environment  | `PROJECT#<Account>#<Project>` | `ENV#<Project>#<Account>#<Environment>` |
| OTP          | `<OTP>`                       | `<OTP>`                                 |

Should projects be renamed to stacks?

## Global secondary index #1 (GSI1)

| Entity  | GSI1PK              | GSI1SK              |
| ------- | ------------------- | ------------------- |
| Account | `ACCOUNT#<Account>` | `ACCOUNT#<Account>` |

## Access patterns

- Fetch account (an account can be a user or an organization)
- Fetch all accounts (scan on sparse global secondary index)
- Fetch all organizations an account belongs to
- Fetch all projects that belong to an account
- Fetch all environments for project
