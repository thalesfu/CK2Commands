package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝卡拉ElbekaraBarony struct {
	feud.BaseBarony
}

var BElbekara贝卡拉 feud.Barony = &贝卡拉ElbekaraBarony{}

func init() {
	f := BElbekara贝卡拉.(*贝卡拉ElbekaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elbekara",
		TitleName: "贝卡拉",
		TitleCode: "b_elbekara",
	}
}
