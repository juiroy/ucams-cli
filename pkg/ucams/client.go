package ucams

import ucams_api "github.com/juiroy/ucams-cli/pkg/ucams/api"

type (
	UcamsClient struct {
		apiClient     *ucams_api.UcamsApi
		accounts      map[string]Account
		activeAccount Account
	}
)

func Create() *UcamsClient {
	apiClient := ucams_api.Create("https://ucams.ufanet.ru")

	client := UcamsClient{
		apiClient: apiClient,
		accounts:  make(map[string]Account),
	}

	client.readAccounts()

	return &client
}
