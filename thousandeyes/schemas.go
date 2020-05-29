package thousandeyes

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

var schemas = map[string]*schema.Schema{
	"agents": {
		Type:        schema.TypeList,
		Description: "agents to use ",
		Required:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"agent_id": {
					Type:        schema.TypeInt,
					Description: "agent id",
					Optional:    true,
				},
			},
		},
	},
	"agents--label": {
		Type:        schema.TypeList,
		Description: "agents to use ",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"agent_id": {
					Type:        schema.TypeInt,
					Description: "agent id",
					Optional:    true,
				},
			},
		},
	},
	"alert_rules": {
		Description: "get ruleId from /alert-rules endpoint. If alertsEnabled is set to 1 and alertRules is not included in a creation/update query, applicable defaults will be used.",
		Optional:    true,
		Required:    false,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"rule_id": {
					Type:        schema.TypeInt,
					Description: "If alertsEnabled is set to 1 and alertRules is not included in a creation/update query, applicable defaults will be used.",
					Optional:    true,
				},
			},
		},
	},
	"alerts_enabled": {
		Type:         schema.TypeInt,
		Description:  "choose 1 to enable alerts, or 0 to disable alerts. Defaults to 1",
		Optional:     true,
		Required:     false,
		Default:      1,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"auth_type": {
		Type:         schema.TypeString,
		Description:  "auth type",
		Optional:     true,
		Default:      "NONE",
		ValidateFunc: validation.StringInSlice([]string{"NONE", "BASIC", "NTLM", "KERBEROS"}, false),
	},
	"auth_user": {
		Type:        schema.TypeString,
		Description: "username for authentication with SIP server",
		Required:    true,
	},
	"bandwidth_measurements": {
		Type:        schema.TypeInt,
		Description: "set to 1 to measure bandwidth; defaults to 0. Only applies to Enterprise Agents assigned to the test, and requires that networkMeasurements is set.",
		Optional:    true,
		Required:    false,
		Default:     1,
	},
	"bgp_measurements": {
		Type:         schema.TypeInt,
		Description:  "choose 1 to enable bgp measurements, 0 to disable; defaults to 1",
		Optional:     true,
		Required:     false,
		Default:      1,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"bgp_monitors": {
		Type:        schema.TypeList,
		Optional:    true,
		Required:    false,
		Description: "array of BGP Monitor objects",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"monitor_id": {
					Type:        schema.TypeInt,
					Description: "monitor id",
					Optional:    true,
				},
			},
		},
	},
	"codec_id": {
		Type:         schema.TypeInt,
		Description:  "codec to use",
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 8),
	},
	"content_regex": {
		Type: schema.TypeString,
		Description: "regular Expressions	Verify content using a regular expression. This field does not require escaping",
		Optional: true,
		Default:  "NONE",
	},
	"description": {
		Type:        schema.TypeString,
		Required:    false,
		Optional:    true,
		Default:     "",
		Description: "defaults to empty string",
	},
	"desired_status_code": {
		Type: schema.TypeString,
		Description: "A valid HTTP response code	Set to the value you’re interested in retrieving",
		Optional: true,
	},
	"direction": {
		Type: schema.TypeString,
		Description: "[TO_TARGET, FROM_TARGET, BIDIRECTIONAL]	Direction of the test (affects how results are shown)",
		Optional:     false,
		Required:     true,
		ValidateFunc: validation.StringInSlice([]string{"TO_TARGET", "FROM_TARGET", "BIDIRECTIONAL"}, false),
	},
	"dns_servers": {
		Description: "array of DNS Server objects {“serverName”: “fqdn of server”}",
		Optional:    false,
		Required:    true,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"server_name": {
					Type:        schema.TypeString,
					Description: "DNS Server name",
					Optional:    true,
				},
			},
		},
	},
	"dns_transport_protocol": {
		Type: schema.TypeString,
		Description: "string	UDP or TCP	transport protocol used for DNS requests; defaults to UDP",
		Optional:     true,
		Required:     false,
		Default:      "UDP",
		ValidateFunc: validation.StringInSlice([]string{"UDP", "TCP"}, false),
	},
	"domain": {
		Type: schema.TypeString,
		Description: "see notes	target record for test, suffixed by record type (ie, www.thousandeyes.com CNAME). If no record type is specified, the test will default to an ANY record.",
		Optional: false,
		Required: true,
	},
	"dscp_id": {
		Type:         schema.TypeInt,
		Description:  "A Differentiated Services Code Point (DSCP) is a value found in an IP packet header which is used to request a level of priority for delivery (Defined in RFC 2474 https://www.ietf.org/rfc/rfc2474.txt). It is one of the Quality of Service management tools used in router configuration to protect real-time and high priority data applications.",
		Required:     false,
		Default:      0,
		Optional:     true,
		ValidateFunc: validation.IntInSlice([]int{0, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 44, 46, 48, 56}),
	},
	"duration": {
		Type:         schema.TypeInt,
		Description:  "duration of test, in seconds (5 to 30)",
		Optional:     true,
		ValidateFunc: validation.IntBetween(5, 30),
	},
	"enabled": {
		Type: schema.TypeInt,
		Description: "0 or 1	choose 1 to enable the test, 0 to disable the test",
		Optional:     true,
		Required:     false,
		Default:      1,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"http_interval": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "interval to run http server test on",
	},
	"http_target_time": {
		Type:         schema.TypeInt,
		Description:  "target time for HTTP server completion; specified in milliseconds",
		Optional:     true,
		Default:      1000,
		ValidateFunc: validation.IntBetween(100, 5000),
	},
	"http_time_limit": {
		Type:        schema.TypeInt,
		Description: "target time for HTTP server limits; specified in seconds",
		Default:     5,
		Optional:    true,
	},
	"http_version": {
		Type:         schema.TypeInt,
		Description:  "2 for default (prefer HTTP/2), 1 for HTTP/1.1 only",
		Default:      2,
		Optional:     true,
		ValidateFunc: validation.IntBetween(1, 2),
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "name of the test",
	},
	"include_covered_prefixes": {
		Type:         schema.TypeInt,
		Description:  "set to 1 to include queries for subprefixes detected under this prefix",
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"include_headers": {
		Type:         schema.TypeInt,
		Description:  "set to 1 to capture response headers for objects loaded by the test.Default is 1.",
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"interval": {
		Type:         schema.TypeInt,
		Required:     true,
		Description:  "interval to run test on, in seconds",
		ValidateFunc: validation.IntInSlice([]int{60, 120, 300, 600, 900, 1800, 3600}),
	},
	"jitter_buffer": {
		Type:         schema.TypeInt,
		Description:  "de-jitter buffer size, in seconds (0 to 150)",
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 150),
	},
	"mss": {
		Type: schema.TypeInt,
		Description: "(30..1400)	Maximum Segment Size, in bytes.",
		ValidateFunc: validation.IntBetween(30, 1400),
		Optional:     true,
		Required:     false,
	},
	"mtu_measurements": {
		Type:         schema.TypeInt,
		Description:  "set to 1 to measure MTU sizes on network from agents to the target.",
		Optional:     true,
		Required:     false,
		Default:      1,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"network_measurements": {
		Type: schema.TypeInt,
		Description: "integer	0 or 1	choose 1 to enable network measurements, 0 to disable; defaults to 1",
		Default:      1,
		Optional:     true,
		Required:     false,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"num_path_traces": {
		Type:         schema.TypeInt,
		Description:  "number of path traces. default 3.",
		Default:      3,
		Optional:     true,
		Required:     false,
		ValidateFunc: validation.IntBetween(3, 10),
	},
	"password": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "password to be used to authenticate with the destination server",
	},
	"path_trace_mode": {
		Type:         schema.TypeString,
		Description:  "choose inSession to perform the path trace within a TCP session; defaults to classic",
		Optional:     true,
		Required:     false,
		Default:      "classic",
		ValidateFunc: validation.StringInSlice([]string{"classic", "inSession"}, false),
	},
	"port": {
		Type:         schema.TypeInt,
		Description:  "target port for agent",
		Default:      80,
		ValidateFunc: validation.IntBetween(1, 65535),
		Optional:     true,
		Required:     false,
	},
	"prefix": {
		Type:        schema.TypeString,
		Description: "BGP network address prefix",
		Required:    true,
		// a.b.c.d is a network address, with the prefix length defined as e.
		// Prefixes can be any length from 8 to 24
		// Can only use private BGP monitors for a local prefix.
	},
	"probe_mode": {
		Type:         schema.TypeString,
		Description:  "probe mode used by End-to-end Network Test; only valid if protocol is set to TCP; defaults to AUTO",
		Optional:     true,
		Required:     false,
		Default:      "AUTO",
		ValidateFunc: validation.StringInSlice([]string{"AUTO", "SACK", "SYN"}, false),
	},
	"protocol": {
		Type:         schema.TypeString,
		Description:  "protocol used by dependent Network tests (End-to-end, Path Trace, PMTUD); defaults to TCP",
		Optional:     true,
		Required:     false,
		Default:      "TCP",
		ValidateFunc: validation.StringInSlice([]string{"TCP", "ICMP"}, false),
	},
	"protocol--sip": {
		Type:         schema.TypeString,
		Description:  "transport layer for SIP communication: TCP, TLS (TLS over TCP), or UDP. Defaults to TCP",
		Required:     true,
		ValidateFunc: validation.StringInSlice([]string{"TCP", "TLS", "UDP"}, false),
	},
	"recursive_queries": {
		Type:         schema.TypeInt,
		Default:      1,
		ValidateFunc: validation.IntBetween(0, 1),
		Description: "0 or 1	set to 1 to run query with RD (recursion desired) flag enabled",
		Optional: true,
		Required: false,
	},
	"request_type": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Set the type of activity for the test: Download, Upload, or List",
	},
	"server": {
		Type:        schema.TypeString,
		Description: "target host",
		Required:    true,
	},
	"ssl_version_id": {
		Type:         schema.TypeInt,
		Description:  "0 for auto, 3 for SSLv3, 4 for TLS v1.0, 5 for TLS v1.1, 6 for TLS v1.2",
		Optional:     true,
		Default:      0,
		ValidateFunc: validation.IntBetween(1, 2),
	},
	"sub_interval": {
		Type:         schema.TypeInt,
		Description:  "subinterval for round-robin testing (in seconds), must be less than or equal to interval",
		Optional:     true,
		ValidateFunc: validation.IntInSlice([]int{60, 120, 300, 600, 900, 1200, 1800, 3600}),
	},
	"target_agent_id": {
		Type:     schema.TypeInt,
		Optional: false,
		Required: true,
		Description: "pull from /agents endpoint	Both the 'agents': [] and the targetAgentID cannot be cloud agents. Can be Enterprise Agent -> Cloud, Cloud -> Enterprise Agent, or Enterprise Agent -> Enterprise Agent",
	},
	"target_time": {
		Type:         schema.TypeInt,
		Description:  "target time for completion, defaults to 50% of time limit; specified in seconds",
		Optional:     true,
		ValidateFunc: validation.IntBetween(1, 60),
	},
	"test_id": {
		Type:        schema.TypeInt,
		Description: "Unique ID of test",
		Computed:    true,
	},
	"test_name": {
		Type:        schema.TypeString,
		Description: "Name of Test",
		Required:    true,
	},
	"tests": {
		Type:        schema.TypeList,
		Description: "list of tests",
		Required:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"test_id": {
					Type:        schema.TypeInt,
					Description: "test id",
					Optional:    true,
				},
			},
		},
	},
	"time_limit": {
		Type:         schema.TypeInt,
		Description:  "time limit for transaction; defaults to 30s",
		Optional:     true,
		Default:      30,
		ValidateFunc: validation.IntBetween(1, 60),
	},
	"throughput_duration": {
		Type:         schema.TypeInt,
		Optional:     true,
		Required:     false,
		Default:      10000,
		Description:  "Defaults to 10000",
		ValidateFunc: validation.IntBetween(5000, 10000),
	},
	"throughput_measurements": {
		Type:         schema.TypeInt,
		ValidateFunc: validation.IntBetween(0, 1),
		Optional:     true,
		Required:     false,
		Default:      0,
		Description: "0 or 1	defaults to 0 (disabled), not allowed when source (or target) of the test is a cloud agent",
	},
	"throughput_rate": {
		Type:         schema.TypeInt,
		Description:  "for UDP only",
		Optional:     true,
		Required:     false,
		Default:      0,
		ValidateFunc: validation.IntBetween(0, 1000),
	},
	"transaction_script": {
		Type:        schema.TypeString,
		Description: "selenium transaction script",
		Required:    true,
	},
	"type": {
		Type: schema.TypeString,
		Description: "Type of test",
		Computed:  true,
	},
	"type--label": {
		Type:         schema.TypeString,
		Description:  "Type of label (tests, agents, or endpoint_agents",
		Required:     true,
		ValidateFunc: validation.StringInSlice([]string{"tests", "agents", "endpoint_agents"}, false),
	},
	"url": {
		Type:        schema.TypeString,
		Description: "target for the test",
		Required:    true,
	},
	"use_public_bgp": {
		Type:         schema.TypeInt,
		Description:  "set to 1 to automatically add all available Public BGP Monitors",
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 1),
	},
	"use_ntlm": {
		Type:        schema.TypeInt,
		Description: "choose 0 to use Basic Authentication, or omit field.Requires username/password to be set",
		Optional:    true,
	},
	"user": {
		Type:        schema.TypeString,
		Description: "username for SIP registration; should be unique within a ThousandEyes Account Group",
		Required:    true,
	},
	"user_agent": {
		Type:        schema.TypeString,
		Description: "user-agent string to be provided during the test",
		Optional:    true,
	},
	"username": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "username to be used to authenticate with the destination server",
	},
	"verify_certificate": {
		Type:        schema.TypeInt,
		Description: "set to 0 to ignore certificate errors (defaults to 1)",
		Optional:    true,
		Default:     1,
	},
}
