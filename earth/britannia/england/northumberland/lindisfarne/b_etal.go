package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊特尔EtalBarony struct {
	feud.BaseBarony
}

var BEtal伊特尔 feud.Barony = &伊特尔EtalBarony{}

func init() {
    f := BEtal伊特尔.(*伊特尔EtalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etal",
		TitleName: "伊特尔",
		TitleCode: "b_etal",
	}
}
