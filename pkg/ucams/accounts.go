package ucams

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	accountsFile = "./data/accounts.json"
)

type (
	Account struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}
)

func (c *UcamsClient) AddAccount(username string, password string) error {
	if _, exists := c.accounts[username]; exists {
		return fmt.Errorf("Failed to add account %q: Already exist", username)
	}

	authResult, err := c.apiClient.Auth(username, password)
	if err != nil {
		return fmt.Errorf("Failed to add account %q: %v", username, err)
	}

	account := Account{
		Username: username,
		Password: password,
		Token:    authResult.Token,
	}

	c.accounts[username] = account
	c.activeAccount = account
	c.saveAccounts()

	return nil
}

func (c *UcamsClient) saveAccounts() {
	jsonBytes, _ := json.Marshal(c.accounts)
	writeJSONToFile(accountsFile, jsonBytes)
}

func (c *UcamsClient) readAccounts() error {
	var accounts map[string]Account
	jsonBytes, _ := readJSONFile(accountsFile)
	// Unmarshal the JSON data into the struct
	err := json.Unmarshal(jsonBytes, &accounts)
	if err != nil {
		return err
	}

	c.accounts = accounts

	return nil
}

// TODO: move this into utils
func writeJSONToFile(filename string, data []byte) error {
	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func readJSONFile(filename string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
