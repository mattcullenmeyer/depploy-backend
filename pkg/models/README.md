Authentication models and controllers relied heavily on https://codevoweb.com/golang-and-gorm-user-registration-email-verification/

## Entity chart

| Entity / Type | PK                           | SK                                     |
| ------------- | ---------------------------- | -------------------------------------- |
| User          | `ACCOUNT#<Username>`         | `ACCOUNT#<Username>`                   |
| Organization  | `ACCOUNT#<Organization>`     | `ACCOUNT#<Organization>`               |
| Membership    | `ACCOUNT#<Organization>`     | `MEMBER#<UserId>`                      |
| Project       | `ACCOUNT#<UserId>`           | `PROJECT#<UserId>#<Project>`           |
| Environment   | `PROJECT#<UserId>#<Project>` | `ENV#<Project>#<UserId>#<Environment>` |
| OTP           | `<OTP>`                      | `<OTP>`                                |

Should projects be renamed to stacks?

## Global secondary index #1 (GSI1) // Sparse index

| Entity     | GSI1PK        | GSI1SK        |
| ---------- | ------------- | ------------- |
| Account ID | `ID#<UserId>` | `ID#<UserId>` |

## Global secondary index #2 (GSI2) // Not a sparse index

| Entity     | GSI2PK          | GSI2SK          |
| ---------- | --------------- | --------------- |
| User Email | `EMAIL#<Email>` | `EMAIL#<Email>` |

## Access patterns

- Fetch account by account name (an account can be a user or an organization)
- Fetch account by id (use GSI1; need this for changing account names and logging)
- Fetch user account by email (use GSI2; need this for resending email verification and password reset)
- Fetch all accounts (users & orgs; scan on sparse GSI1)
- Fetch all organizations of which a user is a member
- Fetch all members of an organization
- Fetch all projects that belong to an account
- Fetch all environments for a project
- Change username for an account
- Change email for an account

## Accounts

- An account is a collection of projects with its own namespace (ie no 2 accounts can share the same name)
- The Depploy app console will include account name in the URL path to signal to user which account they're viewing
- An account can be a user or an organization
- A user represents a single person with one set of credentials
- An organization is a shared account, not a user (ie there is no login for an organization)
- A user must create an organization for other users to access shared projects
- A user can create and belong to multiple organizations
- A user account should only reference organization IDs and organizations should only reference user IDs (in case names are subsequently changed)
