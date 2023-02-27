package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊曼达YmandaBarony struct {
	feud.BaseBarony
}

var BYmanda伊曼达 feud.Barony = &伊曼达YmandaBarony{}

func init() {
    f := BYmanda伊曼达.(*伊曼达YmandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ymanda",
		TitleName: "伊曼达",
		TitleCode: "b_ymanda",
	}
}
