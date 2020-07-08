package config

import "hotelAPI/utils"

//AuthSwitch app
func AuthSwitch() bool {
	isTrue := utils.ViperGetEnv("Auth", "0")
	if isTrue == "0" {
		return true
	}
	return false
}
