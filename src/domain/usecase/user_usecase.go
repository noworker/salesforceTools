package usecase

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/infrastructure/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(userModel *model.User) (model.UserSignUpResponse, error)
	Login(userModel *model.User) (string, error)
}

type UserUsecase struct {
	ur repositories.IUserRepository
}

func NewUserUsecase(userRepo repositories.IUserRepository) IUserUsecase {
	return &UserUsecase{userRepo}
}

func (uu *UserUsecase) SignUp(userModel *model.User) (model.UserSignUpResponse, error) {
	retUser := model.UserSignUpResponse{}
	hash, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), 10)
	if err != nil {
		return retUser, err
	}
	userModel.Password = string(hash)
	if err != nil {
		return retUser, err
	}
	uuidWithHyphen := uuid.New()
	generatedUuid := strings.ReplaceAll(uuidWithHyphen.String(), "-", "")
	userModel.Id = generatedUuid
	err = uu.ur.CreateUser(userModel)
	if err != nil {
		return retUser, err
	}
	retUser = model.UserSignUpResponse{
		Id:   userModel.Id,
		Name: userModel.Name,
	}
	return retUser, nil
}

func (uc *UserUsecase) Login(userModel *model.User) (string, error) {
	storedUser := &model.User{}
	err := uc.ur.GetUserByNameId(storedUser, userModel.Name)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userModel.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.Id,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
