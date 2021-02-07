### mqtt_tw
mqtt 连接 发布 订阅

go get github.com/tw-iot/mqtt_tw

### 示例
```
package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"github.com/tw-iot/mqtt_tw"
	"time"
)

func main()  {
	clientId := uuid.NewV4()
	fmt.Println(clientId)
	mqttInfo := mqtt_tw.NewMqttInfo("192.168.146.19", "",
		"", fmt.Sprintf("%s", clientId), 1883)
	mqtt_tw.MqttInit(&mqttInfo)
	subs(mqtt_tw.MqttTw)
	publi(mqtt_tw.MqttTw)
	for {
		time.Sleep(time.Second * 2)
	}
}

func publi(client mqtt.Client) {
	token := client.Publish("topic/test", 0, false, "hello")
	token.Wait()
}

func subs(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
```
