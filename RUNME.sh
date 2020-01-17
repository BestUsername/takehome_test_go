#/usr/bin/env bash

# Setup
TEMPOUTPUT="temp_output.json"

# kill owned processes if script killed
trap "kill 0" EXIT

# Start service
go run main.go &
SERVERPID=$!
sleep 1

# Make a request and capture the response
curl -X POST -d @ex_input.xml localhost:8080/process -s > $TEMPOUTPUT
if [ $? -ne 0 ]
then
    echo "ERROR POSTING DATA"
    exit 1
fi

# Stop service
kill $SERVERPID

# Compare output with expected
diff $TEMPOUTPUT ex_output.json
if [[ $? -ne 0 ]]
then
    echo "FAILURE"
else
    echo "SUCCESS"
fi
