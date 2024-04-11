package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"palworld/rpc/user/dao"
	"time"

	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	if in.Username == " " || in.Password == " " {
		return nil, errors.New("用户名或密码不能为空")
	}
	var ret = &user.LoginResp{}
	p := dao.User
	loginUser, err := p.WithContext(l.ctx).Where(p.Name.Eq(in.Username)).First()
	if err != nil {
		return nil, err
	}
	if !VerifyPassword(loginUser.Password, in.Password) {
		return nil, errors.New("密码错误")
	}
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	now := time.Now().Unix()

	token, err := GetJwtToken(secret, now, expire, loginUser.ID)
	if err != nil {
		return nil, err
	}
	ret.Code = http.StatusOK
	ret.Message = "ok"

	data := &user.LoginInfo{}
	data.Token = &user.Token{
		AccessToken:  token,
		AccessExpire: now + expire,
		RefreshAfter: now + (expire / 2),
	}

	data.User = &user.User{
		Id:        loginUser.ID,
		Username:  loginUser.Name,
		Role:      mapRole[int(loginUser.Role)],
		Avatar:    loginUser.Avatar,
		CreatedAt: loginUser.CreatedAt.Format("2006-02-01 15:04:05"),
	}

	ret.Data = data
	return ret, nil
}

// EncryptPassword 使用 bcrypt 对密码进行加密
func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword 检查加密的密码是否匹配明文密码
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
