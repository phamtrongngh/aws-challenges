# EBS Volume Management Challenge

## Objectives:
- Write an app backend server that receives message from clients and save them to a data log file. The server should also be able to retrieve the stored messages.
- Launch an EC2 instance using the Amazon Linux AMI.
- Create and associate an EBS volume with the EC2 instance.
- Connect to the instance via SSH.
- Generate a file system on the EBS volume, mount it, and read/write data.
- Create a snapshot of the EBS volume.
- Terminate the initial EC2 instance, remove the original EBS volume, and create a new volume from the snapshot.
- Attach and mount this volume to a new EC2 instance to retrieve the stored data.
- Clean up all resources to prevent unnecessary expenses.

## Steps:
> IMPORTANT! You must install and configure your AWS CLI with the necessary credentials before proceeding with the steps below.

### 1. Provision the infrastructure with CloudFormation
Set the CloudFormation stack name. Replace `<stack-name>` with any name you want.

```
$ export STACK_NAME=<stack-name>
```

Create the CloudFormation stack and wait for the stack creation to complete.

```
$ aws cloudformation create-stack \
    --stack-name $STACK_NAME \
    --template-body file://cfn-templates/init.yaml \
    --capabilities CAPABILITY_NAMED_IAM
```

### 2. Test APIs
Get the server public IP address of the backend server from the CloudFormation stack outputs:

```
$ export SERVER_PUBLIC_IP=$(aws cloudformation describe-stacks \
    --stack-name $STACK_NAME \
    --query "Stacks[0].Outputs[?OutputKey=='PublicIp'].OutputValue" \
    --output text)
```

Post a message to the backend server

```
$ curl -X POST -H "Content-Type: application/json" -d '{"message": "Hello, World!"}' http://${SERVER_PUBLIC_IP}:3000
```

Retrieve the stored messages from the backend server

```
$ curl http://${SERVER_PUBLIC_IP}:3000
```

### 3. Create a snapshot of the EBS volume to backup the data
Get the instance ID from the CloudFormation stack outputs:
```
$ export INSTANCE_ID=$(aws cloudformation describe-stacks \
    --stack-name $STACK_NAME \
    --query "Stacks[0].Outputs[?OutputKey=='InstanceId'].OutputValue" \
    --output text)
```

Get the volume ID of the secondary EBS volume attached to the EC2 instance which contains the stored app messages.
    
```
$ export VOLUME_ID=$(aws ec2 describe-volumes \
    --filters Name=attachment.instance-id,Values=${INSTANCE_ID} \
    --query "Volumes[?Attachments[0].Device == '/dev/sdb'].VolumeId" \
    --output text)
```

Create the snapshot and store the snapshot ID

```
$ export SNAPSHOT_ID=$(aws ec2 create-snapshot \
                --volume-id ${VOLUME_ID} \
                --description "Snapshot of the secondary EBS volume" \
                --query "SnapshotId" \
                --output text)
```
    

### 4. Terminate the EC2 instance.
Now the snapshot has been created, so it's safe to terminate the EC2 instance and remove all the volumes attached to it.
    
```
$ aws ec2 terminate-instances --instance-ids ${INSTANCE_ID}
```

### 5. Create a new EC2 instance with the snapshot of the EBS volume
Update the CloudFormation stack with the snapshot ID to create a new EBS volume from the snapshot and attach it to a new EC2 instance.

```
$ aws cloudformation update-stack \
    --stack-name $STACK_NAME \
    --template-body file://cfn-templates/init.yaml \
    --parameters ParameterKey=SnapshotId,ParameterValue=${SNAPSHOT_ID} \
    --capabilities CAPABILITY_NAMED_IAM
```


### 6. Test the new EC2 instance
Get the new server public IP address of the backend server from the CloudFormation stack outputs:

```
$ export NEW_SERVER_PUBLIC_IP=$(aws cloudformation describe-stacks \
    --stack-name $STACK_NAME \
    --query "Stacks[0].Outputs[?OutputKey=='PublicIp'].OutputValue" \
    --output text)
```

Retrieve the stored messages from the backend server

```
$ curl http://${NEW_SERVER_PUBLIC_IP}:3000
```
If the messages are retrieved successfully, the data has been successfully restored from the snapshot.

### 7. Clean up
Delete the CloudFormation stack to clean up all resources.

```
$ aws cloudformation delete-stack --stack-name $STACK_NAME
```

Note that the snapshot is not a part of the CloudFormation stack, so we need to delete the snapshot manually.

```
$ aws ec2 delete-snapshot --snapshot-id ${SNAPSHOT_ID}
```