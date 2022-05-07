# AWS profile Go

This is a pet project as part of learning Go lang and helping devs learning FE use the API for their portfolio.
The API is deployed to the AWS using Serverless framework it will create:

- CloudFormation stack to create all required resources
- Lambda functions exposed under single API gateway
- DynamoDB database to store the data
- S3 bucket to store the deployed resources
- Access policies for the resources to work properly

This should fall into the AWS Free Tier limitations.

The bins are included into the repo to ease the deployment process by mitigating the need of installing extra toolings.

## API definition

Look at [the OpenAPI docs](openapi.yaml) to discover available endpoints.

## Deploying the project

For deployment [Serverless framework](https://www.serverless.com/) is used. It will create all required resources
on the AWS. All resources are listed in the [serverless.yml](serverless.yml).

1. Clone the repo
   ```shell
   git clone https://github.com/gultyaev/aws-profile-go.git
   ```
2. Authenticate in AWS using [AWS CLI](https://aws.amazon.com/cli/)
3. Install npm packages
   ```shell
   npm i
   ```
4. Update `serverless.yml` values on lines 19 and 20 (after the semicolons).
   Those are the username & password you will be using to authenticate data changing requests.
   Don't commit them or expose in any way to the public, otherwise other people will be able to authenticate.
6. Run the deployment command
   ```shell
   npm run deploy
   ```

## Removing deployed resources from AWS

1. Make S3 bucket which name starts with `ImagesBucket` empty, otherwise the stack won't be removed.
2. Execute the command below. It will remove all resources created on AWS.
   ```shell
   sls remove
   ```
