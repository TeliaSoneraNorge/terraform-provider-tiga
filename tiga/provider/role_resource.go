package tiga

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tct "github.com/telia-company/tiga-go-client/pkg"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &roleResource{}
	_ resource.ResourceWithConfigure   = &roleResource{}
	_ resource.ResourceWithImportState = &roleResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewRoleResource() resource.Resource {
	return &roleResource{}
}

// orderResource is the resource implementation.
type roleResource struct {
	//client *Client
	client *tct.Client
}

// roleResourceModel maps the resource schema data.
type roleResourceModel struct {
	ID                 types.String         `tfsdk:"id"`
	LastUpdated        types.String         `tfsdk:"last_updated"`
	HID                types.String         `tfsdk:"hid"`
	Instance           types.String         `tfsdk:"instance"`
	Name               types.String         `tfsdk:"name"`
	Template           types.String         `tfsdk:"template"`
	ValidFrom          types.String         `tfsdk:"valid_from"`
	ValidTo            types.String         `tfsdk:"valid_to"`
	PreventSelfService types.Bool           `tfsdk:"prevent_self_service"`
	Description        types.String         `tfsdk:"description"`
	SystemInstance     types.String         `tfsdk:"system_instance"`
	ProvisioningType   types.String         `tfsdk:"provisioning_type"`
	ApprovalSettings   roleApprovalSettings `tfsdk:"approval_settings"`
	Owners             []types.String       `tfsdk:"owners"`
	UserRequirements   roleUserRequirements `tfsdk:"user_requirements"`
	ChildRoles         []types.String       `tfsdk:"child_roles"`
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
	resp.TypeName = req.ProviderTypeName + "_roleresource"
}

// Schema defines the schema for the data source.
func (r *roleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a role.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Numeric identifier of the order.",
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
				Optional:    false,
				Computed:    false,
				Description: "Prevent self service of the role.",
			},
			"description": schema.StringAttribute{
				Description: "The role description.",
				Required:    true,
				Computed:    false,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"system_instance": schema.StringAttribute{
				Description: "The system instance.",
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

	//Create new role
	newRole := tct.Role{
		// hid:                plan.HID.ValueString(),
		// instance:           plan.Instance.ValueString(),
		Name:      plan.Name.ValueString(),
		Template:  plan.Template.String(),
		ValidFrom: plan.ValidFrom.ValueString(),
		// otherBool:          false,
		ValidTo:            plan.ValidTo.ValueString(),
		PreventSelfService: plan.PreventSelfService.ValueBool(),
		Description:        plan.Description.ValueString(),
		SystemInstance:     plan.SystemInstance.ValueString(),
		ProvisioningType:   plan.ProvisioningType.ValueString(),
		ApprovalSettings: tct.ApprovalSettings{
			SkipSystemOwnerApproval: plan.ApprovalSettings.SkipSystemOwnerApproval.ValueBool(),
			SkipManagerApproval:     plan.ApprovalSettings.SkipManagerApproval.ValueBool(),
			SkipRoleOwnerApproval:   plan.ApprovalSettings.SkipRoleOwnerApproval.ValueBool(),
		},
		UserRequirements: tct.UserRequirements{
			DigitalCommittment: plan.UserRequirements.DigitalCommittment.ValueBool(),
			TermsAndConditions: plan.UserRequirements.TermsAndConditions.ValueString(),
		},
	}

	//newRole.otherBool = false
	for _, namedApprover := range plan.ApprovalSettings.NamedApprovers {
		newRole.ApprovalSettings.NamedApprovers = append(newRole.ApprovalSettings.NamedApprovers, namedApprover.ValueString())
	}
	for _, securityClearanceApprover := range plan.ApprovalSettings.SecurityClearanceApprovers {
		newRole.ApprovalSettings.SecurityClearanceApprovers = append(newRole.ApprovalSettings.SecurityClearanceApprovers, securityClearanceApprover.ValueString())
	}
	for _, owner := range plan.Owners {
		newRole.Owners = append(newRole.Owners, owner.ValueString())
	}

	for _, country := range plan.UserRequirements.Countries {
		newRole.UserRequirements.Countries = append(newRole.UserRequirements.Countries, country.ValueString())
	}

	for _, businessContext := range plan.UserRequirements.BusinessContexts {
		newRole.UserRequirements.BusinessContexts = append(newRole.UserRequirements.BusinessContexts, businessContext.ValueString())
	}
	for _, childRole := range plan.ChildRoles {
		childRoles := []string{}
		childRoles = append(childRoles, childRole.ValueString())
		newRole.ChildRoles = childRoles
	}

	role, err := r.client.CreateRole(&newRole)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"Could not create order, unexpected error: "+err.Error(),
		)
		return
	}

	// role.hid = plan.HID.ValueString()
	// role.instance = plan.Instance.ValueString()

	plan.ID = types.StringValue(plan.Instance.ValueString() + "/" + role.Name)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	plan.Name = types.StringValue(role.Name)
	plan.ValidFrom = types.StringValue(role.ValidFrom)
	plan.ValidTo = types.StringValue(role.ValidTo)

	// Check this
	plan.PreventSelfService = types.BoolValue(role.PreventSelfService)
	plan.Description = types.StringValue(role.Description)
	plan.SystemInstance = types.StringValue(role.SystemInstance)
	plan.ProvisioningType = types.StringValue(role.ProvisioningType)
	plan.ApprovalSettings.SkipSystemOwnerApproval = types.BoolValue(role.ApprovalSettings.SkipSystemOwnerApproval)
	plan.ApprovalSettings.SkipManagerApproval = types.BoolValue(role.ApprovalSettings.SkipManagerApproval)
	plan.ApprovalSettings.SkipRoleOwnerApproval = types.BoolValue(role.ApprovalSettings.SkipRoleOwnerApproval)

	//Check this
	plan.UserRequirements.DigitalCommittment = types.BoolValue(role.UserRequirements.DigitalCommittment)
	plan.UserRequirements.TermsAndConditions = types.StringValue(role.UserRequirements.TermsAndConditions)
	plan.ApprovalSettings.NamedApprovers = compareSlices(plan.ApprovalSettings.NamedApprovers, role.ApprovalSettings.NamedApprovers)
	plan.ApprovalSettings.SecurityClearanceApprovers = compareSlices(plan.ApprovalSettings.SecurityClearanceApprovers, role.ApprovalSettings.SecurityClearanceApprovers)
	plan.Owners = compareSlices(plan.Owners, role.Owners)
	plan.UserRequirements.Countries = compareSlices(plan.UserRequirements.Countries, role.UserRequirements.Countries)
	plan.UserRequirements.BusinessContexts = compareSlices(plan.UserRequirements.BusinessContexts, role.UserRequirements.BusinessContexts)
	plan.ChildRoles = compareSlices(plan.ChildRoles, role.ChildRoles)

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
	// role.hid = state.HID.ValueString()
	// role.instance = state.Instance.ValueString()

	// state.HID = types.StringValue(role.hid)
	// state.Instance = types.StringValue(role.instance)
	state.ValidFrom = types.StringValue(role.ValidFrom)
	state.ValidTo = types.StringValue(role.ValidTo)
	state.PreventSelfService = types.BoolValue(role.PreventSelfService)
	state.Description = types.StringValue(role.Description)
	state.SystemInstance = types.StringValue(role.SystemInstance)
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
			"can be made through the Tiga web interface at https://tiga.teliacompany.net/workitemdlg.aspx?ACTTEMP=1003274",
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
			"can be made through the Tiga web interface at https://tiga.teliacompany.net/workitemdlg.aspx?ACTTEMP=1003274",
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *roleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// compareSlices - compare the TF Plan/State with Role info from Tiga
func compareSlices(s1 []basetypes.StringValue, s2 []string) []basetypes.StringValue {
	// create a map to store the strings in s2
	s2Map := make(map[basetypes.StringValue]bool)
	for _, str := range s2 {
		s2Map[types.StringValue(str)] = true
	}

	// create a new slice to store the updated strings
	updatedSlice := make([]basetypes.StringValue, 0)

	// iterate over the strings in s1 and check if they're in s2
	for _, str := range s1 {
		if _, ok := s2Map[str]; ok {
			// string is in s2, add it to the updated slice
			updatedSlice = append(updatedSlice, str)
			// remove the string from the s2 map to avoid duplicates
			delete(s2Map, str)
		}
	}

	// iterate over the remaining strings in s2 and add them to the updated slice
	for str := range s2Map {
		updatedSlice = append(updatedSlice, str)
	}

	return updatedSlice
}
