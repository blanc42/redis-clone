package protocol

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func DecodeRESP(reader *bufio.Reader) ([]string, error) {
	// Simplified RESP decoding for array type only
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	if line[0] != '*' {
		return nil, fmt.Errorf("expected array, got: %s", line)
	}

	count, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return nil, err
	}

	result := make([]string, count)
	for i := 0; i < count; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		if line[0] != '$' {
			return nil, fmt.Errorf("expected bulk string, got: %s", line)
		}

		length, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			return nil, err
		}

		value := make([]byte, length)
		_, err = io.ReadFull(reader, value)
		if err != nil {
			return nil, err
		}

		result[i] = string(value)

		// Read and discard CRLF
		_, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func EncodeRESP(writer io.Writer, value interface{}) error {
	switch v := value.(type) {
	case string:
		return encodeString(writer, v)
	case int:
		return encodeInteger(writer, v)
	case error:
		return encodeError(writer, v)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

func encodeString(writer io.Writer, s string) error {
	_, err := fmt.Fprintf(writer, "$%d\r\n%s\r\n", len(s), s)
	return err
}

func encodeInteger(writer io.Writer, i int) error {
	_, err := fmt.Fprintf(writer, ":%d\r\n", i)
	return err
}

func encodeError(writer io.Writer, err error) error {
	_, writeErr := fmt.Fprintf(writer, "-ERR %s\r\n", err.Error())
	return writeErr
}
