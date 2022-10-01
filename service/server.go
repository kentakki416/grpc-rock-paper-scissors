package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kentakki416/grpc-rock-paper-scissors/pb"
	"github.com/kentakki416/grpc-rock-paper-scissors/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func inti() {
	rand.Seed(time.Now().UnixNano())
}

// DBを使わずに対戦結果の履歴を表示できるように構造値を用意
type RockPaperScissorsService struct {
	numberOfGames int32
	numberOfWins  int32
	matchResults  []*pb.MatchResult
}

func NewRockPaperScissorsService() *RockPaperScissorsService {
	return &RockPaperScissorsService{
		numberOfGames: 0,
		numberOfWins:  0,
		matchResults:  make([]*pb.MatchResult, 0),
	}
}

func (s *RockPaperScissorsService) PlayGame(ctx context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	if req.HandShapes == pb.HandShapes_HAND_SHAPES_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors")
	}

	// ランダムに1~3の数値を生成し、相手のじゃんけんとの手とする
	opponentHandShapes := pkg.EncodeHandShapes(int32(rand.Intn(3) + 1))

	// じゃんけんの勝敗を判定
	var result pb.Result
	if req.HandShapes == opponentHandShapes {
		result = pb.Result_DRAW
	} else if (req.HandShapes.Number()-opponentHandShapes.Number()+3)%3 == 1 {
		result = pb.Result_WIN
	} else {
		result = pb.Result_LOSE
	}

	now := time.Now()
	// 自動生成された型を元に対戦結果を生成
	matchResult := &pb.MatchResult{
		YourHandShapes:     req.HandShapes,
		OpponentHandShapes: opponentHandShapes,
		Result:             result,
		CreateTime: &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	//　試合数を１増やし、プレイヤーが勝利した場合は勝利数も１増やす
	s.numberOfGames = s.numberOfGames + 1
	if result == pb.Result_WIN {
		s.numberOfWins = s.numberOfWins + 1
	}
	s.matchResults = append(s.matchResults, matchResult)

	// 自動生成されたレスポンス用のコードを使ってレスポンスを作り返却
	return &pb.PlayResponse{
		MatchResult: matchResult,
	}, nil
}

func (s *RockPaperScissorsService) ReportMatchResults(ctx context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	return &pb.ReportResponse{
		Report: &pb.Report{
			NumberOfGames: s.numberOfGames,
			NumberOfWins:  s.numberOfWins,
			MatchResults:  s.matchResults,
		},
	}, nil
}
