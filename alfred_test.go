package Alfred

import (
	"os"
	"path"
	"testing"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInitAlfred(t *testing.T) {
	Convey("Test Init Alfred", t, func() {
		id := "test init alfred"
		al, err := NewAlfred(id)

		Convey("Need No Err", func() {
			So(err, ShouldBeNil)
		})

		Convey("Check ID", func() {
			So(al.GetID(), ShouldEqual, id)
			t.Log("\tID: " + al.GetID())
		})

		Convey("Check Home Dir", func() {
			So(al.GetHomeDir(), ShouldEqual, os.Getenv("HOME"))
			t.Log("\tHome Dir: " + al.GetHomeDir())
		})

		Convey("Check Bundle ID", func() {
			So(al.GetBundleID(), ShouldEqual, "com.harry.alfred.workflow.devtoolkit")
			t.Log("\tBundle ID: " + al.GetBundleID())
		})

		Convey("Check Bundle Dir", func() {
			pwd, err := os.Getwd()
			if err != nil {
				t.Fatal("get pwd fail", err)
				t.FailNow()
			}
			So(al.GetBundleDir(), ShouldEqual, pwd)
			t.Log("\tBundle Dir: " + al.GetBundleDir())
		})

		Convey("Check Cache Dir", func() {
			So(al.GetCacheDir(), ShouldEqual, path.Join(al.GetHomeDir(), defaultCacheDir, al.GetBundleID()))
			t.Log("\tCache Dir: " + al.GetCacheDir())
		})

		Convey("Check Data Dir", func() {
			So(al.GetDataDir(), ShouldEqual, path.Join(al.GetHomeDir(), defaultDataDir, al.GetBundleID()))
			t.Log("\tData Dir: " + al.GetDataDir())
		})
	})
}

func TestAlfredResult(t *testing.T) {
	Convey("Test Alfred Result", t, func() {
		al, err := NewAlfred("test alfred result")

		Convey("Need No Err", func() {
			So(err, ShouldBeNil)
		})

		Convey("Need No Json Err", func() {
			al.ResultAppend(NewNoResultItem())

			err = errors.New("error1")
			err2 := errors.Wrap(err, "wrap error")
			al.ResultAppend(NewErrorItem(err))
			al.ResultAppend(NewErrorItem(err2))

			json, err := al.ResultToIndentJson()
			So(err, ShouldBeNil)
			t.Log("\nResult Data:\n" + string(json))
		})
	})
}
