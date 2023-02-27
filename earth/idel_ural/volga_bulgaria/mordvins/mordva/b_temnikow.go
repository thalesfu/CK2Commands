package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷姆尼科夫TemnikowBarony struct {
	feud.BaseBarony
}

var BTemnikow捷姆尼科夫 feud.Barony = &捷姆尼科夫TemnikowBarony{}

func init() {
    f := BTemnikow捷姆尼科夫.(*捷姆尼科夫TemnikowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "temnikow",
		TitleName: "捷姆尼科夫",
		TitleCode: "b_temnikow",
	}
}
