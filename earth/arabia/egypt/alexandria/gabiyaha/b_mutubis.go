package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆图比斯MutubisBarony struct {
	feud.BaseBarony
}

var BMutubis穆图比斯 feud.Barony = &穆图比斯MutubisBarony{}

func init() {
    f := BMutubis穆图比斯.(*穆图比斯MutubisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mutubis",
		TitleName: "穆图比斯",
		TitleCode: "b_mutubis",
	}
}
