package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥德OudhBarony struct {
	feud.BaseBarony
}

var BOudh奥德 feud.Barony = &奥德OudhBarony{}

func init() {
    f := BOudh奥德.(*奥德OudhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oudh",
		TitleName: "奥德",
		TitleCode: "b_oudh",
	}
}
