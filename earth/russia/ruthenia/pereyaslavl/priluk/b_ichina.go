package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊奇纳IchinaBarony struct {
	feud.BaseBarony
}

var BIchina伊奇纳 feud.Barony = &伊奇纳IchinaBarony{}

func init() {
    f := BIchina伊奇纳.(*伊奇纳IchinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ichina",
		TitleName: "伊奇纳",
		TitleCode: "b_ichina",
	}
}
