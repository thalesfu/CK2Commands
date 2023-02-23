package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯纳亚KrasnayaBarony struct {
	feud.BaseBarony
}

var BKrasnaya克拉斯纳亚 feud.Barony = &克拉斯纳亚KrasnayaBarony{}

func init() {
	f := BKrasnaya克拉斯纳亚.(*克拉斯纳亚KrasnayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasnaya",
		TitleName: "克拉斯纳亚",
		TitleCode: "b_krasnaya",
	}
}
