package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈丁HattinBarony struct {
	feud.BaseBarony
}

var BHattin哈丁 feud.Barony = &哈丁HattinBarony{}

func init() {
	f := BHattin哈丁.(*哈丁HattinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hattin",
		TitleName: "哈丁",
		TitleCode: "b_hattin",
	}
}
