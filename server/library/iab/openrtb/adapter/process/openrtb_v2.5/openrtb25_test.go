package openrtb_v2_5

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.5"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func Test_deepCopySource(t *testing.T) {
	var ortb25source *openrtb_v2_5.Source
	var ortb26source *openrtb_v2_6.Source

	json.Unmarshal([]byte(j25s), &ortb25source)
	//json.Unmarshal([]byte(j26s), &ortb26source)
	ortb26source = deepCopySource(ortb25source)
	fmt.Println(ortb26source.SupplyChain.Ver)
}
func Test_Source2(t *testing.T) {
	var ortb25source = new(openrtb_v2_5.Source)
	var ortb26source = new(openrtb_v2_6.Source)

	//json.Unmarshal([]byte(j25s), ortb25source)
	json.Unmarshal([]byte(j26s), ortb26source)
	ortb25source = deepCopySource2(ortb26source)
	fmt.Println(string(ortb25source.Ext))
}

func Test_SupplyChain(t *testing.T) {
	s25 := &openrtb_v2_5.Source{}
	s26 := &openrtb_v2_6.Source{}

	t.Run("26的ext放到26的supplychain", func(t *testing.T) {
		json.Unmarshal([]byte(`{"ext":{"schain":{"ver":"123"},"key6":"val6"}}`), s26)
		supplyChainAdapte(s26)
		assert.NotNil(t, s26.SupplyChain)
		assert.NotContains(t, string(s26.Ext), "123")
	})
	t.Run("26的supplychain放到26的ext", func(t *testing.T) {
		json.Unmarshal([]byte(`{"schain":{"ver":"123"},"ext":{"key6":"val6"}}`), s26)
		supplyChainAdapte2(s26)
		assert.Nil(t, s26.SupplyChain)
		assert.Contains(t, string(s26.Ext), "123")
	})
	t.Run("25的ext放到26的supplychain", func(t *testing.T) {
		json.Unmarshal([]byte(`{"ext":{"schain":{"ver":"123"},"key6":"val6"}}`), s25)
		json.Unmarshal([]byte(`{"ext":{"key6":"val6"}}`), s26)
		supplyChain25to26(s25, s26)
		assert.Contains(t, string(s25.Ext), s26.SupplyChain.Ver)
	})
	t.Run("26的supplychain放到26的ext", func(t *testing.T) {
		json.Unmarshal([]byte(`{"ext":{"key6":"val6"}}`), s25)
		json.Unmarshal([]byte(`{"schain":{"ver":"123"},"ext":{"key6":"val6"}}`), s26)
		supplyChain26to25(s26, s25)
		assert.Contains(t, string(s25.Ext), s26.SupplyChain.Ver)
	})
}

func TestOpenRTB25_Request(t *testing.T) {
	a := new(OpenRTB25)
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
func TestOpenRTB25_Response(t *testing.T) {
	a := new(OpenRTB25)
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

var j25s = `{
	"fd":1,
	"tid":"2",
	"pchain":"pchain1",
	"ext":{
		"key5":"val5",
		"schain":{
			"ver":"123"
		}
	}
}`
var j26s = `{
	"fd":1,
	"tid":"2",
	"pchain":"pchain1",
	"schain":{
			"ver":"123"
	},
	"ext":{
		"key6":"val6"
	}
}`
