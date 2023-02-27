package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济赫ZikhBarony struct {
	feud.BaseBarony
}

var BZikh济赫 feud.Barony = &济赫ZikhBarony{}

func init() {
    f := BZikh济赫.(*济赫ZikhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zikh",
		TitleName: "济赫",
		TitleCode: "b_zikh",
	}
}
