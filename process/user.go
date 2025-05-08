package process

import (
	"fmt"
	"sync"
	"time"

	data_pattern "github.com/kacpi52/t1_users/data"
	repository_createUser "github.com/kacpi52/t1_users/repository/createUser"
	repository_fetchUser "github.com/kacpi52/t1_users/repository/fetchUser"
)

func GetUAndCreteUserConcurrently() {
	successChan := make(chan bool)
	successCount := 0
	successAddedUsersArray := []data_pattern.UserCredentials{}

	go func() {
		for range successChan {
			successCount++
			fmt.Printf("Created %d/%d users \n", successCount, data_pattern.USER_COUNT_TARGET)
			if successCount >= data_pattern.USER_COUNT_TARGET {
				close(successChan)
				return
			}
		}
	}()

	for {
		if successCount >= data_pattern.USER_COUNT_TARGET {
			break
		}

		var wg sync.WaitGroup
		for i := 0; i < data_pattern.FETCH_STEP_QUANTITY; i++ {
			wg.Add(1)
			go func(workerId int) {
				defer wg.Done()
				user, err := repository_fetchUser.GetAndPrepareUserData()
				if err != nil {
					return
				}
				err = repository_createUser.SaveUserToLinux(*user)
				if err == nil {
					successAddedUsersArray = append(successAddedUsersArray, *user)
					successChan <- true
				}
			}(i)
		}
		
		wg.Wait()
		time.Sleep(data_pattern.FETCH_TIME_DELAY)
	}
}
