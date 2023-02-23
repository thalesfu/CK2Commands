package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里卡拉TrikkalaBarony struct {
	feud.BaseBarony
}

var BTrikkala特里卡拉 feud.Barony = &特里卡拉TrikkalaBarony{}

func init() {
	f := BTrikkala特里卡拉.(*特里卡拉TrikkalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trikkala",
		TitleName: "特里卡拉",
		TitleCode: "b_trikkala",
	}
}
