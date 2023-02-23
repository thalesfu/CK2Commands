package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊拉洪IrahunBarony struct {
	feud.BaseBarony
}

var BIrahun伊拉洪 feud.Barony = &伊拉洪IrahunBarony{}

func init() {
	f := BIrahun伊拉洪.(*伊拉洪IrahunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irahun",
		TitleName: "伊拉洪",
		TitleCode: "b_irahun",
	}
}
