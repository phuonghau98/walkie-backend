package main

import (
	"context"
	"github.com/phuonghau98/walkie3/srv/user/proto/user"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Service) Authenticate(ctx context.Context, req *user.User) (*user.TokenInfo, error) {
	if len(req.Email) == 0 || len(req.Password) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Insufficient information")
	}
	foundUser, err := s.repo.GetByEmail(req.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password)) != nil {
		return nil, status.Errorf(codes.Unauthenticated, "You have entered an invalid username or password")
	}

	generatedToken, _ := EncodeToken(foundUser)
	rspTokenInfo := &user.TokenInfo{
		Token:   generatedToken,
		IsValid: true,
	}
	return rspTokenInfo, nil
}


func (s *Service) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	foundUser, err := s.repo.GetByEmail(user.Email)
	if foundUser != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Email " + user.Email + " has been used by another user")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPass)
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetUserProfile(ctx context.Context, _ *user.Empty) (*user.UserProfile, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	foundUser, err := s.GetUserFromToken(ctx, &user.TokenInfo{Token:headers["authorization"][0]})
	if err != nil {
		return nil, err
	}
	userProfile := &user.UserProfile{
		Id: foundUser.Id,
		Email: foundUser.Email,
		FirstName: foundUser.FirstName,
		LastName: foundUser.LastName,
	}
	return userProfile, nil
}

func (s *Service) GetUserFromToken(ctx context.Context, userToken *user.TokenInfo) (*user.User, error) {
	tokenClaims, err := DecodeToken(userToken.Token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid credentials")
	}
	foundUser, err := s.repo.GetByID(tokenClaims.UserInfo.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not existed")
	}
	userInfo := &user.User{
		Id: foundUser.ID.String(),
		FirstName: foundUser.Firstname,
		LastName: foundUser.Lastname,
		Email: foundUser.Email,
	}
	return userInfo, nil
}

func (s *Service) GetAllUsers(ctx context.Context, _ *user.Empty) (*user.UserList, error) {
	var users []*User
	var result []*user.User
	if  err := s.repo.db.Select("id, firstname, lastname").Find(&users).Error; err != nil {
		return nil, err
	}
	for _, e := range users {
		result = append(result, &user.User{Id: e.ID.String(), LastName: e.Lastname, FirstName:e.Firstname})
	}
	return &user.UserList{Users: result}, nil
}