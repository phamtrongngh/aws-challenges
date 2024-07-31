# CodeBuild Web App
Difficulty: ★ ★ ★ ☆ ☆

## Overview
In this lab, you will learn how to use AWS CodeBuild to build a simple NextJS application.

## Scenario
You are a DevOps engineer at a startup company. You are tasked with automating the build process for a NextJS web application. You have decided to use AWS CodeBuild to build the app with Docker and push the image to AWS Elastic Container Registry (ECR).
Because this lab focuses on CodeBuild, we won't cover the deployment part, e.g., deploying the image to EC2, ECS, or EKS.However, you can use the Docker image built and pushed to Docker Hub in the this lab anytime you want.

## Steps

### 0. Prerequisites
You need to have the following tools installed on your machine:
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html), should be configured with the appropriate permissions, e.g., `AdministratorAccess`.

### 1. Provision the resources
Use the provided CloudFormation template file `cfn-templates/codebuild-web-app.yml` to provision the resources for this lab. 
Run the following command to create the stack:
```bash
aws cloudformation create-stack --stack-name codebuild-web-app --template-body file://cfn-templates/codebuild-web-app.yaml --capabilities CAPABILITY_NAMED_IAM
```
Wait for the stack to be created successfully.

After the stack is created, you will see the following resources in your AWS account:
- CodeBuild project
- ECR repository

### 2. Start the build
Navigate to the CodeBuild console and start the build process. You can use the AWS CLI to start the build process:
```bash
aws codebuild start-build --project-name codebuild-web-app
```