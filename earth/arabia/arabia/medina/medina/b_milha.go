package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米哈MilhaBarony struct {
	feud.BaseBarony
}

var BMilha米哈 feud.Barony = &米哈MilhaBarony{}

func init() {
    f := BMilha米哈.(*米哈MilhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "milha",
		TitleName: "米哈",
		TitleCode: "b_milha",
	}
}
