package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 底比斯EgypthebesBarony struct {
	feud.BaseBarony
}

var BEgypthebes底比斯 feud.Barony = &底比斯EgypthebesBarony{}

func init() {
    f := BEgypthebes底比斯.(*底比斯EgypthebesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egypthebes",
		TitleName: "底比斯",
		TitleCode: "b_egypthebes",
	}
}
