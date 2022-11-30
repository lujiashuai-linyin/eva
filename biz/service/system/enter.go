package system

type ServiceGroup struct {
	UserService
	CasbinService
	JwtService
	OperationRecordService
}
