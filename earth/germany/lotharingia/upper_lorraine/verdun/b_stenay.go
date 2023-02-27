package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特奈StenayBarony struct {
	feud.BaseBarony
}

var BStenay斯特奈 feud.Barony = &斯特奈StenayBarony{}

func init() {
    f := BStenay斯特奈.(*斯特奈StenayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stenay",
		TitleName: "斯特奈",
		TitleCode: "b_stenay",
	}
}
