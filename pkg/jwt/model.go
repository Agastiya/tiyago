package jwt

import "github.com/agastiya/tiyago/dto"

type Jwt struct {
	JwtPackage dto.JwtSetting
}

type JwtConfig struct {
	Key string
	Exp int64
}

func NewJwt(jwt dto.JwtSetting) IJwt {
	return &Jwt{JwtPackage: jwt}
}
