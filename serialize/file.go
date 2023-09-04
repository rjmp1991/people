package serialize

import (
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshall roto message to binary: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary date to file: %w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	jsonData, err := MarshalJSON(message)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// MarshalJSON marshals a protocol buffer message to JSON bytes
func MarshalJSON(message proto.Message) ([]byte, error) {
	return json.MarshalIndent(message, "", "  ")
}

func WriteProtobufToJsonFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}
	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}
	return nil
}

/*
// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return UnmarshalJSON([]byte(data), message)
}
// UnmarshalJSON unmarshals JSON bytes to a protocol buffer message
func UnmarshalJSON(data []byte, message proto.Message) error {
	return json.Unmarshal(data, message)
}
*/
