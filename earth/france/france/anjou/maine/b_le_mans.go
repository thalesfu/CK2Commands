package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒芒Le_mansBarony struct {
	feud.BaseBarony
}

var BLe_mans勒芒 feud.Barony = &勒芒Le_mansBarony{}

func init() {
    f := BLe_mans勒芒.(*勒芒Le_mansBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "le_mans",
		TitleName: "勒芒",
		TitleCode: "b_le_mans",
	}
}
