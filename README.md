# go-nextjs-gcp-portfolio

## terraform

1. GCPのAPIは手動で有効にします。 
   Please enable GCP API manually.　

2. GCPのシークレットマネジャーは手動で値を設定します。  
   Please set the value manually for GCP's secret manager.

3. terraformコマンド実行前に環境変数を設定します。（terraformコマンドは`make apply`等）  
   Please set the environment variables before executing the terraform command. (terraform command is `make apply` etc.)
```
export GOOGLE_CREDENTIALS="$(< tf-dev.json)"
export TF_ENV=dev
```

## 環境構築
以下の手順を実施することでGCPのCloudRunにアプリケーションがデプロイされます。  
The application will be deployed to CloudRun on GCP by following the steps below.

1. make init-pre

2. make init-plan

3. make init-apply

4. Build and Push to Artifact Registry Backend（GithubActionsで手動実行）

5. Build and Push to Artifact Registry Frontend（GithubActionsで手動実行）

6. make-init

7. make-plan

8. make-apply

9. Deploy to Cloud Run Backend（GithubActionsで手動実行）

10. Deploy to Cloud Run Frontend（GithubActionsで手動実行）

11. DBへDDL実行

## CI
リポジトリへのpush、Pull requests作成時に単体テストと結合テストが実行されます。  
Unit tests and integration tests are executed when creating push and pull requests to the repository.

## CD
Pull requestsのマージ時にCloudRunへのデプロイされます。  
Deployed to CloudRun when merging pull requests.