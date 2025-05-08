package internal_fetchUser

import (
	"encoding/json"
	"io"
	"net/http"

	data_pattern "github.com/kacpi52/t1_users/data"
)

func fetchUserCollection() (*data_pattern.UserCollection, error) {
	resp, err := http.Get(data_pattern.URL_USER_API)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userCol := data_pattern.UserCollection{}
	err = json.Unmarshal(bytesData, &userCol)
	if err != nil {
		return nil, err
	}

	return &userCol, nil
}

func GetUserCredentials(userCol *data_pattern.UserCollection) data_pattern.UserCredentials {
	userCredentials := data_pattern.UserCredentials{}
	singleUser := userCol.Results[0]

	userCredentials.Name = singleUser.Name.First
	userCredentials.Surname = singleUser.Name.Last
	userCredentials.Username = singleUser.Login.Username

	return userCredentials
}
