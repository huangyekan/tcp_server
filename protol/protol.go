package protol

import (
	"tcp_server/util"
)

const HEAD_LENGTH = 4

func Encode() {

}

func Decode(buf []byte, messages []string) (result []string, leftBuf []byte) {
	if len(buf) < HEAD_LENGTH {
		return messages, buf
	}
	messageLength := getMessageLength(buf[:HEAD_LENGTH])
	if messageLength > len(buf) - HEAD_LENGTH {
		return messages, buf
	}
	result = append(messages, string(buf[HEAD_LENGTH:HEAD_LENGTH + messageLength]))
	leftBuf = buf[HEAD_LENGTH + messageLength:len(buf)]
	return Decode(leftBuf, result)
}

func getMessageLength(b []byte) int {
	return util.BytesToInt(b)
}
