package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔戈YrgoBarony struct {
	feud.BaseBarony
}

var BYrgo伊尔戈 feud.Barony = &伊尔戈YrgoBarony{}

func init() {
    f := BYrgo伊尔戈.(*伊尔戈YrgoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yrgo",
		TitleName: "伊尔戈",
		TitleCode: "b_yrgo",
	}
}
