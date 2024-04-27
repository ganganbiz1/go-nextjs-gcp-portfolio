# go-nextjs-gcp-portfolio

## terraform

1. GCPのAPIは手動で有効にしてください。

2. GCPのシークレットマネジャーは手動で値を設定してください。

3. 環境変数を設定してください。
```
export GOOGLE_CREDENTIALS="$(< tf-dev.json)"
export TF_ENV=dev
```