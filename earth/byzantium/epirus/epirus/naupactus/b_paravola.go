package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕拉沃拉ParavolaBarony struct {
	feud.BaseBarony
}

var BParavola帕拉沃拉 feud.Barony = &帕拉沃拉ParavolaBarony{}

func init() {
    f := BParavola帕拉沃拉.(*帕拉沃拉ParavolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paravola",
		TitleName: "帕拉沃拉",
		TitleCode: "b_paravola",
	}
}
