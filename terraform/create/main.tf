terraform {
  required_providers {
    tiga = {
      source = "teliacompany.net/api/tiga"
    }
  }
  required_version = ">= 1.1.0"
}

provider "tiga" {
  # host = "https://api.tiga-sandbox.teliacompany.net"
  we_agree_to_terms_and_conditions = true
}

resource "tiga_roleresource" "AWS_999999990100_Administrator_Role" {
  
    hid = "Hid100000006"
    instance = "HID100000006.TEST"
    name = "AWS_00000000000002_Administrator_Role"
    template = "Amazon Web Services Cloud (AWS)"
    valid_from = "2023-05-03T13:13:13Z"
    valid_to = "2024-04-25T13:13:13Z"
    prevent_self_service = false
    description = "AWS_999999990100_Administrator"
    provisioning_type = "activeDirectory"

    owners = [
      "zkv293", "mdr449", "nju840"
    ]

    approval_settings = {
        
        skip_system_owner_approval = true
        skip_manager_approval = true
        skip_role_owner_approval = true

        named_approvers = [
          "sbh881", "zkv293"
        ]
        
        security_clearance_approvers = [
          "nju840", "zkv293"
        ]
        
    }

    user_requirements = {
        digital_committment = false
        terms_and_conditions = "/v1/termsAndConditions/Terms+and+Conditions+Jfrog"
        countries = [
          "SE","NO","DK"
        ]
        business_contexts = [
          "/v1/businessContexts/Any"
        ]
    }

    child_roles = [

    ]
}

output "role_admin" {
  value = tiga_roleresource.AWS_999999990100_Administrator_Role
}
