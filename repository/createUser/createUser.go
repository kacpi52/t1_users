package repository_createUser

import (
	"errors"
	"fmt"
	"os/exec"

	data_pattern "github.com/kacpi52/t1_users/data"
)

func SaveUserToLinux(user data_pattern.UserCredentials) error {
	checkCmd := exec.Command("id", user.Username)
	if err := checkCmd.Run(); err == nil {
		return errors.New("user already exists")
	}

	fullName := fmt.Sprintf("%s %s", user.Name, user.Surname)

	createCmd := exec.Command("sudo", "useradd", "-m", "-c", fullName, user.Username)
	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("error during user create: %v", err)
	}

	return nil
}
