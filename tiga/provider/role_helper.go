package tiga

import (
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tct "github.com/telia-company/tiga-go-client/pkg"
)

func createRoleFromPlan(plan *roleResourceModel) *tct.Role {
	newRole := tct.Role{
		Name:               plan.Name.ValueString(),
		Template:           plan.Template.String(),
		ValidFrom:          plan.ValidFrom.ValueString(),
		ValidTo:            plan.ValidTo.ValueString(),
		PreventSelfService: plan.PreventSelfService.ValueBool(),
		Description:        plan.Description.ValueString(),
		SystemInstance:     "/v1/systems/" + strings.ToUpper(plan.HID.ValueString()) + "/instances/" + plan.Instance.ValueString(),
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

	return &newRole
}

func updatePlan(plan *roleResourceModel, role *tct.Role) {
	plan.ID = types.StringValue(role.ID)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	plan.Name = types.StringValue(role.Name)
	plan.ValidFrom = types.StringValue(role.ValidFrom)
	plan.ValidTo = types.StringValue(role.ValidTo)

	// Check this
	plan.PreventSelfService = types.BoolValue(role.PreventSelfService)
	plan.Description = types.StringValue(role.Description)
	// plan.SystemInstance = types.StringValue(role.SystemInstance)
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
