package grpc

import (
	"context"
	userpb "user-service/proto/user"
	"user-service/usecase"
)

type ServerGRPC struct {
	userpb.UnimplementedUserServiceServer
	UseCase *usecase.UseCase
}

func (s *ServerGRPC) GetUser(ctx context.Context, in *userpb.GetUserReq) (*userpb.User, error) {
	user, err := s.UseCase.UserUseCase.GetOneUser(in.GetId())
	if err != nil {
		return nil, err
	}
	reply := &userpb.User{
		Id:       user.ID,
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
	}
	return reply, nil
}
func (s *ServerGRPC) GetBalance(ctx context.Context, in *userpb.String) (*userpb.Amount, error) {
	wallet, err := s.UseCase.WalletUseCase.GetOneWallet(in.GetValue())
	if err != nil {
		return nil, err
	}
	reply := &userpb.Amount{
		Value: int64(wallet.Balance),
	}
	return reply, nil
}
