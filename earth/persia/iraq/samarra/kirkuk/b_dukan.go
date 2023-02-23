package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜坎DukanBarony struct {
	feud.BaseBarony
}

var BDukan杜坎 feud.Barony = &杜坎DukanBarony{}

func init() {
	f := BDukan杜坎.(*杜坎DukanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dukan",
		TitleName: "杜坎",
		TitleCode: "b_dukan",
	}
}
