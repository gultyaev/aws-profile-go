# AWS profile Go

This is a pet project as part of learning Go lang and helping devs learning FE use the API for their portfolio.
The API is deployed to the AWS using Serverless framework it will create:

- CloudFormation stack to create all required resources
- Lambda functions exposed under single API gateway
- DynamoDB database to store the data
- S3 bucket to store the deployed resources

This should fall into the AWS Free Tier limitations.

The bins are included into the repo to ease the deployment process by mitigating the need of installing extra toolings.

## Deploying the project

```shell
sls deploy
```

## Removing project from deployments

```shell
sls remove
```
