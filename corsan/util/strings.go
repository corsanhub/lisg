package util

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// func GetFrame(skipFrames int) runtime.Frame {
// 	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
// 	targetFrameIndex := skipFrames + 2

// 	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
// 	programCounters := make([]uintptr, targetFrameIndex+2)
// 	n := runtime.Callers(0, programCounters)

// 	frame := runtime.Frame{Function: "unknown"}
// 	if n > 0 {
// 		frames := runtime.CallersFrames(programCounters[:n])
// 		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
// 			var frameCandidate runtime.Frame
// 			frameCandidate, more = frames.Next()
// 			if frameIndex == targetFrameIndex {
// 				frame = frameCandidate
// 			}
// 		}
// 	}

// 	return frame
// }

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ArrayToString(items []string, delimiter string) string {
	var buffer bytes.Buffer
	for _, item := range items {
		buffer.WriteString(delimiter + item + delimiter + " ")
	}
	return buffer.String()
}

func PointersToString(items []*string, delimiter string) string {
	var buffer bytes.Buffer
	for _, item := range items {
		buffer.WriteString(delimiter + *item + delimiter + " ")
	}
	return buffer.String()
}

func Xs(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return stringWithCharset(length, charset)
}
