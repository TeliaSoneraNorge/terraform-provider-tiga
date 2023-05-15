package main

import (
	"context"

	tiga "terraform-provider-tiga/tiga/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name telia-tiga

func main() {
	providerserver.Serve(context.Background(), tiga.New, providerserver.ServeOpts{
		// NOTE: This is not a typical Terraform Registry provider address,
		// such as registry.terraform.io/hashicorp/tiga. This specific
		// provider address is used in these tutorials in conjunction with a
		// specific Terraform CLI configuration for manual development testing
		// of this provider.
		Debug:   false,
		Address: "teliacompany.net/api/tiga",
	})
}
