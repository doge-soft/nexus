package modules

import (
	"github.com/doge-soft/dogego_module_tasks"
	"github.com/doge-soft/dogego_module_tasks/managers"
	"github.com/doge-soft/dogego_module_tasks/triggers"
)

var TaskManager *managers.TaskManager
var AsyncTrigger *triggers.AsyncTrigger
var TimeTrigger *triggers.TimeTrigger

func InitTasksModule() {
	TaskManager, AsyncTrigger, TimeTrigger = dogego_module_tasks.NewDogeGoTaskModule(RedisMQ, RedisMutex)

	registerTasks()
}

func registerTasks() {
}
