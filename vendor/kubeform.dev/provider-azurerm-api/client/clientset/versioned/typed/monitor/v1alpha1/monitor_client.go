/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/monitor/v1alpha1"
	"kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type MonitorV1alpha1Interface interface {
	RESTClient() rest.Interface
	AadDiagnosticSettingsGetter
	ActionGroupsGetter
	ActionRuleActionGroupsGetter
	ActionRuleSuppressionsGetter
	ActivityLogAlertsGetter
	AutoscaleSettingsGetter
	DiagnosticSettingsGetter
	LogProfilesGetter
	MetricAlertsGetter
	PrivateLinkScopesGetter
	PrivateLinkScopedServicesGetter
	ScheduledQueryRulesAlertsGetter
	ScheduledQueryRulesLogsGetter
	SmartDetectorAlertRulesGetter
}

// MonitorV1alpha1Client is used to interact with features provided by the monitor.azurerm.kubeform.com group.
type MonitorV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MonitorV1alpha1Client) AadDiagnosticSettings(namespace string) AadDiagnosticSettingInterface {
	return newAadDiagnosticSettings(c, namespace)
}

func (c *MonitorV1alpha1Client) ActionGroups(namespace string) ActionGroupInterface {
	return newActionGroups(c, namespace)
}

func (c *MonitorV1alpha1Client) ActionRuleActionGroups(namespace string) ActionRuleActionGroupInterface {
	return newActionRuleActionGroups(c, namespace)
}

func (c *MonitorV1alpha1Client) ActionRuleSuppressions(namespace string) ActionRuleSuppressionInterface {
	return newActionRuleSuppressions(c, namespace)
}

func (c *MonitorV1alpha1Client) ActivityLogAlerts(namespace string) ActivityLogAlertInterface {
	return newActivityLogAlerts(c, namespace)
}

func (c *MonitorV1alpha1Client) AutoscaleSettings(namespace string) AutoscaleSettingInterface {
	return newAutoscaleSettings(c, namespace)
}

func (c *MonitorV1alpha1Client) DiagnosticSettings(namespace string) DiagnosticSettingInterface {
	return newDiagnosticSettings(c, namespace)
}

func (c *MonitorV1alpha1Client) LogProfiles(namespace string) LogProfileInterface {
	return newLogProfiles(c, namespace)
}

func (c *MonitorV1alpha1Client) MetricAlerts(namespace string) MetricAlertInterface {
	return newMetricAlerts(c, namespace)
}

func (c *MonitorV1alpha1Client) PrivateLinkScopes(namespace string) PrivateLinkScopeInterface {
	return newPrivateLinkScopes(c, namespace)
}

func (c *MonitorV1alpha1Client) PrivateLinkScopedServices(namespace string) PrivateLinkScopedServiceInterface {
	return newPrivateLinkScopedServices(c, namespace)
}

func (c *MonitorV1alpha1Client) ScheduledQueryRulesAlerts(namespace string) ScheduledQueryRulesAlertInterface {
	return newScheduledQueryRulesAlerts(c, namespace)
}

func (c *MonitorV1alpha1Client) ScheduledQueryRulesLogs(namespace string) ScheduledQueryRulesLogInterface {
	return newScheduledQueryRulesLogs(c, namespace)
}

func (c *MonitorV1alpha1Client) SmartDetectorAlertRules(namespace string) SmartDetectorAlertRuleInterface {
	return newSmartDetectorAlertRules(c, namespace)
}

// NewForConfig creates a new MonitorV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MonitorV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MonitorV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MonitorV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MonitorV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MonitorV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MonitorV1alpha1Client {
	return &MonitorV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MonitorV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
