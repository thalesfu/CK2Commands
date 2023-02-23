package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼遮罗ManchalBarony struct {
	feud.BaseBarony
}

var BManchal曼遮罗 feud.Barony = &曼遮罗ManchalBarony{}

func init() {
	f := BManchal曼遮罗.(*曼遮罗ManchalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manchal",
		TitleName: "曼遮罗",
		TitleCode: "b_manchal",
	}
}
