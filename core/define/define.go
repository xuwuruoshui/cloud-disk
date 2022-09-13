package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id int
	Identity string
	Name string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var TokenExpireTime = 3600*12
var RefreshTokenExpireTime = 3600*24

// 验证码长度
var CodeLength = 6

// 过期时间
var CodeExpire = 3600*24

// 图片上传
var SecretID = "xxx"
var SecretKey = "xxx"
var CosBucket = "xxx"

// 分页默认参数
var PageSize  = 20

// 时间格式化
var Datetime = "2006-01-02 15:04:05"
