package driver

/*

driver用来实现subscriber。

OnEvent回调：
	1. 将状态同步到数据库
	2. 将状态变更通知发送到第三方组件。


*/

func Init() {
	initRuleStatus()
	initRuleActive()
}

const eventDriverLogTitle = "[EventDriver]"
