package hit

import (
	"errors"
	"image"
	"log"
	"strings"
	"testing"
)

func TestIf0(t *testing.T) {
	a := If(500)
	if a != 500 {
		t.Errorf(`a := If(500)`)
	}
}
func TestIf0a(t *testing.T) {
	var b int
	a := If(func() { b = 10 })
	if a != nil && b != 10 {
		t.Errorf(`a := If(func() { b = 10 })`)
	}
}

func TestIf0b(t *testing.T) {
	var b int
	a := If(func() interface{} { b = 10; return 30 })
	if a != 30 && b != 10 {
		t.Errorf(`a := If(func() { b = 10 })`)
	}
}

func TestIf1(t *testing.T) {
	expect := 5
	a := If(true, 5, 50)
	if a != expect {
		t.Errorf(`a := If(true, 5, 50)`)
	}
}

func TestIf1b(t *testing.T) {
	var b = []string{"aa", "bb"}
	a := If(len(b) > 10, b[1], "cancel")
	if a != "cancel" {
		t.Errorf(`a := If(len(b) > 1, b[1], "cancel")`)
	}
}

func TestIf2(t *testing.T) {
	expect := 50
	a := If(false, 5, 50)
	if a != expect {
		t.Errorf(`a := If(false, 5, 50)`)
	}
}

func TestIf3(t *testing.T) {
	expect := 50
	a := If(nil, 5, 50)
	if a != expect {
		t.Errorf(`a := If(nil, 5, 50)`)
	}
}

func TestIf4(t *testing.T) {
	expect := 50
	a := If("", 5, 50)
	if a != expect {
		t.Errorf(`a := If("", 5, 50)`)
	}
}

func TestIf5(t *testing.T) {
	expect := 50
	a := If("0", 5, 50)
	if a != expect {
		t.Errorf(`a := If("0", 5, 50)`)
	}
}

func TestIf6(t *testing.T) {
	expect := 50
	a := If("false", 5, 50)
	if a != expect {
		t.Errorf(`a := If("false", 5, 50)`)
	}
}

func TestIf7(t *testing.T) {
	expect := 50
	a := If(0, 5, 50)
	if a != expect {
		t.Errorf(`a := If(0, 5, 50)`)
	}
}

func TestIf7a(t *testing.T) {
	expect := 50
	a := If(0.0, 5, 50)
	if a != expect {
		t.Errorf(`a := If(0.0, 5, 50)`)
	}
}

func TestIf7b(t *testing.T) {
	expect := 50
	a := If(uint(0), 5, 50)
	if a != expect {
		t.Errorf(`a := If(uint(0), 5, 50)`)
	}
}

func TestIf8(t *testing.T) {
	expect := 5
	a := If(299, 5, 50)
	if a != expect {
		t.Errorf(`a := If(299, 5, 50)`)
	}
}

func TestIf9(t *testing.T) {
	expect := 5
	a := If("test", 5, 50)
	if a != expect {
		t.Errorf(`a := If("test", 5, 50)`)
	}
}

func TestIf10(t *testing.T) {
	expect := 5
	a := If(log.Printf, 5, 50)
	if a != expect {
		t.Errorf(`a := If(log.Printf, 5, 50)`)
	}
}

func TestIf11(t *testing.T) {
	expect := 5
	a := If(func() bool { return true }(), 5, 50)
	if a != expect {
		t.Errorf(`a := If(func() bool { return true }(), 5, 50)`)
	}
}

func TestIf12(t *testing.T) {
	a := If(true, func() interface{} { return "run-func" }, 50)
	log.Println(a)
	if a != "run-func" {
		t.Errorf(`a := If(true, func() interface{} { return "run-func" }, 50)`)
	}
}

func TestIf13(t *testing.T) {
	a := If(func() interface{} { return "false" }, 5, 50)
	log.Println(a)
	if a != 50 {
		t.Errorf(`a := If(func() interface{} { return "false" }, 5, 50)`)
	}
}

func TestIf14(t *testing.T) {
	a := If(func() interface{} { return 123 }, 5, 50)
	log.Println(a)
	if a != 5 {
		t.Errorf(`a := If(func() interface{} { return 123 }, 5, 50)`)
	}
}

func TestIf15(t *testing.T) {
	a := If(20 > 5, func() interface{} { return "ok" }, func() interface{} { return "cancel" })
	if a != "ok" {
		t.Errorf(`a := If(20 > 5, func() interface{} { return "ok" }, func() interface{} { return "cancel" })`)
	}
}

func TestAnd1(t *testing.T) {
	expect := 5
	a := If(true, 5)
	if a != expect {
		t.Errorf(`a := If(true, 5)`)
	}
}

func TestAnd2(t *testing.T) {
	a := If(false, 5)
	if a != nil {
		t.Errorf(`a := If(false, 5)`)
	}
}

func TestAnd3(t *testing.T) {
	a := If(200, 5)
	if a != 5 {
		t.Errorf(`a := If(200, 5)`)
	}
}

func TestAnd4(t *testing.T) {
	a := If("false", 5)
	if a != nil {
		t.Errorf(`a := If("false", 5)`)
	}
}

func TestAnd5(t *testing.T) {
	var b string
	a := If(func() { b = "ok" }, 5)
	if a != nil {
		t.Errorf(`a := If("false", 5)`)
	}
	if b != "ok" {
		t.Errorf(`a := If("false", 5)`)
	}
}

func TestAnd6(t *testing.T) {
	var b string
	a := If(1, func() interface{} { b = "ok"; return 10 })
	if a != 10 {
		t.Errorf(`a := If("false", 5)`)
	}
	if b != "ok" {
		t.Errorf(`a := If("false", 5)`)
	}
}

func TestAnd7(t *testing.T) {
	a := If(0, 5)
	if a != nil {
		t.Errorf(`a := If(0, 5)`)
	}
}

func TestAnd7a(t *testing.T) {
	a := If(0.0, 5)
	if a != nil {
		t.Errorf(`a := If(0.0, 5)`)
	}
}

func TestAnd7b(t *testing.T) {
	a := If(uint(0), 5)
	if a != nil {
		t.Errorf(`a := If(uint(0), 5)`)
	}
}

func TestOr0(t *testing.T) {
	a := Or(500)
	if a != 500 {
		t.Errorf(`a := If(500)`)
	}
}
func TestOr0a(t *testing.T) {
	var b int
	a := Or(func() { b = 10 })
	if a != nil && b != 10 {
		t.Errorf(`a := If(func() { b = 10 })`)
	}
}

func TestOr0b(t *testing.T) {
	var b int
	a := Or(func() interface{} { b = 10; return 30 })
	if a != 30 && b != 10 {
		t.Errorf(`a := If(func() { b = 10 })`)
	}
}

func TestOr1(t *testing.T) {
	a := Or(true, 5)
	if a != true {
		t.Errorf(`a := Or(true, 5)`)
	}
}

func TestOr2(t *testing.T) {
	a := Or(400, 5)
	if a != 400 {
		t.Errorf(`a := Or(400, 5)`)
	}
}

func TestOr3(t *testing.T) {
	a := Or(0, 5)
	if a != 5 {
		t.Errorf(`a := Or(0, 5)`)
	}
}

func TestOr3a(t *testing.T) {
	a := Or(0.0, 5)
	if a != 5 {
		t.Errorf(`a := Or(0.0, 5)`)
	}
}

func TestOr3b(t *testing.T) {
	a := Or(uint(0), 5)
	if a != 5 {
		t.Errorf(`a := Or(uint(0), 5)`)
	}
}

func TestOr4(t *testing.T) {
	a := Or("0", 5)
	if a != 5 {
		t.Errorf(`a := Or("0", 5)`)
	}
}

func TestOr5(t *testing.T) {
	a := Or("false", 5)
	if a != 5 {
		t.Errorf(`a := Or("false", 5)`)
	}
}

func TestOr6(t *testing.T) {
	a := Or("test", 5)
	if a != "test" {
		t.Errorf(`a := Or("test", 5)`)
	}
}

func TestOr7(t *testing.T) {
	a := Or(nil, 5)
	if a != 5 {
		t.Errorf(`a := Or(nil, 5)`)
	}
}

func TestOr8(t *testing.T) {
	var b int
	a := Or(func() { b = 100 }, 5)
	if a != 5 && b != 100 {
		t.Errorf(`a := Or(func() { b = 100 }, 5)`)
	}
}

func TestOr9(t *testing.T) {
	var b int
	a := Or(func() interface{} { b = 100; return 500 }, 5)
	if a != 500 && b != 100 {
		t.Errorf(`a := Or(func() { b = 100 }, 5)`)
	}
}

func TestOr10(t *testing.T) {
	var b int
	a := Or("", func() interface{} { b = 100; return 500 })
	if a != 500 && b != 100 {
		t.Errorf(`a := Or("", func() interface{} { b = 100; return 500 })`)
	}
}

func TestOr11(t *testing.T) {
	var b int
	a := Or("", func() interface{} { b = 100; return 500 })
	if a != 500 && b != 100 {
		t.Errorf(`a := Or("", func() interface{} { b = 100; return 500 })`)
	}
}

func TestOr12(t *testing.T) {
	var b int
	a := Or(false, func() interface{} { b = 100; return 500 })
	if a != 500 && b != 100 {
		t.Errorf(`a := Or("", func() interface{} { b = 100; return 500 })`)
	}
}

func TestIfError1(t *testing.T) {
	a := If(errors.New("test-error"))
	if _, ok := a.(error); ok == false {
		t.Errorf(`a := If(errors.New("test-error"))`)
	}
}

func TestIfError2(t *testing.T) {
	var b int
	a := If(errors.New("test-error"), func() { b = 10 })
	if _, ok := a.(error); ok == false && b == 10 {
		t.Errorf(`a := If(errors.New("test-error"), func() { b = 10 })`)
	}
}

func TestFnTime1(t *testing.T) {
	a := TestFnTime(func() {})
	if strings.Contains(a, "TestFnTime1") == false {
		t.Errorf(`a := TestFnTime(func() {})`)
	}
}

func TestCallFn1(t *testing.T) {
	var b int
	a := callFn(func() { b = 100 })
	if a != nil && b != 100 {
		t.Errorf(`a := callFn(func() { b = 100 })`)
	}
}

func TestCallFn2(t *testing.T) {
	var b int
	a := callFn(func() interface{} { b = 100; return 10 })
	if a != 10 && b != 100 {
		t.Errorf(`a := callFn(func() interface{} { b = 100; return 10 })`)
	}
}

func TestCallFn3(t *testing.T) {
	var b int
	a := callFn(func() int { b = 100; return 10 })
	if a != 10 && b != 100 {
		t.Errorf(`a := callFn(func() int { b = 100; return 10 })`)
	}
}

func TestCallFn4(t *testing.T) {
	var b int
	a := callFn(func() image.Point { b = 100; return image.Point{0, 0} })
	if v, ok := a.(image.Point); !ok || (!v.Eq(image.Point{0,0}) && b != 100) {
		t.Errorf(`a := callFn(func() image.Point { b = 100; return image.Point{0, 0} })`)
	}
}

func TestOrError1(t *testing.T) {
	a := Or(errors.New("test-error"))
	if _, ok := a.(error); ok == false {
		t.Errorf(`a := If(errors.New("test-error"))`)
	}
}

func TestOrError2(t *testing.T) {
	var b int
	a := Or(errors.New("test-error"), func() { b = 10 })
	if _, ok := a.(error); ok == false && b == 10 {
		t.Errorf(`a := If(errors.New("test-error"), func() { b = 10 })`)
	}
}
