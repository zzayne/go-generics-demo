package reflect

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetStructFields(t *testing.T) {
	Convey("TestGetStructFields", t, func() {
		p := Person{}
		out := GetStructFieldNames(p)
		So(len(out), ShouldEqual, 2)

		fmt.Println(out)
	})
}

func TestGetStructTypes(t *testing.T) {
	Convey("TestGetStructTypes", t, func() {
		p := Person{}
		out := GetStructFieldTypes(p)
		So(len(out), ShouldEqual, 2)

		fmt.Println(out)
	})
}

func TestSetFieldValue(t *testing.T) {
	Convey("TestSetFieldValue", t, func() {
		p := Person{}
		err := SetFieldValue(p, "Name", "zayne")
		So(err, ShouldEqual, ErrShouldBePrt)

		err = SetFieldValue(&p, "Name", 11111)
		So(err, ShouldEqual, ErrValueTypeNotMatch)

		err = SetFieldValue(&p, "Name", "zayne")
		So(err, ShouldBeNil)

		fmt.Println(p)
	})
}
