package ginkgo_test

import (
	"github.com/rrd1986/ginkgo_gomega_example/service"
	set "github.com/rrd1986/ginkgo_gomega_example/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var s *service.Set

var _ = Describe("Set", func() {

	BeforeEach(func() {
		s = set.NewSet()
	})

	Describe("Emptyness", func() {
		Context("When the set doe not contain items", func() {
			It("Should be empty", func() {
				s := set.NewSet()
				Expect(s.IsEmpty()).To(BeTrue())
			})
		})

		Context("When the set contain items", func() {
			It("Should not be empty", func() {
				s := set.NewSet()
				s.Add("item1")
				Expect(s.IsEmpty()).To(BeFalse())
			})
		})

	})

	Describe("Size", func() {
		Context("As items are added", func() {
			It("Should return an increasing size", func() {
				By("Empty set size being 0", func() {
					Expect(s.Size()).To(BeZero())
				})

				By("Adding a first item", func() {
					s.Add("red")

					Expect(s.Size()).To(Equal(1))
				})

				By("Adding a second item", func() {
					s.Add("blue")

					Expect(s.Size()).To(Equal(2))
				})
			})
		})
	})

	Describe("Contains", func() {
		Context("When red has not been added", func() {
			It("Should not contain red", func() {
				Expect(s.Contains("red")).To(BeFalse())
			})
		})

		Context("When red has been adding", func() {
			It("Should contain red", func() {
				s.Add("red")

				Expect(s.Contains("red")).To(BeTrue())
			})
		})
	})

})
