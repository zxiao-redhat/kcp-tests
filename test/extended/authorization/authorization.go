package authorization

import (
	"fmt"

	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"

	exutil "github.com/kcp-dev/kcp-tests/test/extended/util"
)

var _ = g.Describe("[sig-authorization]", func() {
	defer g.GinkgoRecover()

	var (
		k = exutil.NewCLIWithWorkSpace("kcp-workspace")
	)

	// author: zxiao@redhat.com
	g.It("Author:zxiao-Medium-[Smoke] Basic workspace content authorizer should work", func() {
		g.By("# Create a test workspace as parent workspace")
		k.SetupWorkSpace()
		parentWorkSpace := k.WorkSpace()

		g.By("# Create child workspace under parent workspace")
		k.SetupWorkSpaceWithSpecificPath(parentWorkSpace.ServerURL)
		childWorkSpace := k.WorkSpace()

		g.By("# Check if a service account with 'admin' access to parent workspace can access child workspace")
		g.By("## Create test service account")
		user := k.CreateUser("smoke")
		fmt.Println(user)
		sa1 := k.CreateServiceAccount()
		fmt.Println(sa1)

		g.By("## Create cluster role")
		clusterRoleName := fmt.Sprintf("cr-admin-%s", exutil.GetRandomString())
		clusterRoleTemplate := exutil.FixturePath("testdata", "authorization", "create-workspace-cluster-role.yaml")
		exutil.CreateClusterResourceFromTemplate(k, clusterRoleTemplate, "NAME="+clusterRoleName, "WORKSPACE="+parentWorkSpace.Name, "VERB=admin")

		g.By(fmt.Sprintf("## Give test service account %s 'admin' access to parent workspace", sa1))
		clusterRoleBindName := fmt.Sprintf("cr-bind-admin-%s", exutil.GetRandomString())
		clusterRoleBindTemplate := exutil.FixturePath("testdata", "authorization", "bind-workspace-cluster-role.yaml")
		exutil.ApplyClusterResourceFromTemplate(k, clusterRoleBindTemplate, "NAME="+clusterRoleBindName, "WORKSPACE="+parentWorkSpace.Name, "USERNAME="+sa1, "CLUSTER_ROLE="+clusterRoleName)

		g.By(fmt.Sprintf("## Check if service account %s can access parent workspace", sa1))
		err := k.ChangeUser(sa1).Run("kcp").Args("workspace", parentWorkSpace.Name).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		g.By(fmt.Sprintf("## Check if service account %s can access child workspace", sa1))
		err = k.ChangeUser(sa1).Run("kcp").Args("workspace", childWorkSpace.Name).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		g.By("# Check if a service account with 'admin' access to child workspace cannot access parent workspace")

	})
})
