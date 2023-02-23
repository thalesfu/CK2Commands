package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔克利IlkleyBarony struct {
	feud.BaseBarony
}

var BIlkley伊尔克利 feud.Barony = &伊尔克利IlkleyBarony{}

func init() {
	f := BIlkley伊尔克利.(*伊尔克利IlkleyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilkley",
		TitleName: "伊尔克利",
		TitleCode: "b_ilkley",
	}
}
