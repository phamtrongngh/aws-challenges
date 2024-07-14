include .env

create-stack:
	aws cloudformation create-stack --stack-name ${STACK_NAME} --template-body file://cfn-templates/init.yaml --capabilities CAPABILITY_NAMED_IAM

delete-stack:
	aws cloudformation delete-stack --stack-name ${STACK_NAME}

update-stack:
	aws cloudformation update-stack --stack-name ${STACK_NAME} --template-body file://cfn-templates/init.yaml --capabilities CAPABILITY_NAMED_IAM

recreate-stack: delete-stack create-stack

cancel-update-stack:
	aws cloudformation cancel-update-stack --stack-name ${STACK_NAME}

detech-drift:
	aws cloudformation detect-stack-drift --stack-name ${STACK_NAME}

deploy-stack:
	aws cloudformation deploy --stack-name ${STACK_NAME} --template-file init.yaml --capabilities CAPABILITY_NAMED_IAM

create-snaphot:
	INSTANCE_ID=$(aws cloudformation describe-stacks --stack-name ${STACK_NAME} --query "Stacks[0].Outputs[?OutputKey=='InstanceId'].OutputValue" --output text)
	aws ec2 create-snapshot --volume-id ${VOLUME_ID} --description "Backup for volume ${VOLUME_ID}"