# Database Migration Basic On AWS

## Overview
In this lab, you will learn how to migrate a PostgreSQL database from an on-premises environment to AWS RDS using AWS Database Migration Service (DMS).

## Scenario
You are working as a DevOps Engineer for a multinational company using an on-premises PostgreSQL database to store the employee's information. The company is planning to migrate the PostgreSQL database to AWS RDS to reduce the operational overhead of managing the database. You have been tasked to migrate that database to AWS RDS using AWS Database Migration Service (DMS).

## Steps
### 0. Prerequisites
You need to have the following tools installed on your machine:
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html), should be configured with the appropriate permissions, e.g., `AdministratorAccess`.

### 1. Provision the on-premises system
We will use a PostgreSQL database running on an EC2 instance to simulate an on-premises environment. You can use the provided CloudFormation template `onprem.yaml` to provision the resources:
```bash
aws cloudformation create-stack --stack-name on-prem-system --template-body file://onprem.yaml
```
Wait for a few minutes for the CloudFormation stack to be created.

### 2. Review the on-premises PostgreSQL database
For simplicity, we will use the `psql` command-line tool to interact with the `employees` database. Alternatively, you can use any PostgreSQL GUI tool (e.g., pgAdmin, DBeaver, TablePlus).

Use the following command to connect to the `employees` database:
```bash
docker exec -ti postgres psql -h localhost -U postgres -d employees
```

After connecting to the database, you can use commands to interact with the database.

Use `\dt` to list all tables in the `employees` database, and you should see two tables: `country` and `engineer`:
```sql
          List of relations
 Schema |   Name   | Type  |  Owner   
--------+----------+-------+----------
 public | country  | table | postgres
 public | engineer | table | postgres
```
Get number of rows in the `engineer` table:
```sql
SELECT COUNT(*) FROM engineer; -- 1000000 rows
```
Get number of rows in the `country` table:
```sql
SELECT COUNT(*) FROM country; -- 240 rows
```
You can also update and discover more about the database if you want.

### 3. Create an AWS RDS PostgreSQL instance
Before migrating the database, you need to create an AWS RDS PostgreSQL instance. You can use the AWS Management Console or AWS CLI to create the RDS instance.

Use the following command to create an AWS RDS PostgreSQL instance:
```bash
aws rds create-db-instance \
    --db-instance-identifier postgres \
    --db-instance-class db.t2.micro \
    --engine postgres \
    --master-username postgres \
    --master-user-password postgres \
    --allocated-storage 20 \
    --no-publicly-accessible
```

Wait for a few minutes for the RDS instance to be created.
