package adapter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/vast/adapter/base"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetTrackInfo(t *testing.T) {
	Range(func(content []byte) {

		vi, err := GetTrackInfo(string(content), base.VAST20)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(vi)

	}, t)

	Range(func(content []byte) {
		vi := &base.VASTInfo{
			ImpTrack:     []string{"http://imptrack.com?a=1"},
			ClickTrack:   []string{"http://clicktrack.com?a=2"},
			ClickThrough: []string{"http://clickthrough.com?a=3"},
		}

		adm, err := SetTrackInfo(string(content), vi, base.VAST20)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(adm)

	}, t)

}

func TestSetTrackInfo(t *testing.T) {
	Range(func(content []byte) {
		vi := &base.VASTInfo{
			ImpTrack:     []string{"http://imptrack.com?a=1"},
			ClickTrack:   []string{"http://clicktrack.com?a=2"},
			ClickThrough: []string{"http://clickthrough.com?a=3"},
		}

		adm, err := SetTrackInfo(string(content), vi, base.VAST20)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(adm)

	}, t)
}

func Range(tester func(content []byte), t *testing.T) {
	_, curFilename, _, _ := runtime.Caller(0)

	dir := filepath.Join(path.Dir(curFilename), "testdata")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Errorf("读取目录失败：%v", err)
	}
	for _, file := range files {
		if !file.IsDir() {

			t.Log(file.Name())

			content, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				t.Errorf("读取文件 %s 失败：%v", file.Name(), err)
			}
			tester(content)

		}
	}
}
