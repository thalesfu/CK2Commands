package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯泰恩斯StainesBarony struct {
	feud.BaseBarony
}

var BStaines斯泰恩斯 feud.Barony = &斯泰恩斯StainesBarony{}

func init() {
    f := BStaines斯泰恩斯.(*斯泰恩斯StainesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "staines",
		TitleName: "斯泰恩斯",
		TitleCode: "b_staines",
	}
}
