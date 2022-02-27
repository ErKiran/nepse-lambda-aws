### More than Hello World Example of Lambda function   

##### This repo is just for learning purpose and I plan to create other repos too to enhance my knowledge on AWS and I will be using part of code from my Nepse Application. 


To Deploy the Lambda Function, you need to follow the below steps: 
1. Build the Binary from the source code  
`GOOS=linux GOARCH=amd64 go build main.go`
2. Zip the build binary file.  
`zip main.zip test.crt main`
3. Create the lambda function using the zip file.  
`aws lambda create-function --function-name nepse-lambda-aws --runtime go1.x --role <role-arn> --handler main --zip-file fileb://main.zip` or you can do it manually as well.
4. Check if lambda function is created or not using  
`aws lambda get-function --function-name nepse-lambda-aws`

To update the Lambda function code  

`aws lambda update-function-code --function-name nepse-lambda-aws --zip-file fileb://main.zip --publish`

To update the Lamda function configuration 

`aws lambda update-function-configuration --function-name nepse-lambda-aws --environment "Variables={NEPSE=https://*******.com/}"`