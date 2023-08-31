package pwdutils

type PwdUtil struct {
	salt       string
	iterations int
}

func NewPwdUtil(salt string, iterations int) *PwdUtil {
	// 根据项目替换，值越大加密时间越长
	if iterations == 0 {
		iterations = 1000
	}

	if pwdUtil == nil {
		pwdUtil = &PwdUtil{
			salt:       salt,
			iterations: iterations,
		}
	}

	return pwdUtil
}
