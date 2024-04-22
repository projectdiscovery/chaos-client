#!/bin/bash

echo "::group::Build chaos"
rm integration-test chaos 2>/dev/null
cd ../cmd/chaos
go build
mv chaos ../../integration_tests/chaos
echo "::endgroup::"

echo "::group::Build chaos integration-test"
cd ../integration-test
go build
mv integration-test ../../integration_tests/integration-test
cd ../../integration_tests
echo "::endgroup::"

./integration-test
if [ $? -eq 0 ]
then
  exit 0
else
  exit 1
fi
