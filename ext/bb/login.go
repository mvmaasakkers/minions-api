package bb

import (
	"github.com/BeyondBankingDays/minions-api"
	"fmt"
	"net/http"
	"encoding/json"
)

type Conn struct {
	Token string
}

func Login(bankUser hackathon_api.BankUser) (*TokenResponse, error) {

	authorizationHeader := fmt.Sprintf(
		`DirectLogin username="%s",password="%s",consumer_key="%s"`,
		bankUser.Username,
		bankUser.Password,
		ConsumerKey,
		)

	client := http.Client{}

	req, err := http.NewRequest("POST", BaseURL+"/my/logins/direct", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", authorizationHeader)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	token := &TokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
		return nil, err
	}

	return token, nil
}

type TokenResponse struct {
	Token string `json:"token"`
}