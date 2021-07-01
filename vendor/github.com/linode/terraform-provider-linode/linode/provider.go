package azurerm

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/azurerm/azurermgo"
)

// Provider creates and manages the resources in a Linode configuration.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_TOKEN", nil),
				Description: "The token that allows you access to your Linode account",
			},
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("LINODE_URL", nil),
				Description:  "The HTTP(S) API address of the Linode API to use.",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"ua_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_UA_PREFIX", nil),
				Description: "An HTTP User-Agent Prefix to prepend in API requests.",
			},
			"api_version": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_API_VERSION", nil),
				Description: "An HTTP User-Agent Prefix to prepend in API requests.",
			},

			"skip_instance_ready_poll": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Skip waiting for a azurerm_instance resource to be running.",
			},

			"skip_instance_delete_poll": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Skip waiting for a azurerm_instance resource to finish deleting.",
			},

			"min_retry_delay_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Minimum delay in milliseconds before retrying a request.",
			},
			"max_retry_delay_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum delay in milliseconds before retrying a request.",
			},

			"event_poll_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_EVENT_POLL_MS", 300),
				Description: "The rate in milliseconds to poll for events.",
			},

			"lke_event_poll_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     300,
				Description: "The rate in milliseconds to poll for LKE events.",
			},

			"lke_node_ready_poll_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     500,
				Description: "The rate in milliseconds to poll for an LKE node to be ready.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"azurerm_account":                dataSourceLinodeAccount(),
			"azurerm_domain":                 dataSourceLinodeDomain(),
			"azurerm_domain_record":          dataSourceLinodeDomainRecord(),
			"azurerm_firewall":               dataSourceLinodeFirewall(),
			"azurerm_image":                  dataSourceLinodeImage(),
			"azurerm_images":                 dataSourceLinodeImages(),
			"azurerm_instances":              dataSourceLinodeInstances(),
			"azurerm_instance_backups":       dataSourceLinodeInstanceBackups(),
			"azurerm_instance_type":          dataSourceLinodeInstanceType(),
			"azurerm_kernel":                 dataSourceLinodeKernel(),
			"azurerm_lke_cluster":            dataSourceLinodeLKECluster(),
			"azurerm_networking_ip":          dataSourceLinodeNetworkingIP(),
			"azurerm_nodebalancer":           dataSourceLinodeNodeBalancer(),
			"azurerm_nodebalancer_config":    dataSourceLinodeNodeBalancerConfig(),
			"azurerm_nodebalancer_node":      dataSourceLinodeNodeBalancerNode(),
			"azurerm_object_storage_cluster": dataSourceLinodeObjectStorageCluster(),
			"azurerm_profile":                dataSourceLinodeProfile(),
			"azurerm_region":                 dataSourceLinodeRegion(),
			"azurerm_sshkey":                 dataSourceLinodeSSHKey(),
			"azurerm_stackscript":            dataSourceLinodeStackscript(),
			"azurerm_user":                   dataSourceLinodeUser(),
			"azurerm_vlans":                  dataSourceLinodeVLANs(),
			"azurerm_volume":                 dataSourceLinodeVolume(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"azurerm_domain":                resourceLinodeDomain(),
			"azurerm_domain_record":         resourceLinodeDomainRecord(),
			"azurerm_firewall":              resourceLinodeFirewall(),
			"azurerm_image":                 resourceLinodeImage(),
			"azurerm_instance":              resourceLinodeInstance(),
			"azurerm_instance_ip":           resourceLinodeInstanceIP(),
			"azurerm_lke_cluster":           resourceLinodeLKECluster(),
			"azurerm_nodebalancer":          resourceLinodeNodeBalancer(),
			"azurerm_nodebalancer_config":   resourceLinodeNodeBalancerConfig(),
			"azurerm_nodebalancer_node":     resourceLinodeNodeBalancerNode(),
			"azurerm_object_storage_bucket": resourceLinodeObjectStorageBucket(),
			"azurerm_object_storage_key":    resourceLinodeObjectStorageKey(),
			"azurerm_object_storage_object": resourceLinodeObjectStorageObject(),
			"azurerm_rdns":                  resourceLinodeRDNS(),
			"azurerm_sshkey":                resourceLinodeSSHKey(),
			"azurerm_stackscript":           resourceLinodeStackscript(),
			"azurerm_token":                 resourceLinodeToken(),
			"azurerm_user":                  resourceLinodeUser(),
			"azurerm_volume":                resourceLinodeVolume(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}
	return provider
}

type ProviderMeta struct {
	Client azurermgo.Client
	Config *Config
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := &Config{
		AccessToken: d.Get("token").(string),
		APIURL:      d.Get("url").(string),
		APIVersion:  d.Get("api_version").(string),
		UAPrefix:    d.Get("ua_prefix").(string),

		SkipInstanceReadyPoll:  d.Get("skip_instance_ready_poll").(bool),
		SkipInstanceDeletePoll: d.Get("skip_instance_delete_poll").(bool),

		MinRetryDelayMilliseconds: d.Get("min_retry_delay_ms").(int),
		MaxRetryDelayMilliseconds: d.Get("max_retry_delay_ms").(int),

		EventPollMilliseconds:    d.Get("event_poll_ms").(int),
		LKEEventPollMilliseconds: d.Get("lke_event_poll_ms").(int),

		LKENodeReadyPollMilliseconds: d.Get("lke_node_ready_poll_ms").(int),
	}
	config.terraformVersion = terraformVersion
	client := config.Client()

	// Ping the API for an empty response to verify the configuration works
	if _, err := client.ListTypes(context.Background(), azurermgo.NewListOptions(100, "")); err != nil {
		return nil, fmt.Errorf("Error connecting to the Linode API: %s", err)
	}
	return &ProviderMeta{
		Client: client,
		Config: config,
	}, nil
}
