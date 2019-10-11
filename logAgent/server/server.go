package server

import (
	"encoding/json"
	"logAgent/producer"
	"logAgent/task"
	"time"

	"github.com/astaxie/beego/logs"
)

// 启动logagent服务
func ServerRun() (err error) {
	logs.Info("Log Agent start is running...")
	for true {
		// 获取一行日志数据
		msg := task.GetOneLine()
		// 发送一行日志数据到kafka
		err = sendToKafka(msg.Msg, msg.Topic)
		if err != nil {
			logs.Error("send to kafka msg:[%v] topic:[%v] failed, err:[%v]", msg.Msg, msg.Topic, err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

// 发送数据到kafka
func sendToKafka(msg task.LogContent, topic string) (err error) {
	smsg, err := json.Marshal(&msg)
	if err != nil {
		logs.Error("send to kafka marshal failed --> msg: [%v], topic:[%s], error: %s", msg, topic, err)
		return
	}
	logs.Debug("send to kafka --> msg:[%v], topic:[%v]", string(smsg), topic)
	err = producer.SendMsgToKafka(string(smsg), topic)
	return
}