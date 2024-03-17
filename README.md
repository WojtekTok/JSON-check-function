# JSON-check-function

## Task Description
Write a method verifying the input JSON data. Input data format is defined as AWS::IAM::Role Policy - definition and example ([AWS IAM Role JSON definition and example](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html)). Input JSON might be read from a file. 
Method shall return logical false if an input JSON Resource field contains a single asterisk and true in any other case. 

## Requirements
Installed Go, version: 1.22.1 or newer ([download](https://go.dev/doc/install))

## Running
After cloning the repository open cmd in the repository's directory and type:
<br>`go run verify.go`<br/>
Then paste path to JSON file you want to verify. There is an example.json in the repository that you can verify by simpy typing:
<br>`example.json`<br/>
### Tests
To run all tests type:
<br>`go test`<br/>
To run tests and see the outcome of each of them type:
<br>`go test -v`<br/>
