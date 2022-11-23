// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/runfinch/common-tests/command"
	"github.com/runfinch/common-tests/option"
)

var testVersion = func(o *option.Option) {
	ginkgo.Context("Version", func() {
		ginkgo.Specify("Test version", func() {
			//exec.Command("finch", "version").Start()
			cmd := exec.Command("finch", "version") //nolint:gosec // G204 is not an issue because cmdName is fully controlled by the user.
			session, err := gexec.Start(cmd, ginkgo.GinkgoWriter, ginkgo.GinkgoWriter)
			ginkgo.GinkgoWriter.Println("checkpoint 1")
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			ginkgo.GinkgoWriter.Println("checkpoint 2")
			session.Wait(10 * time.Second)
			ginkgo.GinkgoWriter.Println("checkpoint 3")
			gomega.Expect(session).Should(gexec.Exit(0))
			ginkgo.GinkgoWriter.Println("checkpoint 4")
			gomega.Expect(command.StdoutStr(o, "version")).To(gomega.Equal(fmt.Sprintf("Finch version: %s", "v0.1.0")))
		})
	})
}
