package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSlicesEqual(t *testing.T) {
	Convey("测试切片相等性函数", t, func() {
		a := []int{1, 2, 3, 4}
		b := []int{1, 2, 3, 4}
		So(SlicesEqual(a, b), ShouldBeTrue) // a和b相等，这个判定应该为true，如果确实相等，则单测绘PASS，否侧不通过
	})
}
