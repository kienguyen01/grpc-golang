package serializer_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	pb "github.com/grpc-golang/pcbook/pb/proto"
	"github.com/grpc-golang/pcbook/sample"
	"github.com/grpc-golang/pcbook/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))

	err = WriteProtobufMessageIntoJsonFile(laptop1, jsonFile)
	require.NoError(t, err)
}

func WriteProtobufMessageIntoJsonFile(message proto.Message, fileName string) error {
	data, err := serializer.WriteProtobufIntoJsonFile(message)

	if err != nil {
		return fmt.Errorf("can not write into json")
	}

	err = os.WriteFile(fileName, []byte(data), 0644)

	if err != nil {
		return fmt.Errorf("can not write file from json")
	}

	return nil
}
