package openrtb_v2_5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func Test_Request(t *testing.T) {
	a := new(OpenRTB26)
	reqs := getData("testdata/request")
	for filename, val := range reqs {
		obj25, err := a.UnmarshalRequest([]byte(val))
		assert.Nil(t, err, filename)

		r, err := a.MarshalRequest(obj25)
		t.Log(string(r))
		assert.Nil(t, err, filename)

		obj26 := a.TransformReqTo(obj25)
		assert.Nil(t, err, filename)

		obj25 = a.TransformReqFrom(obj26)
		assert.Nil(t, err, filename)

		byte25, err := a.DenormalizeRequest(obj26)
		assert.Nil(t, err, filename, err)

		obj26, err = a.NormalizeRequest(byte25)
		assert.Nil(t, err, filename)
	}
}
func Test_Response(t *testing.T) {
	a := new(OpenRTB26)
	resps := getData("testdata/response")
	for filename, val := range resps {
		obj25, err := a.UnmarshalResponse([]byte(val))
		assert.Nil(t, err, filename)

		r, err := a.MarshalResponse(obj25)
		t.Log(string(r))
		assert.Nil(t, err, filename)

		obj26 := a.TransformRespTo(obj25)
		assert.Nil(t, err, filename)

		obj25 = a.TransformRespFrom(obj26)
		assert.Nil(t, err, filename)

		byte25, err := a.DenormalizeResponse(obj26)
		assert.Nil(t, err, filename, err)

		obj26, err = a.NormalizeResponse(byte25)
		assert.Nil(t, err, filename)
	}
}

func getData(dirname string) (data map[string]string) {
	_, curFilename, _, _ := runtime.Caller(0)

	dirName := filepath.Join(path.Dir(curFilename), dirname)
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	data = make(map[string]string, 16)
	// 读取文件夹中的文件和子文件夹信息
	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// 遍历文件信息
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		// 构建文件路径
		filePath := filepath.Join(dirName, file.Name())

		// 读取文件内容
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		data[file.Name()] = string(content)

	}
	return
}
