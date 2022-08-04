package launchdarkly

import (
	"context"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Environment Variables
const (
	LAUNCHDARKLY_ACCESS_TOKEN = "LAUNCHDARKLY_ACCESS_TOKEN"
	LAUNCHDARKLY_API_HOST     = "LAUNCHDARKLY_API_HOST"
	LAUNCHDARKLY_OAUTH_TOKEN  = "LAUNCHDARKLY_OAUTH_TOKEN"
)

// Provider keys
const (
	access_token = "access_token"
	oauth_token  = "oauth_token"
	api_host     = "api_host"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			access_token: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(LAUNCHDARKLY_ACCESS_TOKEN, nil),
				Description: "The LaunchDarkly API key",
			},
			oauth_token: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(LAUNCHDARKLY_OAUTH_TOKEN, nil),
				Description: "The LaunchDarkly OAuth token",
			},
			api_host: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(LAUNCHDARKLY_API_HOST, "https://app.launchdarkly.com"),
				Description: "The LaunchDarkly host address, e.g. https://app.launchdarkly.com",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"launchdarkly_team":                      resourceTeam(),
			"launchdarkly_project":                   resourceProject(),
			"launchdarkly_environment":               resourceEnvironment(),
			"launchdarkly_feature_flag":              resourceFeatureFlag(),
			"launchdarkly_webhook":                   resourceWebhook(),
			"launchdarkly_custom_role":               resourceCustomRole(),
			"launchdarkly_segment":                   resourceSegment(),
			"launchdarkly_team_member":               resourceTeamMember(),
			"launchdarkly_feature_flag_environment":  resourceFeatureFlagEnvironment(),
			"launchdarkly_destination":               resourceDestination(),
			"launchdarkly_access_token":              resourceAccessToken(),
			"launchdarkly_flag_trigger":              resourceFlagTrigger(),
			"launchdarkly_audit_log_subscription":    resourceAuditLogSubscription(),
			"launchdarkly_relay_proxy_configuration": resourceRelayProxyConfig(),
			"launchdarkly_metric":                    resourceMetric(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"launchdarkly_team":                      dataSourceTeam(),
			"launchdarkly_team_member":               dataSourceTeamMember(),
			"launchdarkly_team_members":              dataSourceTeamMembers(),
			"launchdarkly_project":                   dataSourceProject(),
			"launchdarkly_environment":               dataSourceEnvironment(),
			"launchdarkly_feature_flag":              dataSourceFeatureFlag(),
			"launchdarkly_feature_flag_environment":  dataSourceFeatureFlagEnvironment(),
			"launchdarkly_webhook":                   dataSourceWebhook(),
			"launchdarkly_segment":                   dataSourceSegment(),
			"launchdarkly_flag_trigger":              dataSourceFlagTrigger(),
			"launchdarkly_audit_log_subscription":    dataSourceAuditLogSubscription(),
			"launchdarkly_relay_proxy_configuration": dataSourceRelayProxyConfig(),
			"launchdarkly_metric":                    dataSourceMetric(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	host := d.Get(api_host).(string)
	if strings.HasPrefix(host, "http") {
		u, _ := url.Parse(host)
		host = u.Host
	}
	accessToken := d.Get(access_token).(string)
	oauthToken := d.Get(oauth_token).(string)

	if oauthToken == "" && accessToken == "" {
		return nil, diag.Errorf("either an %q or %q must be specified", access_token, oauth_token)
	}

	if oauthToken != "" {
		client, err := newClient(oauthToken, host, true)
		if err != nil {
			return client, diag.FromErr(err)
		}
		return client, diags
	}

	client, err := newClient(accessToken, host, false)
	if err != nil {
		return client, diag.FromErr(err)
	}
	return client, diags
}
