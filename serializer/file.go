package serializer

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
)



func WriteProtobufToBinaryFile(message proto.Message, fileName string) error{
	//serial to binary with .Marshal
	data, err := proto.Marshal(message)

	if err != nil{
		return fmt.Errorf("cannot serialize message")
	}

	err = os.WriteFile(fileName, data, 0644)

	if err != nil{
		return fmt.Errorf("cannot write file")
	}


	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error{
	data, err := os.ReadFile(filename)

	if err != nil{
		return fmt.Errorf("can not read binary data from file")
	}

	err = proto.Unmarshal(data, message)

	if err != nil{
		return fmt.Errorf("can not unserialize data")

	}

	return nil 
}