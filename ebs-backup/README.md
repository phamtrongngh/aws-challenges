# EBS Volume Management Challenge

## Objectives:
- Write an app backend server that receives message from clients and save them to a log file.
- Launch an EC2 instance using the Amazon Linux AMI.
- Create and associate an EBS volume with the EC2 instance.
- Connect to the instance via SSH.
- Generate a file system on the EBS volume, mount it, and read/write data.
- Create a snapshot of the EBS volume.
- Terminate the initial EC2 instance, remove the original EBS volume, and create a new volume from the snapshot.
- Attach and mount this volume to a new EC2 instance to retrieve the stored data.
- Clean up all resources to prevent unnecessary expenses.

## High-Level Architecture:
Refer to the architecture diagram you provided. This includes:

- A VPC in the us-east-1 region with a public subnet.
- An EC2 instance with a root volume and an additional backup volume.
- EBS snapshots for backups.

## Deliverables:
- Successfully perform each step outlined in the objectives.
- Verify data persistence by retrieving data from the EBS snapshot on a new instance.
- Ensure all resources are cleaned up after completion.