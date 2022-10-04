package main

import (
	"context"
	"fmt"
	"myfinnplan/entity"
	"myfinnplan/repository"
	"myfinnplan/service/userproto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type user struct {
	repo repository.UserRepository
	userproto.UnimplementedUserServiceServer
}

func main() {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	env := entity.SetEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASS, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	userproto.RegisterUserServiceServer(srv, NewUser(db))
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
func NewUser(db *gorm.DB) *user {
	return &user{repo: repository.NewUserRepository(db)}
}

func (u *user) GetAllUSer(ctx context.Context, empty *userproto.Empty) (*userproto.UserList, error) {
	listUser, err := u.repo.FindAll()

	if err != nil {
		return &userproto.UserList{}, err
	}

	var result userproto.UserList

	for _, v := range listUser {
		user := userproto.User{
			Id:          int64(v.Id),
			UserName:    v.UserName,
			Password:    v.Password,
			Email:       v.Email,
			Telephone:   v.Telephone,
			Photo:       v.Photo,
			IsVerified:  v.IsVerified,
			CreatedBy:   v.CreatedBy,
			CreatedDate: timestamppb.New(v.CreatedDate),
			UpdatedDate: timestamppb.New(v.CreatedDate),
			DeletedDate: timestamppb.New(v.DeletedDate),
			UpdatedBy:   v.UpdatedBy,
			DeletedBy:   v.DeletedBy,
		}
		result.User = append(result.User, &user)
	}
	return &result, nil
}
