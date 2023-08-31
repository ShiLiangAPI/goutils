package pwdutils

var pwdUtil *PwdUtil

func GetPwdUtil() *PwdUtil {
	if pwdUtil == nil {
		panic("Please run the InitFlake method first")
	}

	return pwdUtil
}
