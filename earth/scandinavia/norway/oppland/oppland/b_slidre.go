package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯利德勒SlidreBarony struct {
	feud.BaseBarony
}

var BSlidre斯利德勒 feud.Barony = &斯利德勒SlidreBarony{}

func init() {
	f := BSlidre斯利德勒.(*斯利德勒SlidreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slidre",
		TitleName: "斯利德勒",
		TitleCode: "b_slidre",
	}
}
