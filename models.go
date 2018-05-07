package mcmqttrouter

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"os"
)

type MQTTRouter struct {
	Client *mqtt.Client
	prefix string
	Routes []string
	standardQOS byte
}

func NewMQTTRouter(c *mqtt.Client, standardQOS byte) *MQTTRouter {
	return &MQTTRouter{
		c,
		"",
		[]string{},
		standardQOS,
	}
}

func (this *MQTTRouter) Subscribe(path string, handler mqtt.MessageHandler) {
	this.Routes = append(this.Routes, path)
	if token := (*this.Client).Subscribe(this.prefix + path, 0, handler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func (this *MQTTRouter) UnSubscribeFromAll() {
	if token := (*this.Client).Unsubscribe(this.Routes...); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Println("Usubscribed from: ")
	for _, t := range this.Routes {
		fmt.Println("Topic name: " + t)
	}
	fmt.Println("")
}

func (this *MQTTRouter) PublishCustom(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return (*this.Client).Publish(topic, qos, retained, payload)
}

func (this *MQTTRouter) Publish(topic string, payload interface{}) mqtt.Token {
	return (*this.Client).Publish(topic, this.standardQOS, false, payload)
}

func (this *MQTTRouter) Group(path string) *MQTTRouter{
	return &MQTTRouter{this.Client, this.prefix + path, []string{}, this.standardQOS}
}
