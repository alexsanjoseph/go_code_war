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

func dotest(k, start, nd int, exp []int) {
	var ans = Step(k, start, nd)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Check for Primes", func() {

	It("should handle basic cases", func() {
		Expect(true).To(Equal(IsPrime(7)))
		Expect(true).To(Equal(IsPrime(2)))
		Expect(true).To(Equal(IsPrime(17)))
		Expect(true).To(Equal(IsPrime(101)))
		Expect(false).To(Equal(IsPrime(1)))
		Expect(false).To(Equal(IsPrime(1001)))
		Expect(false).To(Equal(IsPrime(50)))
		Expect(false).To(Equal(IsPrime(25)))
		Expect(false).To(Equal(IsPrime(30148)))
	})
})

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(2, 100, 110, []int{101, 103})
		dotest(4, 100, 110, []int{103, 107})
		dotest(6, 100, 110, []int{101, 107})
		dotest(8, 300, 400, []int{359, 367})
		dotest(10, 300, 400, []int{307, 317})
		dotest(11, 30000, 100000, nil)
	})
})

// func dotest(f FParam, arr []int, init int, exp []int) {
// 	var ans = OperArray(f, arr, init)
// 	Expect(ans).To(Equal(exp))
// }

// var _ = Describe("Test Example", func() {

// 	It("should handle basic cases gcdi", func() {
// 		var dta = []int{18, 69, -90, -78, 65, 40}
// 		var sol = []int{18, 3, 3, 3, 1, 1}
// 		dotest(Gcdi, dta, dta[0], sol)
// 	})

// 	It("should handle basic cases som", func() {
// 		var dta = []int{357, 112, 28, -52, 644, 119}
// 		var sol = []int{357, 469, 497, 445, 1089, 1208}
// 		dotest(Som, dta, 0, sol)
// 	})
// 	It("should handle basic cases maxi", func() {
// 		var dta = []int{10, -32, 190, 300, -42, -38, 50, 405, -46, 225, -31}
// 		var sol = []int{10, 10, 190, 300, 300, 300, 300, 405, 405, 405, 405}
// 		dotest(Maxi, dta, dta[0], sol)
// 	})
// 	It("should handle basic cases lcmu", func() {
// 		var dta = []int{6, -72, -62, -22, -23, 80}
// 		var sol = []int{6, 72, 2232, 24552, 564696, 5646960}
// 		dotest(Lcmu, dta, dta[0], sol)
// 	})
// 	It("should handle basic cases mini", func() {
// 		var dta = []int{64, -67, -43, 12, -15, 108, 12, 104, -36}
// 		var sol = []int{64, -67, -67, -67, -67, -67, -67, -67, -67}
// 		dotest(Mini, dta, dta[0], sol)
// 	})
// })

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
