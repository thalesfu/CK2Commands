package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈科普MyequapeBarony struct {
	feud.BaseBarony
}

var BMyequape迈科普 feud.Barony = &迈科普MyequapeBarony{}

func init() {
    f := BMyequape迈科普.(*迈科普MyequapeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myequape",
		TitleName: "迈科普",
		TitleCode: "b_myequape",
	}
}
