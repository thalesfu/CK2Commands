package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩戈埃斯PegoesBarony struct {
	feud.BaseBarony
}

var BPegoes佩戈埃斯 feud.Barony = &佩戈埃斯PegoesBarony{}

func init() {
    f := BPegoes佩戈埃斯.(*佩戈埃斯PegoesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pegoes",
		TitleName: "佩戈埃斯",
		TitleCode: "b_pegoes",
	}
}
