package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般遮补罗PancapuraBarony struct {
	feud.BaseBarony
}

var BPancapura般遮补罗 feud.Barony = &般遮补罗PancapuraBarony{}

func init() {
	f := BPancapura般遮补罗.(*般遮补罗PancapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pancapura",
		TitleName: "般遮补罗",
		TitleCode: "b_pancapura",
	}
}
