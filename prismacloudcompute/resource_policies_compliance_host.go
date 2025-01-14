package prismacloudcompute

import (
	"log"
	"time"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyComplianceHost"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesComplianceHost() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyComplianceHost,
		Read:   readPolicyComplianceHost,
		Update: updatePolicyComplianceHost,
		Delete: deletePolicyComplianceHost,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the policy set.",
			},
			"policytype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of policy. For example: 'docker', 'containerVulnerability', 'containerCompliance', etc.",
				Default:     true,
			},
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of policy rules.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Action to take.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"alertthreshold": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The compliance container policy alert threshold. Threshold values typically vary between 0 and 10 (non-inclusive).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables alerts.",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', suppresses alerts for all compliance containers.",
									},
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Minimum severity to trigger alerts. Supported values range from 0 to 9, where 0=off, 1=low, and 9=critical.",
									},
								},
							},
						},
						"allcompliance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', reports the results of all (passed and failed) compliance checks.",
						},
						"auditallowed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', Prisma Cloud audits successful transactions.",
						},
						"blockmsg": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Represents the block message in a policy.",
						},
						"blockthreshold": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The compliance container policy block threshold. Threshold values typically vary between 0 and 10 (non-inclusive).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables blocking.",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', suppresses blocking for all compliance containers.",
									},
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Minimum severity to trigger blocking. Supported values range from 0 to 9, where 0=off, 1=low, and 9=critical.",
									},
								},
							},
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of collections. Used to scope the rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Output.
									"accountids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of account IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"appids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of application IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of Kubernetes cluster names.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"coderepos": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of code repositories.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"color": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A hex color code for a collection.",
									},
									"containers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of containers.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A free-form text description of the collection.",
									},
									"functions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of functions.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hosts": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of hosts.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"images": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of images.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of labels.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"modified": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Date/time when the collection was last modified.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique collection name.",
									},
									"namespaces": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of Kubernetes namespaces.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "User who created or last modified the collection.",
									},
									"prisma": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', this collection originates from Prisma Cloud.",
									},
									"system": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', this collection was created by the system (i.e., a non-user). Otherwise (false) a real user.",
									},
								},
							},
						},
						"condition": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Rule conditions. Conditions only apply for their respective policy type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Allowed volume host device (wildcard). If a 'container create' command specifies a non-matching host device, the action is blocked. Only applies to rules in certain policy types.",
									},
									"readonly": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', the condition applies only to read-only commands. For example: HTTP GET requests.",
									},
									"vulnerabilities": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Block and scan severity-based vulnerabilities conditions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"block": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the effect is blocked.",
												},
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Compliance container ID.",
												},
											},
										},
									},
								},
							},
						},
						"cverules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of Common Vulnerability and Exposure (CVE) IDs classified for special handling/exceptions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Free-form text for documenting the exception.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the relevant action for a compliance container. Can be set to 'ignore', 'alert', or 'block'.",
									},
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CVE ID",
									},
									"expiration": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The compliance container expiration date.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"date": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The date the compliance container expires.",
												},
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the grace period is enabled.",
												},
											},
										},
									},
								},
							},
						},
						"disabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', the rule is currently disabled.",
						},
						"effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The effect of evaluating the given policy. Can be set to 'allow', 'deny', 'block', or 'alert'.",
						},
						"gracedays": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of days to suppress the rule's block effect. Measured from date the compliance container was fixed. If there's no fix, measured from the date the compliance container was published.",
						},
						"group": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Applicable groups.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"license": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The configuration of the compliance policy license.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alertThreshold": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The license severity threshold to indicate whether to perform an alert action. Threshold values typically vary between 0 and 10 (non-inclusive).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the alert action is enabled.",
												},
												"disabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', suppresses alerts for all compliance containers.",
												},
												"value": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The minimum severity score for which the alert action is enabled.",
												},
											},
										},
									},
									"blockThreshold": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The license severity threshold to indicate whether to perform a block action. Threshold values typically vary between 0 and 10 (non-inclusive).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the block action is enabled.",
												},
												"disabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', suppresses blocking for all compliance containers.",
												},
												"value": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The minimum severity score for which the block action is enabled.",
												},
											},
										},
									},
									"critical": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with critical severity.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"high": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with high severity.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"low": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with low severity.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"medium": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with medium severity.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"modified": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date/time when the rule was last modified.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the rule.",
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text notes.",
						},
						"onlyfixed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', applies rule only when vendor fixes are available.",
						},
						"owner": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "User who created or last modified the rule.",
						},
						"previousname": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Previous name of the rule. Required for rule renaming.",
						},
						"principal": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Applicable users.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of tags classified for special handling/exceptions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Free-form text for documenting the exception.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the relevant action for a compliance container. Can be set to 'ignore', 'alert', or 'block'.",
									},
									"expiration": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The compliance container expiration date.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"date": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Date of the compliance container expiration.",
												},
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the grace period is enabled.",
												},
											},
										},
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Tag name.",
									},
								},
							},
						},
						"verbose": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', displays a detailed message when an operation is blocked.",
						},
					},
				},
			},
		},
	}
}

func parsePolicyComplianceHost(d *schema.ResourceData, id string) policyComplianceHost.Policy {
	ans := policyComplianceHost.Policy{
		PolicyId: id,
	}
	if d.Get("policytype") != nil {
		ans.PolicyType = d.Get("policytype").(string)
	}

	rules := d.Get("rules").([]interface{})
	ans.Rules = parseRules(rules)

	return ans
}

func savePolicyComplianceHost(d *schema.ResourceData, obj policyComplianceHost.Policy) {
	d.Set("_id", obj.PolicyId)
	d.Set("policytype", obj.PolicyType)
	d.Set("rules", obj.Rules)

	// Rule.
	if len(obj.Rules) > 0 {
		rv := map[string]interface{}{
			"advancedprotection":       obj.Rules[0].AdvancedProtection,
			"cloudmetadataenforcement": obj.Rules[0].CloudMetadataEnforcement,
			"collections":              obj.Rules[0].Collections,
			"customrules":              obj.Rules[0].CustomRules,
			"disabled":                 obj.Rules[0].Disabled,
			"dns":                      obj.Rules[0].Dns,
			"filesystem":               obj.Rules[0].Filesystem,
			"kubernetesenforcement":    obj.Rules[0].KubernetesEnforcement,
			"modified":                 obj.Rules[0].Modified,
			"name":                     obj.Rules[0].Name,
			"network":                  obj.Rules[0].Network,
			"notes":                    obj.Rules[0].Notes,
			"owner":                    obj.Rules[0].Owner,
			"previousname":             obj.Rules[0].PreviousName,
			"processes":                obj.Rules[0].Processes,
			"wildfireanalysis":         obj.Rules[0].WildFireAnalysis,
		}

		if err := d.Set("rules", []interface{}{rv}); err != nil {
			log.Printf("[WARN] Error setting 'rules' for %q: %s", d.Id(), err)
		}
	}

}

func createPolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parsePolicyComplianceHost(d, "")

	if err := policyComplianceHost.Create(client, obj); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := policyComplianceHost.Get(client)
		return err
	})

	pol, err := policyComplianceHost.Get(client)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicyComplianceHost(d, meta)
}

func readPolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	obj, err := policyComplianceHost.Get(client)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	savePolicyComplianceHost(d, obj)

	return nil
}

func updatePolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parsePolicyComplianceHost(d, id)

	if err := policyComplianceHost.Update(client, obj); err != nil {
		return err
	}

	return readPolicyComplianceHost(d, meta)
}

func deletePolicyComplianceHost(d *schema.ResourceData, meta interface{}) error {
	/*	client := meta.(*pc.Client)
		id := d.Id()

		err := policy.Delete(client, id)
		if err != nil {
			if err != pc.ObjectNotFoundError {
				return err
			}
		}*/

	d.SetId("")
	return nil
}
