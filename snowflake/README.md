
# 项目编码
WorkerId = 1
# 服务器编码
DataCenterId = 1
# 项目启动时运行
NewFlake(WorkerId, DataCenterId)

# 获取int64类型
GetFlake().GetNextID()
# 获取string类型
GetFlake().GetNextStringID()


