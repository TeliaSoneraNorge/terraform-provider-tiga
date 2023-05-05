# terraform-provider-tiga
A terraform provider for creating Tiga-roles (~security groups in Active Directory)

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

Set environment variables:
```
SIDM_HOST=https://staging.securityservice.teliacompany.com
SIDM_SERVICEID=2bee5a4c-6070-4b37-b13f-689e27a4d2a8
SIDM_SECRET=0b6ad3d80d060fa0c6317673b38cad59d8a249611444ecd63c2919fda9ed358dd89f68b6e5feaa2a522b58a578fc0fa925e5a28c101244cbaa82f90f4540e999

openssl base64 -e <<< 2bee5a4c-6070-4b37-b13f-689e27a4d2a8:0b6ad3d80d060fa0c6317673b38cad59d8a249611444ecd63c2919fda9ed358dd89f68b6e5feaa2a522b58a578fc0fa925e5a28c101244cbaa82f90f4540e

```