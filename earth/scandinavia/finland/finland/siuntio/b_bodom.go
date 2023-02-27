package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博多姆BodomBarony struct {
	feud.BaseBarony
}

var BBodom博多姆 feud.Barony = &博多姆BodomBarony{}

func init() {
    f := BBodom博多姆.(*博多姆BodomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodom",
		TitleName: "博多姆",
		TitleCode: "b_bodom",
	}
}
