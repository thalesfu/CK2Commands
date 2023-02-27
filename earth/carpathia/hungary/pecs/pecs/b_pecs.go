package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩奇PecsBarony struct {
	feud.BaseBarony
}

var BPecs佩奇 feud.Barony = &佩奇PecsBarony{}

func init() {
    f := BPecs佩奇.(*佩奇PecsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pecs",
		TitleName: "佩奇",
		TitleCode: "b_pecs",
	}
}
