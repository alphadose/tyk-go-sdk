package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "github.com/alphadose/tyk-go-sdk"
)

const secretKey = "352d20ee67be67f6340b4c0605b044b7"

const defaultEndpoint = "http://localhost:8080"

func main() {
	aPIDefinition := *openapiclient.NewAPIDefinition() // APIDefinition |  (optional)
	ctx := context.WithValue(context.Background(), openapiclient.ContextAPIKeys, map[string]openapiclient.APIKey{
		"api_key": {Key: secretKey},
	})
	aPIDefinition.SetName("TestAPI")
	aPIDefinition.SetUseKeyless(true)
	aPIDefinition.SetActive(true)

	proxy := *openapiclient.NewAPIDefinitionProxy()
	proxy.SetListenPath("/test")

	aPIDefinition.SetProxy(proxy)

	configuration := &openapiclient.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: openapiclient.ServerConfigurations{
			{
				URL:         defaultEndpoint,
				Description: "Tyk CLI POC",
			},
		},
		OperationServers: map[string]openapiclient.ServerConfigurations{},
	}
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIsApi.CreateApi(ctx).APIDefinition(aPIDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIsApi.CreateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r.Body)
	}
	// response from `CreateApi`: ApiModifyKeySuccess
	fmt.Fprintf(os.Stdout, "Response from `APIsApi.CreateApi`: %v\n", resp)
	a, _ := resp.MarshalJSON()
	fmt.Printf("%s\n", string(a))
}
