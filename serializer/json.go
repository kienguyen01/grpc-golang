package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func WriteProtobufIntoJsonFile(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts: false,
	}

	return marshaler.MarshalToString(message)
}