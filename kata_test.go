package kata_test

import (
	. "codewarrior/kata"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

var _ = Describe("Test Example", func() {
	It("should work for sample test cases", func() {
		Expect(ToJadenCase("most trees are blue")).To(Equal("Most Trees Are Blue"))
		Expect(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own.")).To(Equal("All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."))
		Expect(ToJadenCase("When I die. then you will realize")).To(Equal("When I Die. Then You Will Realize"))
		Expect(ToJadenCase("Jonah Hill is a genius")).To(Equal("Jonah Hill Is A Genius"))
		Expect(ToJadenCase("Dying is mainstream")).To(Equal("Dying Is Mainstream"))
	})
})

// var _ = Describe("Example Test", func() {
// 	It("should test that the solution returns the correct value", func() {
// 		Expect(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")).To(Equal("42 -9"))
// 	})
// })

// func dotest(s string, exp string) {
// 	var ans = Accum(s)
// 	Expect(ans).To(Equal(exp))
// }

// var _ = Describe("Test Example", func() {

// 	It("should handle basic cases", func() {
// 		dotest("ZpglnRxqenU", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
// 		dotest("NyffsGeyylB", "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb")
// 		dotest("MjtkuBovqrU", "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu")
// 		dotest("EvidjUnokmM", "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm")
// 		dotest("HbideVbxncC", "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc")
// 	})
// })
