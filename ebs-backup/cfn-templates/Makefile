include .env

create-stack:
	aws cloudformation create-stack --stack-name ${STACK_NAME} --template-body file://root.yaml --capabilities CAPABILITY_NAMED_IAM

delete-stack:
	aws cloudformation delete-stack --stack-name ${STACK_NAME}

update-stack:
	aws cloudformation update-stack --stack-name ${STACK_NAME} --template-body file://root.yaml --capabilities CAPABILITY_NAMED_IAM

recreate-stack: delete-stack create-stack

cancel-update-stack:
	aws cloudformation cancel-update-stack --stack-name ${STACK_NAME}