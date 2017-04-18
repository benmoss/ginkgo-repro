package ginkgo_repro_test

import (
	. "ginkgo-repro"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Foo", func() {
	It("works", func() {
		Expect(One).To(Equal("one"))
	})
})
