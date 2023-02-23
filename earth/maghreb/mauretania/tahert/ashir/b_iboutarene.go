package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊布塔赫恩IboutareneBarony struct {
	feud.BaseBarony
}

var BIboutarene伊布塔赫恩 feud.Barony = &伊布塔赫恩IboutareneBarony{}

func init() {
	f := BIboutarene伊布塔赫恩.(*伊布塔赫恩IboutareneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iboutarene",
		TitleName: "伊布塔赫恩",
		TitleCode: "b_iboutarene",
	}
}
