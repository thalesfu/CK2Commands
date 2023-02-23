package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克兰ClunBarony struct {
	feud.BaseBarony
}

var BClun克兰 feud.Barony = &克兰ClunBarony{}

func init() {
	f := BClun克兰.(*克兰ClunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clun",
		TitleName: "克兰",
		TitleCode: "b_clun",
	}
}
