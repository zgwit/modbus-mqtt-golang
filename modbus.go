package modbus_mqtt

import "github.com/goburrow/modbus"

var modbusClient modbus.Client

func modbusOpen() error {
	handler := modbus.NewRTUClientHandler("")
	err := handler.Connect()
	if err != nil {
		return err
	}

	modbusClient = modbus.NewClient(handler)
	//client.ReadHoldingRegisters()
	return nil
}

func ttt() {
	//buf, err := modbusClient.ReadCoils(0, 10)
}
