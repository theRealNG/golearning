package reqres

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Create(user *User) {

	requestParams, err := json.Marshal(*user)

	resp, err := http.Post(CreateUserURL, "application/octet-stream", bytes.NewBuffer(requestParams))
	if err != nil {
		panic(fmt.Sprintf("Failed to make post request: %v", err))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed to read response body: %v", err))
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal the response body: %v", err))
	}
}
