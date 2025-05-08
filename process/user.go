package process

import (
	"log"
	"sync"
	"time"

	data_pattern "github.com/kacpi52/t1_users/data"
	repository_createUser "github.com/kacpi52/t1_users/repository/createUser"
	repository_fetchUser "github.com/kacpi52/t1_users/repository/fetchUser"
)

func GetAndCreateUserConcurrently() *data_pattern.UserCredentialsCollection {
	successChan := make(chan bool)
	successCount := 0
	successAddedUsersCol := &data_pattern.UserCredentialsCollection{}

	go func() {
		for range successChan {
			successCount++
			log.Printf("Created %d/%d users", successCount, data_pattern.USER_COUNT_TARGET)
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
					log.Printf("worker %d: error fetching user: %v", workerId, err)
					return
				}
				err = repository_createUser.SaveUserToLinux(*user)
				if err == nil {
					successAddedUsersCol.Mutex.Lock()
					successAddedUsersCol.Collection = append(successAddedUsersCol.Collection, *user)
					successAddedUsersCol.Mutex.Unlock()
					successChan <- true
				}
			}(i)
		}

		wg.Wait()
		time.Sleep(data_pattern.FETCH_TIME_DELAY)
	}

	close(successChan)

	return successAddedUsersCol
}
