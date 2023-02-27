package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊兰茨IllanzBarony struct {
	feud.BaseBarony
}

var BIllanz伊兰茨 feud.Barony = &伊兰茨IllanzBarony{}

func init() {
    f := BIllanz伊兰茨.(*伊兰茨IllanzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illanz",
		TitleName: "伊兰茨",
		TitleCode: "b_illanz",
	}
}
