package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昏驮多KhandudBarony struct {
	feud.BaseBarony
}

var BKhandud昏驮多 feud.Barony = &昏驮多KhandudBarony{}

func init() {
    f := BKhandud昏驮多.(*昏驮多KhandudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khandud",
		TitleName: "昏驮多",
		TitleCode: "b_khandud",
	}
}
