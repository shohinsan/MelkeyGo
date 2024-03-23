
AWS_ACCOUNT_ID := $(shell aws sts get-caller-identity --query Account --output text)
AWS_REGION := $(shell aws configure get region)
FILE := ./lambda/function.zip
EVENTS := ./lambda/bootstrap

## =======================================================
## General

tidy:
	go mod tidy
	go mod vendor

## =======================================================
## Cloud Development Kit (CDK)

cdk-build:
	@GOOS=linux GOARCH=amd64 go build -o ${EVENTS}
	@zip ${FILE} ${EVENTS}


cdk-zip:
	zip ${FILE} ${EVENTS}

cdk-diff:
	cdk diff

cdk-bootstrap:
	cdk bootstrap aws://$(AWS_ACCOUNT_ID)/$(AWS_REGION)
	
cdk-deploy:
	cdk deploy

cdk-curl:
	curl -X POST POST https://{API_ID}/register \
		-H "Content-Type: application/json" \
		-d '{"username": "insidious", "password": "primePass"}'



