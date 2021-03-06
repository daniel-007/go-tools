package pkg

type I interface {
	f1()
	f2()
}

func init() {
	var _ I
}

type t1 struct{}
type T2 struct{ t1 }

func (t1) f1() {}
func (T2) f2() {}

func Fn() {
	var v T2
	_ = v.t1
}

type I2 interface {
	f3()
	f4()
}

type t3 struct{}
type t4 struct {
	x int // MATCH /x is unused/
	y int // MATCH /y is unused/
	t3
}

func (*t3) f3() {}
func (*t4) f4() {}

func init() {
	var i I2 = &t4{}
	i.f3()
	i.f4()
}

type i3 interface {
	F()
}

type I4 interface {
	i3
}

type T5 struct {
	t6
}

type t6 struct {
	F int
}

type t7 struct{ X int }
type t8 struct{ t7 }
type t9 struct{ t8 }

var _ = t9{}
