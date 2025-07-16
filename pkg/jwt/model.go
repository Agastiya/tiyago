package jwt

import "github.com/agastiya/tiyago/dto"

type Jwt struct {
	JwtPackage dto.JwtSetting
}

func NewJwt(jwt dto.JwtSetting) IJwt {
	return &Jwt{
		JwtPackage: jwt,
	}
}
