Build and Deploy:
template.yaml:
Properties:
      CodeUri: golang-hello.zip
      Handler: main
      Runtime: go1.x

cd golang-hello
unset GOROOT
unset GOPATH
# skip these:
# export GOPATH=`pwd`
# go get github.com/aws/aws-lambda-go/lambda
# go get github.com/aws/aws-lambda-go/events
GOARCH=amd64 GOOS=linux go build -o bin/main main.go
cd bin
zip -r ../../golang-hello.zip main
cd ../..
sam local invoke
sam package --s3-bucket charliesappdata
sam deploy --stack-name golangsamstack --s3-bucket charliesappdata --capabilities CAPABILITY_IAM

Issues:
exec format error","errorType":"PathError"
1) compile for amd linux by setting 
2) main must be defined in main pkg: root folder, not an arbitrary folder

