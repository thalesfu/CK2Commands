package leeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍本SherburnBarony struct {
	feud.BaseBarony
}

var BSherburn舍本 feud.Barony = &舍本SherburnBarony{}

func init() {
	f := BSherburn舍本.(*舍本SherburnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sherburn",
		TitleName: "舍本",
		TitleCode: "b_sherburn",
	}
}
