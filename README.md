# terraform-provider-tiga
A terraform provider for creating Tiga-roles (~security groups in Active Directory)

To try this provider, follow these steps:

Download the binary for your specific OS.
You can find them in the latest release package.

copy/create the .terraformrc -file into your home directory. 
Change the following line inside the .terraformrc -file: 
'to point to where you downloaded the binary.'

```
 "teliacompany.net/api/tiga" = "/Users/m2s/go/bin"

```

You need to set the proxy to be able to login to SIDM
You need to set the TIGA_HOST for the terraform provider 
You need to set the SIDM_HOST to be able to login
You need to set the SIDM_SERVICEID and SIDM_SECRET to be able to get the JWT token

Therefore, export the following in your terminal:

```
export http_proxy=proxy-se-uan.ddc.teliasonera.net:8080
export HTTP_PROXY=proxy-se-uan.ddc.teliasonera.net:8080
export https_proxy=proxy-se-uan.ddc.teliasonera.net:8080
export HTTPS_PROXY=proxy-se-uan.ddc.teliasonera.net:8080
export TIGA_HOST=https://api.tiga-sandbox.teliacompany.net
export SIDM_HOST=https://staging.securityservice.teliacompany.com
export SIDM_SERVICEID=2bee5a4c-6070-4b37-b13f-689e27a4d2a8
export SIDM_SECRET=0b6ad3d80d060fa0c6317673b38cad59d8a249611444ecd63c2919fda9ed358dd89f68b6e5feaa2a522b58a578fc0fa925e5a28c101244cbaa82f90f4540e999

```

Copy the main.tf -file from this repository in the terraform/create -folder.
You can put that file anywhere you like on your system.

Now, try:

```
terraform plan
```

You should see output like this:

│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - teliacompany.net/api/tiga in /Users/m2s/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become incompatible with published releases.
╵

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # tiga_role.AWS_00000000000004_Administrator_Role will be created
  + resource "tiga_role" "AWS_00000000000004_Administrator_Role" {
      + approval_settings    = {
          + named_approvers              = [
              + "sbh881",
              + "zkv293",
            ]
          + security_clearance_approvers = [
              + "nju840",
              + "zkv293",
            ]
          + skip_manager_approval        = true
          + skip_role_owner_approval     = true
          + skip_system_owner_approval   = true
        }
      + child_roles          = [
        ]
      + description          = "AWS_00000000000004_Administrator_Role"
      + hid                  = "Hid100000006"
      + id                   = (known after apply)
      + instance             = "HID100000006.TEST"
      + last_updated         = (known after apply)
      + name                 = "AWS_00000000000004_Administrator_Role"
      + owners               = [
          + "zkv293",
          + "mdr449",
          + "nju840",
        ]
      + prevent_self_service = false
      + provisioning_type    = "activeDirectory"
      + template             = "Amazon Web Services Cloud (AWS)"
      + user_requirements    = {
          + business_contexts    = [
              + "/v1/businessContexts/Any",
            ]
          + countries            = [
              + "SE",
              + "NO",
              + "DK",
            ]
          + digital_committment  = false
          + terms_and_conditions = "/v1/termsAndConditions/Terms+and+Conditions+Jfrog"
        }
      + valid_from           = "2023-05-03T13:13:13Z"
      + valid_to             = "2024-04-25T13:13:13Z"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + role_admin = {
      + approval_settings    = {
          + named_approvers              = [
              + "sbh881",
              + "zkv293",
            ]
          + security_clearance_approvers = [
              + "nju840",
              + "zkv293",
            ]
          + skip_manager_approval        = true
          + skip_role_owner_approval     = true
          + skip_system_owner_approval   = true
        }
      + child_roles          = []
      + description          = "AWS_00000000000004_Administrator_Role"
      + hid                  = "Hid100000006"
      + id                   = (known after apply)
      + instance             = "HID100000006.TEST"
      + last_updated         = (known after apply)
      + name                 = "AWS_00000000000004_Administrator_Role"
      + owners               = [
          + "zkv293",
          + "mdr449",
          + "nju840",
        ]
      + prevent_self_service = false
      + provisioning_type    = "activeDirectory"
      + template             = "Amazon Web Services Cloud (AWS)"
      + user_requirements    = {
          + business_contexts    = [
              + "/v1/businessContexts/Any",
            ]
          + countries            = [
              + "SE",
              + "NO",
              + "DK",
            ]
          + digital_committment  = false
          + terms_and_conditions = "/v1/termsAndConditions/Terms+and+Conditions+Jfrog"
        }
      + valid_from           = "2023-05-03T13:13:13Z"
      + valid_to             = "2024-04-25T13:13:13Z"
    }

─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions if you run "terraform apply" now.




Login to SIDM and create a token

```
curl --location -g --request POST 'https://staging.securityservice.teliacompany.com/oauth/token' \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Connection: keep-alive' \
--header 'Authorization: Basic MmJlZTVhNGMtNjA3MC00YjM3LWIxM2YtNjg5ZTI3YTRkMmE4OjBiNmFkM2Q4MGQwNjBmYTBjNjMxNzY3M2IzOGNhZDU5ZDhhMjQ5NjExNDQ0ZWNkNjNjMjkxOWZkYTllZDM1OGRkODlmNjhiNmU1ZmVhYTJhNTIyYjU4YTU3OGZjMGZhOTI1ZTVhMjhjMTAxMjQ0Y2JhYTgyZjkwZjQ1NDBlOTk5' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'resource=https://api.tiga-sandbox.teliacompany.net/v1/' \
--data-urlencode 'token_type=jwt'
```

The following environment variables must be set:
Set environment variables:
```
// TIGA_HOST=https://api.tiga-sandbox.teliacompany.net
// For SIDM Login (example values provided)
// SIDM_HOST=https://staging.securityservice.teliacompany.com
// SIDM_SECRET=0b6ad3d80d060fa0c6317673b38cad59d8a249611444ecd63c2919fda9ed358dd89f68b6e5feaa2a522b58a578fc0fa925e5a28c101244cbaa82f90f4540e999
// SIDM_SERVICEID=2bee5a4c-6070-4b37-b13f-689e27a4d2a8

export TIGA_HOST=https://api.tiga-sandbox.teliacompany.net
export SIDM_HOST=https://staging.securityservice.teliacompany.com
export SIDM_SERVICEID=2bee5a4c-6070-4b37-b13f-689e27a4d2a8
export SIDM_SECRET=0b6ad3d80d060fa0c6317673b38cad59d8a249611444ecd63c2919fda9ed358dd89f68b6e5feaa2a522b58a578fc0fa925e5a28c101244cbaa82f90f4540e999

openssl base64 -e <<< 2bee5a4c-6070-4b37-b13f-689e27a4d2a8:0b6ad3d80d060fa0c6317673b38cad59d8a249611444ecd63c2919fda9ed358dd89f68b6e5feaa2a522b58a578fc0fa925e5a28c101244cbaa82f90f4540e

```

Don't forget to put the .terraformrc -file in your home (user) directory.
To activate debug, in the main.go -file, set debug to true.
Start debugging through vscode debugging tools.

