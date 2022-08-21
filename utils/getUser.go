package utils

import "os/user"

func GetUserName() string {
	user, errUser := user.Current()
	CheckError(errUser)

	username := user.Username

	for i := 0; i < len(username); i++ {
		if username[i] == 92 {
			username = username[i+1:]
		}
	}
	return username
}
