| Entity            | PK                            | SK                            |
| ----------------- | ----------------------------- | ----------------------------- |
| User              | `ACCOUNT#<Username>`          | `ACCOUNT#<Username>`          |
| Organization      | `ACCOUNT#<Organization>`      | `ACCOUNT#<Organization>`      |
| Verification code | `<VerificationCode>`          | `<VerificationCode>`          |
| Membership        | `ACCOUNT#<Organization>`      | `MEMBERSHIP#<Username>`       |
| Project           | `PROJECT#<Account>#<Project>` | `PROJECT#<Account>#<Project>` |
