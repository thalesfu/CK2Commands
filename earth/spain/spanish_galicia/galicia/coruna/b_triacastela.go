package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里亚卡斯特拉TriacastelaBarony struct {
	feud.BaseBarony
}

var BTriacastela特里亚卡斯特拉 feud.Barony = &特里亚卡斯特拉TriacastelaBarony{}

func init() {
    f := BTriacastela特里亚卡斯特拉.(*特里亚卡斯特拉TriacastelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "triacastela",
		TitleName: "特里亚卡斯特拉",
		TitleCode: "b_triacastela",
	}
}
