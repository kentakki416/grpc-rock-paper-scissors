# grpc-rock-paper-scissors
 
 gRPC（Unary RPC）でじゃんけんを行うシステム
 
# DEMO
 
```
start Rock-paper-scissors game.
1: play game
2: show match results
3: exit

Invalid commnad
1: play game
2: show match results
3: exit
1
Please enter Rock, Paper, or Scissors.
1: Rock
2: Paper
3: Scissors
please enter >1
***********************************
Your hand shapes: ROCK 
Opponent hand shapes: SCISSORS 
Result: WIN 
***********************************

```
 
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

// デバッグ用のツールgrpc_cliをインストール
brew tap grpc/grpc
brew install grpc
```
 
# Usage

 
```bash
git clone https://github.com/kentakki416/grpc-rock-paper-scissors.git

// ----------サーバー側の動作確認------------
go run ./cmd/api
// 別のターミナルを立ち上げて下記コマンドを実行
grpc_cli ls localhost:50051 game.RockPaperScissorsService
// 以下の２つが返ってくればサーバーが立ち上がっている
PlayGame
ReportMatchResults
//実際にじゃんけんを行う（1:グーを指定）
grpc_cli call localhost:50051 game.RockPaperScissorsService.PlayGame 'handShapes: 1'
//対戦結果の確認
grpc_cli call localhost:50051 game.RockPaperScissorsService.ReportMatchResults ''

// --------クライアント側の動作確認-----------
go run ./cmd/cli
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