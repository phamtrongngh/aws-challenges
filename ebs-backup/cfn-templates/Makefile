include .env

create-stack:
	aws cloudformation create-stack --stack-name ${STACK_NAME} --template-body file://root.yaml --capabilities CAPABILITY_NAMED_IAM

delete-stack:
	aws cloudformation delete-stack --stack-name ${STACK_NAME}

recreate-stack: delete-stack create-stack