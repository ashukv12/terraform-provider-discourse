package client

import(
	"github.com/stretchr/testify/assert"
	"testing"
	"log"
	"time"
)

func TestClient_NewUser(t *testing.T) {
	testCases := []struct {
		testName  string
		newUser   string
		expectedResp *User
		expectErr bool
	}{
		{
			testName: "Success - User Created",
			newUser: "ashutosh.verma@clevertap.com",
			expectErr: false,
			expectedResp: &User{
				Name: "Ashutosh Verma",
				Id: 133,
				Username: "Ashutosh12345",
				Email: "ashutosh.verma@clevertap.com",
				Admin: false,
				Active: true,
			},
		},
		{
			testName: "User already exists",
			newUser: "ashutoshkverma12@gmail.com",
			expectErr: true,
			expectedResp: nil, 
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://clevertaptest.trydiscourse.com", "7c27290e2ccca7ae4427dfe518f68fb0659cd7df9908ff1cefbf034f9900a568", "Ashwinigaddagiwork")
			err := client.NewUser(tc.newUser)
			if tc.expectErr {
				log.Println("[CREATE ERROR TEST] : ", err)
				assert.Error(t, err)
				return
			}
			time.Sleep(60*time.Second)
			user, err := client.GetUser(tc.newUser)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_GetUser(t *testing.T) {
	testCases := []struct {
		testName     string
		Email     string
		expectErr    bool
		expectedResp *User
	}{
		{
			testName: "user exists",
			Email: "ashutosh.verma@clevertap.com",
			expectErr: false,
			expectedResp: &User {
				Email:  "ashutosh.verma@clevertap.com",
				Name: "Ashutosh Verma",
				Id: 133,
				Username: "Ashutosh12345",
				Admin: false,
				Active: true,
			},
		},
		{
			testName:     "user does not exist",
			Email:     "ashu@gmail.com",
			expectErr:    true,
			expectedResp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://clevertaptest.trydiscourse.com", "7c27290e2ccca7ae4427dfe518f68fb0659cd7df9908ff1cefbf034f9900a568", "Ashwinigaddagiwork")
			user, err := client.GetUser(tc.Email)
			if tc.expectErr {
				log.Println("[READ ERROR TEST] : ", err)
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_UpdateUser(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedUser *User
		expectErr   bool
		expectedResp *User
	}{
		{
			testName: "User exists",
			updatedUser: &User{
				Email:  "ashutoshkverma12@gmail.com",
				Name: "Ashutosh Verma Testing",
				Username: "Ashutosh",
			},
			expectErr: false,
			expectedResp: &User {
				Email:  "ashutoshkverma12@gmail.com",
				Name: "Ashutosh Verma Testing",
				Id: 127,
				Username: "Ashutosh",
				Admin: false,
				Active: true,
			},
		},
		{
			testName: "user does not exist",
			updatedUser: &User{
				Email:  "ashutosh@gmail.com",
				Name: "Ashutosh",
				Username: "Ashutosh123",
			},
			expectErr: true,
			expectedResp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://clevertaptest.trydiscourse.com", "7c27290e2ccca7ae4427dfe518f68fb0659cd7df9908ff1cefbf034f9900a568", "Ashwinigaddagiwork")
			err := client.UpdateUser(tc.updatedUser)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetUser(tc.updatedUser.Email)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

func TestClient_DeactivateUser(t *testing.T) {
	testCases := []struct {
		testName  string
		user_id  int
		expectErr bool
	}{
		{
			testName: "user exists deactivated worked",
			user_id: 81,
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://clevertaptest.trydiscourse.com", "7c27290e2ccca7ae4427dfe518f68fb0659cd7df9908ff1cefbf034f9900a568", "Ashwinigaddagiwork")
 			err := client.DeactivateUser(tc.user_id)
			log.Println(err)
			if tc.expectErr {
				log.Println("[DEACTIVATE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
		})
	}
}

func TestClient_ActivateUser(t *testing.T) {
	testCases := []struct {
		testName  string
		user_id  int
		expectErr bool
	}{
		{
			testName: "user exists activated worked",
			user_id: 81,
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://clevertaptest.trydiscourse.com", "7c27290e2ccca7ae4427dfe518f68fb0659cd7df9908ff1cefbf034f9900a568", "Ashwinigaddagiwork")
 			err := client.ActivateUser(tc.user_id)
			log.Println(err)
			if tc.expectErr {
				log.Println("[ACTIVATE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
		})
	}
}

func TestClient_DeleteUser(t *testing.T) {
	testCases := []struct {
		testName  string
		userName  string
		expectErr bool
	}{
		{
			testName: "user exists",
			userName: "AshutoshVerma",
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://clevertaptest.trydiscourse.com", "7c27290e2ccca7ae4427dfe518f68fb0659cd7df9908ff1cefbf034f9900a568", "Ashwinigaddagiwork")
 			err := client.DeleteUser(tc.userName)
			log.Println(err)
			if tc.expectErr {
				log.Println("[DELETE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
			_, err = client.GetUser(tc.userName)
			log.Println("[DELETE ERROR]: ", err)
			assert.Error(t, err)
		})
	}
}

