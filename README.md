## 使用ツール
- [docker-compose](https://docs.docker.jp/compose/install/index.html)
- [sam cli](https://aws.amazon.com/jp/serverless/sam/)
- [Go](https://go.dev/)

## 実行方法
### LocalStack 起動
```bash
docker-compose up
```

### Go テストコード実行
```bash
go test ./lambda/... -v
```

### API Gateway 起動
```bash
sam build
sam local start-api --docker-network lstack
```

### API Gateway 経由で Lambda 実行
```bash
# ScanSample Lambda
curl http://localhost:3000/data
# GetObjectSample Lambda
curl http://localhost:3000/file
```

## 参考
- [Starting LocalStack with Docker-Compose](https://docs.localstack.cloud/getting-started/installation/#starting-localstack-with-docker-compose)
- [LocalStack Configuration](https://docs.localstack.cloud/references/configuration/)
- [LocalStack Initialization Hooks](https://docs.localstack.cloud/references/init-hooks/)
- [Integrations Go](https://docs.localstack.cloud/user-guide/integrations/sdks/go/)
