package yatenga

import (
	"github.com/thalesfu/CK2Commands/earth/mali/mali/yatenga/wagadougou"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/yatenga/yatenga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YatengaDuke interface {
	feud.Duke
	CWagadougou瓦加杜古() wagadougou.WagadougouCounty
	CYatenga亚滕加() yatenga.YatengaCounty
}

type 亚滕加YatengaDuke struct {
	feud.BaseDuke
	瓦加杜古Wagadougou wagadougou.WagadougouCounty
	亚滕加Yatenga     yatenga.YatengaCounty
}

func (d *亚滕加YatengaDuke) CWagadougou瓦加杜古() wagadougou.WagadougouCounty {
	return d.瓦加杜古Wagadougou
}

func (d *亚滕加YatengaDuke) CYatenga亚滕加() yatenga.YatengaCounty {
	return d.亚滕加Yatenga
}

var DYatenga亚滕加 YatengaDuke = &亚滕加YatengaDuke{}

func init() {
	f := DYatenga亚滕加.(*亚滕加YatengaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "yatenga",
		TitleName: "亚滕加",
		TitleCode: "d_yatenga",
		Counties:  map[string]feud.County{},
	}

	f.瓦加杜古Wagadougou = wagadougou.CWagadougou瓦加杜古
	f.瓦加杜古Wagadougou.SetParent(f)

	f.亚滕加Yatenga = yatenga.CYatenga亚滕加
	f.亚滕加Yatenga.SetParent(f)

}
