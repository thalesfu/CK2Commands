package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊泽尔昂IzerhanBarony struct {
	feud.BaseBarony
}

var BIzerhan伊泽尔昂 feud.Barony = &伊泽尔昂IzerhanBarony{}

func init() {
    f := BIzerhan伊泽尔昂.(*伊泽尔昂IzerhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "izerhan",
		TitleName: "伊泽尔昂",
		TitleCode: "b_izerhan",
	}
}
