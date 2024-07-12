include .env

create:
	aws cloudformation create-stack --stack-name ${STACK_NAME} --template-body file://cfn-templates/ebs-volume-management.yaml --capabilities CAPABILITY_NAMED_IAM

delete:
	aws cloudformation delete-stack --stack-name ${STACK_NAME}