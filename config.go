package modbus_mqtt

type Config struct {
	Mqtt MqttOptions
}

type ModbusOptions struct {
	Port string
}

type MqttOptions struct {
	Server   string
	ClientId string
	Username string
	Password string
}
