package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙罗勒CharollesBarony struct {
	feud.BaseBarony
}

var BCharolles沙罗勒 feud.Barony = &沙罗勒CharollesBarony{}

func init() {
	f := BCharolles沙罗勒.(*沙罗勒CharollesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charolles",
		TitleName: "沙罗勒",
		TitleCode: "b_charolles",
	}
}
