package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊德尔IdarBarony struct {
	feud.BaseBarony
}

var BIdar伊德尔 feud.Barony = &伊德尔IdarBarony{}

func init() {
    f := BIdar伊德尔.(*伊德尔IdarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "idar",
		TitleName: "伊德尔",
		TitleCode: "b_idar",
	}
}
