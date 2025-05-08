package repository_fetchUser

import "testing"

func TestFetchUserCollection(t *testing.T) {
	userCol, err := fetchUserCollection()

	if err != nil {
		t.Errorf(`error, failed to fetch user %e`, err)
	}
	if len(userCol.Results) == 0 {
		t.Errorf(`error, collection is empty`)
	}
	if userCol.Results[0].Name.First == "" {
		t.Errorf(`error, user name is empty`)
	}
}

func TestUserCredentials(t *testing.T) {
	userCol, err := fetchUserCollection()

	if err != nil {
		t.Errorf(`error, failed to fetch user %e`, err)
	}
	if len(userCol.Results) == 0 {
		t.Errorf(`error, collection is empty`)
	}
	if userCol.Results[0].Name.First == "" {
		t.Errorf(`error, user name is empty`)
	}

	userCredentials := userCredentials(userCol)

	if userCol.Results[0].Name.First != userCredentials.Name {
		t.Errorf(`error, user name is not matching`)
	}
	if userCol.Results[0].Name.Last != userCredentials.Surname {
		t.Errorf(`error, user surname is not matching`)
	}
	if userCol.Results[0].Login.Username != userCredentials.Username {
		t.Errorf(`error, user username is not matching`)
	}
}

func TestGetAndPrepareUserCredentials(t *testing.T) {
	userCred, err := GetAndPrepareUserData()

	if err != nil {
		t.Errorf(`error, failed to fetch user %e`, err)
	}
	
	if userCred.Name == "" {
		t.Errorf(`error, user name is empty`)
	}
	if userCred.Username == "" {
		t.Errorf(`error, userName name is empty`)
	}

}
