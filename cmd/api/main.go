package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kentakki416/grpc-rock-paper-scissors/pb"
	"github.com/kentakki416/grpc-rock-paper-scissors/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 起動するポート番号の指定
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}

	// gRPCサーバーの生成
	server := grpc.NewServer()
	//　自動生成された関数に、サーバーと実際に処理を行うメソッドを実装したハンドラを設定する
	pb.RegisterRockPaperScissorsServiceServer(server, service.NewRockPaperScissorsService())

	// シリアライズせずに`grpc_cli`で動作確認するためにサーバーリフレクションを有効化
	reflection.Register(server)
	// サーバー起動
	server.Serve(listenPort)
}
