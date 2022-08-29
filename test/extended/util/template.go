package util

import (
	"fmt"
)

// ApplyClusterResourceFromTemplate apply the changes to the cluster resource.
// For ex: ApplyClusterResourceFromTemplate(oc, "--ignore-unknown-parameters=true", "-f", "TEMPLATE LOCATION")
func ApplyClusterResourceFromTemplate(oc *CLI, filepath string, variables ...string) {
	resourceFromTemplate(oc, false, "", filepath, variables...)
}

// ApplyNsResourceFromTemplate apply changes to the ns resource.
// No need to add a namespace parameter in the template file as it can be provided as a function argument.
// For ex: ApplyNsResourceFromTemplate(oc, "NAMESPACE", "--ignore-unknown-parameters=true", "-f", "TEMPLATE LOCATION")
func ApplyNsResourceFromTemplate(oc *CLI, namespace string, filepath string, variables ...string) {
	resourceFromTemplate(oc, false, namespace, filepath, variables...)
}

// CreateClusterResourceFromTemplate create resource from the template.
// For ex: CreateClusterResourceFromTemplate(oc, "--ignore-unknown-parameters=true", "-f", "TEMPLATE LOCATION")
func CreateClusterResourceFromTemplate(oc *CLI, filepath string, variables ...string) {
	resourceFromTemplate(oc, true, "", filepath, variables...)
}

// CreateNsResourceFromTemplate create ns resource from the template.
// No need to add a namespace parameter in the template file as it can be provided as a function argument.
// For ex: CreateNsResourceFromTemplate(oc, "NAMESPACE", "--ignore-unknown-parameters=true", "-f", "TEMPLATE LOCATION")
func CreateNsResourceFromTemplate(oc *CLI, namespace string, filepath string, variables ...string) {
	resourceFromTemplate(oc, true, namespace, filepath, variables...)
}

func resourceFromTemplate(oc *CLI, create bool, namespace string, filepath string, variables ...string) {
	parameters := []string{"-f", filepath}
	if namespace != "" {
		parameters = append(parameters, "-n", namespace)
	}

	oc.AddEnv(variables...)
	var resourceErr error
	if create {
		resourceErr = oc.AsAdmin().WithoutNamespace().Run("create").Args(parameters...).Execute()
	} else {
		resourceErr = oc.AsAdmin().WithoutNamespace().Run("apply").Args(parameters...).Execute()
	}
	oc.ClearEnv()
	AssertWaitPollNoErr(resourceErr, fmt.Sprintf("fail to create/apply resource %v", resourceErr))
}
