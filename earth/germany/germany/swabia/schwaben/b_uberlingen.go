package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于伯林根UberlingenBarony struct {
	feud.BaseBarony
}

var BUberlingen于伯林根 feud.Barony = &于伯林根UberlingenBarony{}

func init() {
	f := BUberlingen于伯林根.(*于伯林根UberlingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uberlingen",
		TitleName: "于伯林根",
		TitleCode: "b_uberlingen",
	}
}
