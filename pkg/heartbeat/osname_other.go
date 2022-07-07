//go:build !linux

package heartbeat

func osName() string {
	return ""
}
