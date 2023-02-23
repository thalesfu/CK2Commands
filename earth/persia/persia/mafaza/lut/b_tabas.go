package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔巴斯TabasBarony struct {
	feud.BaseBarony
}

var BTabas塔巴斯 feud.Barony = &塔巴斯TabasBarony{}

func init() {
	f := BTabas塔巴斯.(*塔巴斯TabasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabas",
		TitleName: "塔巴斯",
		TitleCode: "b_tabas",
	}
}
