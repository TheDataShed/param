#!/usr/bin/env bash

set -e

# Contact local parameter store with fake credentials
function prm() {
	AWS_REGION=us-east-1 \
	AWS_SESSION_TOKEN=fake \
	AWS_ACCESS_KEY_ID=fake \
	AWS_SECRET_ACCESS_KEY=fake \
	param --endpoint-url=http://localhost:4583 ${@}
}

docker pull localstack/localstack
docker run -d -p 4583:4583 -e SERVICES=ssm --name localstack localstack/localstack
# Allow localstack to stand up
sleep 7

echo "Setting parameter 'test' to 'teststring'..."
prm set test teststring
echo "Passed"

echo "Getting parameter test..."
[[ "$(prm show test)" = "teststring" ]]
echo "Passed"

echo "Setting parameter 'testing' to 'teststring2'..."
prm set testing teststring2

echo "List starting with 't' returns 'test' and 'testing'"
[[ "$(prm list -p t)" = "test
testing" ]]
echo "Passed"

echo "Overwriting value of 'test' to 'teststring3"
prm set -f test teststring3
echo "Passed"

echo "Getting parameter test..."
[[ "$(prm show test)" = "teststring3" ]]
echo "Passed"

echo "Deleting parameter 'test'"
prm delete test
echo "Passed"

echo "List starting with 't' returns 'testing'"
[[ "$(prm list -p t)" = "testing" ]]
echo "Passed"
