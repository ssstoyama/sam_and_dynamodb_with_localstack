#!/bin/bash
#
# LocalStack 起動時に AWS リソースを作成する
#

export AWS_DEFAULT_REGION=ap-northeast-1

# aws コマンドの向き先を Localstack に向ける
aws="aws --endpoint-url http://localhost:4566"

# setup 用ディレクトリに移動
cd "$(dirname "$0")"

# DynamoDB テーブル作成
${aws} dynamodb create-table --cli-input-json file://sample_table.json
# 初期データ投入
${aws} dynamodb put-item --table-name sample --item '{"id":{"S":"id1"}}'
${aws} dynamodb put-item --table-name sample --item '{"id":{"S":"id2"}}'
${aws} dynamodb put-item --table-name sample --item '{"id":{"S":"id3"}}'

# S3 バケット作成
${aws} s3api create-bucket --bucket sample-bucket --create-bucket-configuration LocationConstraint=ap-northeast-1
${aws} s3api put-object --bucket sample-bucket --key sample.txt --body ./sample.txt
