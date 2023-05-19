#!/bin/bash

templateFiles=(
    serverless.yaml
)

while getopts "f:" arg; do
  case $arg in
    f) stage=$OPTARG;;
  esac
done

if [[ "$stage" == "dev" || "$stage" == "staging" || "$stage" == "production" ]];
then
    for template in "${templateFiles[@]}";
    do
        echo "Deploying $template"
        make build
        npm run deploy:$stage
    done
else
    echo "Make sure to set the correct stage"
    exit 1
fi