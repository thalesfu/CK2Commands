package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔切斯特IlchesterBarony struct {
	feud.BaseBarony
}

var BIlchester伊尔切斯特 feud.Barony = &伊尔切斯特IlchesterBarony{}

func init() {
    f := BIlchester伊尔切斯特.(*伊尔切斯特IlchesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilchester",
		TitleName: "伊尔切斯特",
		TitleCode: "b_ilchester",
	}
}
