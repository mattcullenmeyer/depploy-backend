Authentication models and controllers relied heavily on https://codevoweb.com/golang-and-gorm-user-registration-email-verification/

## Entity chart

| Entity / Type | PK                              | SK                                        |
| ------------- | ------------------------------- | ----------------------------------------- |
| Account ID    | `ID#<UserAccountId>`            | `ID#<UserAccountId>`                      |
| Organization  | `ID#<OrgAccountId>`             | `ID#<OrgAccountId>`                       |
| Membership    | `ID#<OrgAccountId>`             | `MEMBER#<UserAccountId>`                  |
| Project       | `ID#<AccountId>`                | `PROJECT#<AccountId>#<Project>`           |
| Environment   | `PROJECT#<AccountId>#<Project>` | `ENV#<Project>#<AccountId>#<Environment>` |
| OTP           | `<OTP>`                         | `<OTP>`                                   |

Should projects be renamed to stacks?

## Global secondary index #1 (GSI1) // Sparse index

| Entity       | GSI1PK                  | GSI1SK                  |
| ------------ | ----------------------- | ----------------------- |
| Account Name | `ACCOUNT#<AccountName>` | `ACCOUNT#<AccountName>` |

## Access patterns

- Fetch account by account id (an account can be a user or an organization)
- Fetch account by account name (use GSI1; also referred to as username)
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
- An organization is a shared account (ie there is no login for an organization)
- A user must create an organization for other users to access shared projects
- A user can create and belong to multiple organizations
- A user account should only reference organization IDs and organizations should only reference user IDs in case account names are changed
