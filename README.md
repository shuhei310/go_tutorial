# go_tutorial
## 概要
Goを勉強する

## 開発環境
MacOS:Mojave 10.14.3  
Go:1.10.3  
Docker:18.09.2  
docker-compose:1.23.2

## Dockerの使い方
1. appディレクトリ配下に移動
2. Dockerをビルド  
   `docker-compose up --build`
3. コンテナに接続  
   `docker exec -i -t app_app_1 bash`
4. プログラムを実行  
   `go run ?[ファイル名]`