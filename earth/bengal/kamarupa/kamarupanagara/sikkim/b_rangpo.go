package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗格波RangpoBarony struct {
	feud.BaseBarony
}

var BRangpo朗格波 feud.Barony = &朗格波RangpoBarony{}

func init() {
    f := BRangpo朗格波.(*朗格波RangpoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rangpo",
		TitleName: "朗格波",
		TitleCode: "b_rangpo",
	}
}
