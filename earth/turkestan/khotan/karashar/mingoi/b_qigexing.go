package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 七个星QigexingBarony struct {
	feud.BaseBarony
}

var BQigexing七个星 feud.Barony = &七个星QigexingBarony{}

func init() {
    f := BQigexing七个星.(*七个星QigexingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qigexing",
		TitleName: "七个星",
		TitleCode: "b_qigexing",
	}
}
