# grpc-rock-paper-scissors
 
 gRPC（Unary RPC）でじゃんけんを行うシステム
 
# DEMO
 
"hoge"の魅力が直感的に伝えわるデモ動画や図解を載せる
 
# Features
 
"hoge"のセールスポイントや差別化などを説明する
 
# Requirement
 
"hoge"を動かすのに必要なライブラリなどを列挙する
 
* huga 3.5.2
* hogehuga 1.0.2
 
# Installation
 
プロジェクトで使用したライブラリと実行した手順
 
```bash
// grpcライブラリ
go get -u google.golang.org/grpc

// protoファイルのコンパイルコマンド
brew install protobuf

// Go言語のコード自動生成プラグイン
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

// コードの実行
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/rock-paper-scissors.proto

// ※エラーが出る場合は下記のコマンドを実行
export GOPATH=$HOME/go
PATH=$PATH:$GOPATH/bin
```
 
# Usage

 
```bash
git clone https://github.com/hoge/~
cd examples
python demo.py
```
 
# Note
 
注意点などがあれば書く
 
# Author
 
作成情報を列挙する
 
* 作成者
* 所属
* E-mail
 
# License
ライセンスを明示する
 
"hoge" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
 
社内向けなら社外秘であることを明示してる
 
"hoge" is Confidential.