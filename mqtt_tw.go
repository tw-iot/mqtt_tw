package mqtt_tw

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var MqttTw mqtt.Client

func MqttInit(port int, ip, username, password, clientId, willTopic, willMsg string) {
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)

	opts := mqtt.NewClientOptions()
	// MQTT的连接设置
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port))
	// clientId
	opts.SetClientID(clientId)
	// 设置连接的用户名
	opts.SetUsername(username)
	// 设置连接的密码
	opts.SetPassword(password)
	// 设置是否清空session,这里如果设置为false表示服务器会保留客户端的连接记录，
	// 把配置里的 cleanSession 设为false，客户端掉线后 服务器端不会清除session，
	// 当重连后可以接收之前订阅主题的消息。当客户端上线后会接受到它离线的这段时间的消息
	opts.SetCleanSession(true)
	// 自动重连
	opts.SetAutoReconnect(true)
	payload := []byte(willMsg)
	// 设置“遗嘱”消息的话题，若客户端与服务器之间的连接意外中断，服务器将发布客户端的“遗嘱”消息。
	opts.SetBinaryWill(willTopic, payload, 0, false)
	// 设置会话心跳时间 单位为秒 服务器会每隔1.5*20秒的时间向客户端发送个消息判断客户端是否在线，但这个方法并没有重连的机制
	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	MqttTw = mqtt.NewClient(opts)
	if token := MqttTw.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//全局 MQTT pub 消息处理
	//fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	//连接的回调
	//fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	//连接丢失的回调
	//fmt.Printf("Connect lost: %v", err)
}

func MqttDisconnect() {
	MqttTw.Disconnect(250)
}
