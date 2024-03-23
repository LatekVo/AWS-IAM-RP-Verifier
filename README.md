# AWS-IAM-RP-Interaction

### Description
This library loads, tests and verifies **AWS::IAM::Role Policy** json files.

---

### Execution
Launch `go run` in project's root

### Testing
Launch `go test` in project's root

---

### Task Reference
**AWS::IAM::Role Policy:** [Amazon AWS IAM Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html)

**Task description** [Google documents](https://docs.google.com/document/d/1zimLpe2F_0an4rAptqiIlCIky8NnUjt7/edit)

### Key notes

#### File structure
This codebase follows project structure recommendations described in the https://go.dev/ documentation.

#### Tests
Tests have no dedicated folder for the tests themselves besides the data folder. This is the standard recommended by go's documentation. 

This ensures high code coverage as there is a clear association being made between the tests and the code they're testing.