package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	base_url   string
	api_key       string
	api_username  string
	httpClient *http.Client
}

type ResponseEmail struct {
	Email string `json:"email"` 
}

type User struct {
	Name string `json:"name,omitempty"`
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Admin bool `json:"admin"`
	Active bool `json:"active"`
}

type Response struct {
	Success bool `json:"success"`
}

var (
    Errors = make(map[int]string)
)

func init() {
	Errors[400] = "Bad Request, StatusCode = 400"
	Errors[404] = "User Does Not Exist , StatusCode = 404"
	Errors[409] = "User Already Exist, StatusCode = 409"
	Errors[401] = "Unautharized Access, StatusCode = 401"
	Errors[429] = "User Has Sent Too Many Request, StatusCode = 429"
}

func NewClient(base_url string, api_key string, api_username string) *Client {
	return &Client{
		base_url:   base_url,
		api_key:       api_key,
		api_username:  api_username,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetUsers() ([]User, error){ 
	body, err := c.httpRequest("admin/users/list/active.json?order=created", "GET", &strings.Reader{}) 
	if err != nil {
		log.Println("[READ ERROR]: ",err)
		return nil, err
	}
	response := []User{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("[READ LIST ERROR]: ",err)
		return nil, err
	}
	return response, nil
}

func (c *Client) GetUser(email string) (*User, error){ 
	userList, err := c.GetUsers()
	if err != nil {
		log.Println("[READ ERROR]: ",err)
		return nil, err
	}
	for _,v := range(userList) {
		body, err := c.httpRequest(fmt.Sprintf("u/%v/emails.json", v.Username), "GET", &strings.Reader{})
		if err != nil {
			log.Println("[READ ERROR]: ",err)
			return nil, err
		}
		resp := ResponseEmail{}
		err = json.Unmarshal(body, &resp)
		if err != nil {
			log.Println("[READ ERROR]: ",err)
			return nil, err
		}
		if email == resp.Email {
			v.Email = email 
			return &v, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (c *Client) NewUser(email string) (error) {
	body := fmt.Sprintf("{\"email\":\"%v\"}", email)
	_, err := c.httpRequest("invites.json", "POST", strings.NewReader(body))
	if err != nil {
		log.Println("[CREATE ERROR]: ",err)
		return err
	}
	return nil 
}

func (c *Client) DeactivateUser(userid int) (error){ 
	_, err := c.httpRequest(fmt.Sprintf("admin/users/%v/deactivate.json", userid), "PUT", &strings.Reader{}) 
	if err != nil {
		log.Println("[ERROR]: ",err)
		return err
	}
	return nil
}

func (c *Client) ActivateUser(userid int) (error){ 
	_, err := c.httpRequest(fmt.Sprintf("admin/users/%v/activate.json", userid), "PUT", &strings.Reader{}) 
	if err != nil {
		log.Println("[ERROR]: ",err)
		return err
	}
	return nil
}
 
func (c *Client) httpRequest(path, method string, body *strings.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, c.requestPath(path), body)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	}
	req.Header.Add("Api-Key", c.api_key)
	req.Header.Add("Api-Username", c.api_username)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	} 
	if resp.StatusCode != 200 {
		fmt.Errorf("Error : %v",Errors[resp.StatusCode])
		return nil, fmt.Errorf(string(respBody))
	} else {
		return respBody, nil
	}
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%v/%v", c.base_url, path)
}

func (c *Client) DeleteUser(username string) error { 
	_, err := c.httpRequest(fmt.Sprintf("u/%v.json", username), "DELETE", &strings.Reader{}) 
	if err != nil {
		log.Println("[ERROR]: ",err)
		return err
	}
	return nil
}

func (c *Client) UpdateUser(user *User) error { 
	body, err := json.Marshal(user)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("u/%v.json", user.Username), "PUT", strings.NewReader(string(body))) 
	if err != nil {
		log.Println("[ERROR]: ",err)
		return err
	}
	return nil
}
