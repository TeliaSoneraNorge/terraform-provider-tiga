package tiga

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	tct "github.com/telia-company/tiga-go-client/pkg"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &tigaProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &tigaProvider{}
}

// tigaProvider is the provider implementation.
type tigaProvider struct{}

// tigaProviderModel maps provider schema data to a Go type.
type tigaProviderModel struct {
	Host               types.String `tfsdk:"host"`
	TermsAndConditions types.Bool   `tfsdk:"we_agree_to_terms_and_conditions"`
}

// Metadata returns the provider type name.
func (p *tigaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tiga"
}

// Schema defines the provider-level schema for configuration data.
func (p *tigaProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Tiga.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "URI for Tiga API. May also be provided via TIGA_HOST environment variable.",
				Optional:    true,
			},
			"we_agree_to_terms_and_conditions": schema.BoolAttribute{
				Description: "You confirm that you agree to the terms and conditions to use Tiga to create 'resources' (roles)\n" +
					"you can read more about it here: https://itwiki.atlassian.teliacompany.net/display/TIGA/Digital+Commitment",
				Required: true,
			},
		},
	}
}

// Configure prepares a Tiga API client for resources.
func (p *tigaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Tiga client")

	// Retrieve provider data from configuration
	var config tigaProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Tiga API Host",
			"The provider cannot create the Tiga API client as there is an unknown configuration value for the Tiga API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TIGA_HOST environment variable.",
		)
	}

	if !config.TermsAndConditions.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			path.Root("we_agree_to_terms_and_conditions"),
			"Must agree to terms and conditions",
			"You must confirm that you agree to the terms and conditions to use Tiga to create 'resources' (roles). "+
				"you can read more about it here: https://itwiki.atlassian.teliacompany.net/display/TIGA/Digital+Commitment",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("TIGA_HOST")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Tiga API Host",
			"The provider cannot create the Tiga API client as there is a missing or empty value for the Tiga API host. "+
				"Set the host value in the configuration or use the TIGA_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "tiga_host", host)

	tflog.Debug(ctx, "Creating Tiga client")

	// Create a new Tiga client using the configuration values
	client, err := tct.New(&tct.Caller{}, true)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Tiga API Client",
			"An unexpected error occurred when creating the Tiga API client. \n\n"+
				"Tiga Client Error: "+err.Error(),
		)
		return
	}

	// Make the Tiga client available during Resource
	// type Configure methods.
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Tiga client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *tigaProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Resources defines the resources implemented in the provider.
func (p *tigaProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewRoleResource,
	}
}
