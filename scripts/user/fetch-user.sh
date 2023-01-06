#!/bin/sh

# This is legacy for reference if we ever use shell scripts in the future

DIR="$( cd "$( dirname "$0" )" && cd .. && cd variables && pwd )"

while getopts e:u: flag
do
  case "${flag}" in
    e) env=${OPTARG};;
    u) username=${OPTARG};;
  esac
done

if [ $env = "local" ]; then
  table=`cat ${DIR}/local/table.txt`
  region=`cat ${DIR}/local/region.txt`
  endpoint=`cat ${DIR}/local/endpoint.txt`
elif [ $env = "prod" ]; then
  table=`cat ${DIR}/prod/table.txt`
  region=`cat ${DIR}/prod/region.txt`
  endpoint=`cat ${DIR}/prod/endpoint.txt`
else 
  echo "who knows"
fi;

aws dynamodb get-item \
  --table-name $table \
  --endpoint-url $endpoint \
  --region $region \
  --key '{ "PK": {"S": "ACCOUNT#'"$username"'"}, "SK": {"S": "ACCOUNT#'"$username"'"}}'
