package tiga

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tct "github.com/telia-company/tiga-go-client/pkg"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &roleResource{}
	_ resource.ResourceWithConfigure = &roleResource{}
)

// NewRoleResource is a helper function to simplify the provider implementation.
func NewRoleResource() resource.Resource {
	return &roleResource{}
}

// roleResource is the resource implementation.
type roleResource struct {
	client *tct.Client
}

// roleResourceModel maps the resource schema data.
type roleResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	LastUpdated        types.String `tfsdk:"last_updated"`
	HID                types.String `tfsdk:"hid"`
	Instance           types.String `tfsdk:"instance"`
	Name               types.String `tfsdk:"name"`
	Template           types.String `tfsdk:"template"`
	ValidFrom          types.String `tfsdk:"valid_from"`
	ValidTo            types.String `tfsdk:"valid_to"`
	PreventSelfService types.Bool   `tfsdk:"prevent_self_service"`
	Description        types.String `tfsdk:"description"`
	// SystemInstance     types.String         `tfsdk:"system_instance"`
	ProvisioningType types.String         `tfsdk:"provisioning_type"`
	ApprovalSettings roleApprovalSettings `tfsdk:"approval_settings"`
	Owners           []types.String       `tfsdk:"owners"`
	UserRequirements roleUserRequirements `tfsdk:"user_requirements"`
	ChildRoles       []types.String       `tfsdk:"child_roles"`
}

type roleApprovalSettings struct {
	SkipSystemOwnerApproval    types.Bool     `tfsdk:"skip_system_owner_approval"`
	SkipManagerApproval        types.Bool     `tfsdk:"skip_manager_approval"`
	SkipRoleOwnerApproval      types.Bool     `tfsdk:"skip_role_owner_approval"`
	NamedApprovers             []types.String `tfsdk:"named_approvers"`
	SecurityClearanceApprovers []types.String `tfsdk:"security_clearance_approvers"`
}

type roleUserRequirements struct {
	DigitalCommittment types.Bool     `tfsdk:"digital_committment"`
	TermsAndConditions types.String   `tfsdk:"terms_and_conditions"`
	Countries          []types.String `tfsdk:"countries"`
	BusinessContexts   []types.String `tfsdk:"business_contexts"`
}

// Metadata returns the data source type name.
func (r *roleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_role"
}

// Schema defines the schema for the data source.
func (r *roleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a role.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of the Tiga role.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Description: "Timestamp of the last Terraform update of the role.",
				Computed:    true,
			},
			"hid": schema.StringAttribute{
				Description: "HID of the system.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"instance": schema.StringAttribute{
				Description: "System instance.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the role.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"template": schema.StringAttribute{
				Description: "Name of the template.",
				Optional:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"valid_from": schema.StringAttribute{
				Description: "A Simple Date in string format '2023-04-25' indicating when the role is valid from.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"valid_to": schema.StringAttribute{
				Description: "A Simple Date in string format '2023-04-25' indicating the date the role is valid up until.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"prevent_self_service": schema.BoolAttribute{
				Required:    true,
				Computed:    false,
				Description: "Prevent self service of the role.",
			},
			"description": schema.StringAttribute{
				Description: "The description of the role.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"provisioning_type": schema.StringAttribute{
				Description: "The provisioning type for the role.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"approval_settings": schema.ObjectAttribute{
				Description: "Approval settings for the role.",
				Required:    true,
				AttributeTypes: map[string]attr.Type{
					"skip_system_owner_approval":   types.BoolType,
					"skip_manager_approval":        types.BoolType,
					"skip_role_owner_approval":     types.BoolType,
					"named_approvers":              types.ListType{ElemType: types.StringType},
					"security_clearance_approvers": types.ListType{ElemType: types.StringType},
				},
			},
			"owners": schema.ListAttribute{
				Description: "List of owners for the role.",
				Required:    true,
				ElementType: types.StringType,
			},
			"user_requirements": schema.ObjectAttribute{
				Description: "User requirements for the role.",
				Required:    true,
				AttributeTypes: map[string]attr.Type{
					"digital_committment":  types.BoolType,
					"terms_and_conditions": types.StringType,
					"business_contexts":    types.ListType{ElemType: types.StringType},
					"countries":            types.ListType{ElemType: types.StringType},
				},
			},
			"child_roles": schema.ListNestedAttribute{
				Description: "List of child roles for the role.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"child_role": schema.StringAttribute{
							Description: "The child role.",
							Optional:    true,
							Computed:    false,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *roleResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*tct.Client)
}

// Create creates the resource and sets the initial Terraform state.
func (r *roleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan roleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newRole := createRoleFromPlan(&plan)

	role, err := r.client.CreateRole(newRole)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"Could not create Role, unexpected error: "+err.Error(),
		)
		return
	}

	updatePlan(&plan, role)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *roleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state roleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	role, err := r.client.GetRole(state.HID.ValueString(), state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Tiga Role",
			"Could not read Tiga Role Name "+state.Name.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state.Name = types.StringValue(role.Name)
	state.ID = types.StringValue(role.ID)
	state.ValidFrom = types.StringValue(role.ValidFrom)
	state.ValidTo = types.StringValue(role.ValidTo)
	state.PreventSelfService = types.BoolValue(role.PreventSelfService)
	state.Description = types.StringValue(role.Description)
	state.ProvisioningType = types.StringValue(role.ProvisioningType)
	state.ApprovalSettings.SkipSystemOwnerApproval = types.BoolValue(role.ApprovalSettings.SkipSystemOwnerApproval)
	state.ApprovalSettings.SkipManagerApproval = types.BoolValue(role.ApprovalSettings.SkipManagerApproval)
	state.ApprovalSettings.SkipRoleOwnerApproval = types.BoolValue(role.ApprovalSettings.SkipRoleOwnerApproval)
	state.UserRequirements.DigitalCommittment = types.BoolValue(role.UserRequirements.DigitalCommittment)
	state.UserRequirements.TermsAndConditions = types.StringValue(role.UserRequirements.TermsAndConditions)
	state.ApprovalSettings.NamedApprovers = compareSlices(state.ApprovalSettings.NamedApprovers, role.ApprovalSettings.NamedApprovers)
	state.ApprovalSettings.SecurityClearanceApprovers = compareSlices(state.ApprovalSettings.SecurityClearanceApprovers, role.ApprovalSettings.SecurityClearanceApprovers)
	state.Owners = compareSlices(state.Owners, role.Owners)
	state.UserRequirements.Countries = compareSlices(state.UserRequirements.Countries, role.UserRequirements.Countries)
	state.UserRequirements.BusinessContexts = compareSlices(state.UserRequirements.BusinessContexts, role.UserRequirements.BusinessContexts)
	state.ChildRoles = compareSlices(state.ChildRoles, role.ChildRoles)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *roleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var state roleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError(
		"Tiga - Updating Role",
		"Updating Tiga resources (roles) are not supported through this terraform provider at the moment, however, updates "+
			"might be made through the Tiga web interface at https://tiga.teliacompany.net/workitemdlg.aspx?ACTTEMP=1003274",
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *roleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state roleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError(
		"Tiga - Deleting Role",
		"Deleting Tiga resources (roles) is not possible, however, disabling of resources (roles) "+
			"might be made through the Tiga web interface at https://tiga.teliacompany.net/workitemdlg.aspx?ACTTEMP=1003274",
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
