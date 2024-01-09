package util

import "github.com/ypMarkJo/Go-gin-samplepkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
