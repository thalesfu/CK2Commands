package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔祖莱UrzuleiBarony struct {
	feud.BaseBarony
}

var BUrzulei乌尔祖莱 feud.Barony = &乌尔祖莱UrzuleiBarony{}

func init() {
    f := BUrzulei乌尔祖莱.(*乌尔祖莱UrzuleiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urzulei",
		TitleName: "乌尔祖莱",
		TitleCode: "b_urzulei",
	}
}
