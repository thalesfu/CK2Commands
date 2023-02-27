package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里什蒙RichemontBarony struct {
	feud.BaseBarony
}

var BRichemont里什蒙 feud.Barony = &里什蒙RichemontBarony{}

func init() {
    f := BRichemont里什蒙.(*里什蒙RichemontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "richemont",
		TitleName: "里什蒙",
		TitleCode: "b_richemont",
	}
}
