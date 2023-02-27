package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥鲍乌伊AbaujBarony struct {
	feud.BaseBarony
}

var BAbauj奥鲍乌伊 feud.Barony = &奥鲍乌伊AbaujBarony{}

func init() {
    f := BAbauj奥鲍乌伊.(*奥鲍乌伊AbaujBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abauj",
		TitleName: "奥鲍乌伊",
		TitleCode: "b_abauj",
	}
}
