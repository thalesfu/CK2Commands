package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩洛皮亚PelopiaBarony struct {
	feud.BaseBarony
}

var BPelopia佩洛皮亚 feud.Barony = &佩洛皮亚PelopiaBarony{}

func init() {
	f := BPelopia佩洛皮亚.(*佩洛皮亚PelopiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pelopia",
		TitleName: "佩洛皮亚",
		TitleCode: "b_pelopia",
	}
}
