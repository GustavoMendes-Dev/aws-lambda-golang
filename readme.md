go mod init
go get github.com/aws/aws-lambda-go/events

# Compilar código

## Para Linux/Mac

GOOS=linux GOARCH=amd64 go build -o main main.go
CGO_ENABLED=0 go build -o bootstrap main.go

## Para Windows

Os desenvolvedores do Windows podem ter problemas para produzir um arquivo zip que marque o binário como executável no Linux. Para criar um .zip que funcione no AWS Lambda, a build-lambda-zipferramenta pode ser útil.

Obtenha a ferramenta

go.exe install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o bootstrap main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -o lambda-handler.zip bootstrap

# Criar ZIP

zip deployment.zip main

# Criar uma Função Lambda na AWS

aws lambda create-function \
 --function-name awsLambdaGoTest \
 --runtime go1.x \
 --handler main \
 --zip-file fileb://./deployment.zip \
 --role <seu-arn-do-papel>

Substitua <seu-arn-do-papel> pelo ARN do papel da AWS Lambda que tem as permissões adequadas. Certifique-se de que o papel tenha permissões para acessar outros serviços AWS, se necessário.

# Função Lambda

aws lambda invoke \
 --function-name awsLambdaGoTest \
 --payload '{}' \
 output.json

https://github.com/aws/aws-lambda-go
