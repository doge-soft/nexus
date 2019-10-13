package modules

func InitAllModules() {
	InitRedisMQModule()
	InitRedisMutex()
	InitTasksModule()
}
