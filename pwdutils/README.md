
# 盐
salt = ...
iterations = 0
# 项目启动时运行
NewPwdUtil(salt, iterations)

# 密码加密（与python相同）
GetPwdUtil().PasswordHash(pwd)
# 密码 和 加密 的密码校验
GetPwdUtil().PasswordVerify(pwd, hash)


