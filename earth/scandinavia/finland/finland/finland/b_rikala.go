package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里卡拉RikalaBarony struct {
	feud.BaseBarony
}

var BRikala里卡拉 feud.Barony = &里卡拉RikalaBarony{}

func init() {
    f := BRikala里卡拉.(*里卡拉RikalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rikala",
		TitleName: "里卡拉",
		TitleCode: "b_rikala",
	}
}
