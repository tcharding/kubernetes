// Package resource provides information on valid resources
package resource

import "bytes"

const bulletPoint string = "* "

type Resource struct {
	name    string
	aliases []string
	info    string
}

var resources = []Resource{
	{name: "all"},
	{name: "certificatesigningrequests", aliases: []string{"csr"}},
	{name: "clusterrolebindings"},
	{name: "clusterroles"},
	{name: "clusters", info: "valid only for federation apiservers"},
	{name: "componentstatuses", aliases: []string{"cs"}},
	{name: "configmaps", aliases: []string{"cm"}},
	{name: "controllerrevisions"},
	{name: "cronjobs"},
	{name: "customresourcedefinition", aliases: []string{"crd"}},
	{name: "daemonsets", aliases: []string{"ds"}},
	{name: "deployments", aliases: []string{"deploy"}},
	{name: "endpoints", aliases: []string{"ep"}},
	{name: "events", aliases: []string{"ev"}},
	{name: "horizontalpodautoscalers", aliases: []string{"hpa"}},
	{name: "ingresses", aliases: []string{"ing"}},
	{name: "jobs"},
	{name: "limitranges", aliases: []string{"limits"}},
	{name: "namespaces", aliases: []string{"ns"}},
	{name: "networkpolicies", aliases: []string{"netpol"}},
	{name: "nodes", aliases: []string{"no"}},
	{name: "persistentvolumeclaims", aliases: []string{"pvc"}},
	{name: "persistentvolumes", aliases: []string{"pv"}},
	{name: "poddisruptionbudgets", aliases: []string{"pdb"}},
	{name: "podpreset"},
	{name: "pods", aliases: []string{"po"}},
	{name: "podsecuritypolicies", aliases: []string{"psp"}},
	{name: "podtemplates"},
	{name: "replicasets", aliases: []string{"rs"}},
	{name: "replicationcontrollers", aliases: []string{"rc"}},
	{name: "resourcequotas", aliases: []string{"quota"}},
	{name: "rolebindings"},
	{name: "roles"},
	{name: "secrets"},
	{name: "serviceaccounts", aliases: []string{"sa"}},
	{name: "services", aliases: []string{"svc"}},
	{name: "statefulsets"},
	{name: "storageclasses"},
}

// ValidResources returns a list of valid resources, formatted for help output.
func ValidResources() string {
	var bb bytes.Buffer

	_, err := bb.WriteString("Valid resource types include:\n\n")
	if err != nil {
		panic(err)
	}
	for _, resource := range resources {
		bb.WriteString(bulletPoint)
		bb.WriteString(resource.name)
		if len(resource.aliases) > 0 {
			bb.WriteString(" (aka '")
			for i, alias := range resource.aliases {
				bb.WriteString(alias)
				if i != len(resource.aliases)-1 {
					bb.WriteString("', '")
				} else {
					bb.WriteString("')")
				}
			}
		}
		if resource.info != "" {
			s := " (" + resource.info + ")"
			bb.WriteString(s)
		}
		bb.WriteString("\n")
	}

	bb.WriteString("\n")

	return bb.String()
}

// ValidResourcesForCommand returns command specific valid resources.
func ValidResourcesForCommand(cmd string) string {
	switch cmd {
	case "explain":
		return validResourcesForExplain()
	default:
		return ValidResources()
	}
}

func validResourcesForExplain() string {

}

// AvailableResources returns list of available resources for use with kubectl get.
func AvailableResources() string {

}
