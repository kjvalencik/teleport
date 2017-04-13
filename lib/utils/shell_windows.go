// +build windows

package utils

func GetLoginShell(username string) (string, error) {
	return "cmd.exe", nil
}
