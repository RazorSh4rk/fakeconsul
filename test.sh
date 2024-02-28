#!/bin/bash

docker build -t fakeconsul .
docker run -d -p 8500:8500 fakeconsul

echo "Setting key-value pair..."
curl -X PUT -d "value1" http://localhost:8500/v1/kv/test-key
echo ""

echo "Getting value for the key..."
curl http://localhost:8500/v1/kv/test-key
echo ""

echo "Deleting the key..."
curl -X DELETE http://localhost:8500/v1/kv/test-key
echo ""

echo "Getting value for the key after deletion..."
curl http://localhost:8500/v1/kv/test-key
echo ""
