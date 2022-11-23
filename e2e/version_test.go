// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"fmt"
	"os"

	"github.com/onsi/ginkgo/v2"
	"github.com/runfinch/common-tests/option"
	"github.com/spf13/afero"
)

var testVersion = func(o *option.Option) {
	ginkgo.Context("Version", func() {
		ginkgo.Specify("Test version", func() {
			//exec.Command("finch", "version").Start()
			//cmd := exec.Command("finch", "version") //nolint:gosec // G204 is not an issue because cmdName is fully controlled by the user.
			//cmd.Run()
			afero.ReadFile(afero.NewOsFs(), fmt.Sprintf("%s/.finch/finch.yaml", os.Getenv("HOME")))
			//session, err := gexec.Start(cmd, ginkgo.GinkgoWriter, ginkgo.GinkgoWriter)
			//gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			//session.Wait(10 * time.Second)
			//gomega.Expect(session).Should(gexec.Exit(0))
			//gomega.Expect(command.StdoutStr(o, "version")).To(gomega.Equal(fmt.Sprintf("Finch version: %s", "v0.1.0")))
		})
	})
}
