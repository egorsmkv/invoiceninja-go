package main

import (
	"context"

	invoiceninja "github.com/egorsmkv/invoiceninja-go"
)

func main() {
	ctx := context.WithValue(context.Background(), invoiceninja.ContextServerIndex, 0)

	cfg := &invoiceninja.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         true,
		Servers: invoiceninja.ServerConfigurations{
			{
				URL:         "https://demo.invoiceninja.com",
				Description: "## Demo API endpoint  You can use the demo API key `TOKEN` to test the endpoints from within this API spec ",
			},
		},
		OperationServers: map[string]invoiceninja.ServerConfigurations{},
	}

	client := invoiceninja.NewAPIClient(cfg)

	req := client.AuthAPI.PostLogin(ctx)

	req = req.XRequestedWith("XMLHttpRequest")
	req = req.XAPITOKEN("2Rds7DiyKX52V2PJQBLsneix6va01KcNoBZ804KXII5Vo25yd9JTQ0N3mcW5IWgy")
	req = req.PostLoginRequest(invoiceninja.PostLoginRequest{
		Email:    "demo@invoiceninja.com",
		Password: "Password0",
	})

	user, response, err := client.AuthAPI.PostLoginExecute(req)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		panic("Error: " + response.Status)
	}

	if user == nil {
		panic("Error: user is nil")
	}
}
