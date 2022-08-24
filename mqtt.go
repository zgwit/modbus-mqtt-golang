package modbus_mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

var client mqtt.Client

func mqttConnect(cfg MqttOptions) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Server).SetClientID(cfg.ClientId).SetCleanSession(true)
	opts.SetUsername(cfg.Username).SetPassword(cfg.Password)
	opts.SetConnectRetry(true)
	opts.SetAutoReconnect(true).SetMaxReconnectInterval(time.Second * 30)

	opts.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		log.Println("MQTT receive ", message.Topic())
		//TODO 处理收到的消息

	})

	client = mqtt.NewClient(opts)
	//modbus

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	log.Println("MQTT Connected to ", cfg.Server)

	return nil
}

func mqttClose() {
	client.Disconnect(0)
}
