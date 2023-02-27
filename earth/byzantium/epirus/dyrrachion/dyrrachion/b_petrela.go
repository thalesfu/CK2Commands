package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩特雷勒PetrelaBarony struct {
	feud.BaseBarony
}

var BPetrela佩特雷勒 feud.Barony = &佩特雷勒PetrelaBarony{}

func init() {
    f := BPetrela佩特雷勒.(*佩特雷勒PetrelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petrela",
		TitleName: "佩特雷勒",
		TitleCode: "b_petrela",
	}
}
