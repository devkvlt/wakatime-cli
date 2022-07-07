//go:build linux

package heartbeat

import (
	"syscall"

	"github.com/wakatime/wakatime-cli/pkg/log"
)

func osName() string {
	var buf syscall.Utsname

	err := syscall.Uname(&buf)
	if err != nil {
		log.Debugf("Uname error: %s", err)

		return ""
	}

	arr := buf.Sysname[:]
	output := make([]byte, 0, len(arr))

	for _, c := range arr {
		if c == 0x00 {
			break
		}

		output = append(output, byte(c))
	}

	return string(output)
}
