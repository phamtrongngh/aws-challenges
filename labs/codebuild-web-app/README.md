# CodeBuild Web App
Difficulty: ★ ★ ★ ☆ ☆

## Overview
In this lab, you will learn how to use AWS CodeBuild to build a simple NextJS application.

## Scenario
You are a teach lead at a startup company that is building a new web application. You don't want your teams to spend time setting up and managing on-premises build servers like Jenkins. Instead, you want to use a fully managed build service that can scale with your team's needs. You have decided to use AWS CodeBuild to automate the build process. For the proof of concept, you will create a simple NextJS web application and configure CodeBuild to build it.

## Steps

### 0. Prerequisites
You need to have the following tools installed on your machine:
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html), should be configured with the appropriate permissions, e.g., `AdministratorAccess`.

### 1. Create a buildspec.yml file
Create a file named `buildspec.yml` in the root of your project directory. This file contains the build commands that CodeBuild will run when building your project. The following is an example of a `buildspec.yml` file for a Node.js project:
```yaml

