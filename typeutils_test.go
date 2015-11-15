package typeutils

import (
	"testing"
	"github.com/wolfgarnet/logging"
)

type typeA struct  {
	Field string
	Name string
	Age int
}

type typeB struct  {
	typeA
	Title string
}

func TestFindField(t *testing.T) {
	instance := typeB{typeA{"fff", "snade", 10}, "the title"}
	result := FindField(instance, "Title", -1)

	if result != "the title" {
		t.Errorf("The result, %v, was not \"the title\"", result)
	}

	result = FindField(instance, "Age", -1)

	if result != 10 {
		t.Errorf("The result, %v, was not \"10\"", result)
	}

	result = FindField(instance, "Snade", -1)

	if result != nil {
		t.Errorf("The result, %v, was not nil", result)
	}
}


type testStruct struct {
	Snade int
}

type testStruct2 struct {
	testStruct
}

func (t testStruct2) Super() interface{} {
	return &t.testStruct
}

func (t testStruct) test() {
	println("CALLED TEST")
}

func (t testStruct) String() string {
	return "Hello"
}

func TestTypeImplements(t *testing.T) {
	test := testStruct2{}

	r := TypeImplements(test, "test")

	test.test()
	println(test.Snade)

	println(r.Name())
}


func init() {
	logging.SetLevel(logging.LevelTrace)
}