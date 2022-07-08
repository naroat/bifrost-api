package cron

import (
	"bifrost/server/global"
	"fmt"
)

type DemoCron struct {
}

func (dc *DemoCron) run() {
	fmt.Println("demo cron")
}

func Demo() {
	//获取Timer对象
	timerTask := global.GVA_Timer

	//添加任务 - */10 * * * * *
	_, err := timerTask.AddTaskByFunc("testFunc", "@daily", func() {
		fmt.Println("testFunc")
	})
	if err != nil {
		fmt.Println("err AddTaskByFunc")
		return
	}

	//timerTask.StopTask("testFunc")

	//timerTask.Remove("testFunc", id)

	timerTask.Close()
}
