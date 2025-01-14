---
page_title: "Provider: prismacloudcompute"
---

# Provider prismacloudcompute

This provider is for the Palo Alto Networks Prisma Cloud platform.

## Example Provider Usage

This code snippet shows how to call the `json` config file:

```hcl
provider "prismacloudcompute" {
  json_config_file = "creds.json"
}
```

Here's what the `creds.json` file might look like:

```json
{
	"url":"localhost",
	"username":"MY_USERNAME",
	"password":"MY_PASSWORD",
	"port":8083,
	"skip_ssl_cert_verification":true
}
```

## Argument Reference

There are multiple ways to specify provider config, and they may all be combined if desired.  The params are taken from the following locations, in order of preference:

1) Any param specified explicitly in the `provider` block
2) From the param's environment variable, where applicable.
3) From the JSON config file, if specified.


The following arguments are supported:

* `url` - (Env: `PRISMACLOUDCOMPUTE_URL`) The API URL without the leading protocol.
* `username` - (Env: `PRISMACLOUDCOMPUTE_USERNAME`) Access key ID.
* `password` - (Env: `PRISMACLOUDCOMPUTE_PASSWORD`) Secret key.
* `customer_name` - (Env: `PRISMACLOUDCOMPUTE_CUSTOMER_NAME`) Customer name.
* `protocol` - (Env: `PRISMACLOUDCOMPUTE_PROTOCOL`) The protocol.  Valid values are `https` or `http`.
* `port` - (Env: `PRISMACLOUDCOMPUTE_PORT`, int) If the port is non-standard for the protocol, the port number to use.
* `timeout` - The default timeout (in seconds) for all communications with Prisma Cloud (default: `90`).
* `logging` - Map of logging options for the API connection.  Valid values are `quiet` (disable logging), `action`, `path`, `send`, and `receive`.
* `disable_reconnect` - (bool) Prisma Cloud invalidates authenticated sessions after 10minutes.  By default the provider will silently get a new JSON web token and continue deploying the plan.  If you do not want the provider to fetch a new JSON web token, set this to `true`.
* `json_web_token` - (Env: `PRISMACLOUDCOMPUTE_JSON_WEB_TOKEN`) A JSON web token.  These are only valid for 10 minutes once issued.  If this is specified but not the `username` / `password` then the provider will not have a way to reauthenticate once the JSON web token expires.
* `json_config_file` - (Env: `PRISMACLOUDCOMPUTE_JSON_CONFIG_FILE`) Retrieve the provider configuration from this JSON file.  When retrieving params from the JSON configuration file, the param names are the same as the provider params, except that underscores in provider params become hyphens in the JSON config file.  For example, the provider param `json_web_token` is `json-web-token` in the config file.

## Support

This template/solution are released under an as-is, best effort, support
policy. These scripts should be seen as community supported and Palo Alto
Networks will contribute our expertise as and when possible. We do not
provide technical support or help in using or troubleshooting the components
of the project through our normal support options such as Palo Alto Networks
support teams, or ASC (Authorized Support Centers) partners and backline
support options. The underlying product used (the VM-Series firewall) by the
scripts or templates are still supported, but the support is only for the
product functionality and not for help in deploying or using the template or
script itself. Unless explicitly tagged, all projects or work posted in our
GitHub repository (at https://github.com/PaloAltoNetworks) or sites other
than our official Downloads page on https://support.paloaltonetworks.com
are provided under the best effort policy.
