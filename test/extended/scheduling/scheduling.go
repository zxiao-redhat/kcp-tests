package scheduling

import (
    g "github.com/onsi/ginkgo"
    o "github.com/onsi/gomega"

    exutil "github.com/kcp-dev/kcp-tests/test/extended/util"
)

var _ = g.Describe("[area-apiexports]", func() {
    defer g.GinkgoRecover()

    var (
        k = exutil.NewCLIWithWorkSpace("kcp-scheduling")
    )

    g.It("Author:zxiao-Critical-[Smoke] Verify syncTarget can be added, used, and deleted", func() {
        // Shared compute could be only accessed from dev-provided test environments
        // Skip for non-supported test environments
        exutil.PreCheckEnvSupport(k, "kcp-stable.apps.kcp-internal", "kcp-unstable.apps.kcp-internal")
        myWs := k.WorkSpace()

    })
})
