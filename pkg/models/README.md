| Entity            | PK                            | SK                            |
| ----------------- | ----------------------------- | ----------------------------- |
| User              | `ACCOUNT#<Username>`          | `ACCOUNT#<Username>`          |
| Organization      | `ACCOUNT#<Organization>`      | `ACCOUNT#<Organization>`      |
| Verification code | `<VerificationCode>`          | `<VerificationCode>`          |
| Membership        | `ACCOUNT#<Organization>`      | `MEMBERSHIP#<Username>`       |
| Project           | `PROJECT#<Account>#<Project>` | `PROJECT#<Account>#<Project>` |

Authentication models and controllers relied heavily on https://codevoweb.com/golang-and-gorm-user-registration-email-verification/
