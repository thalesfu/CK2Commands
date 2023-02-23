package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱宁根LeiningenBarony struct {
	feud.BaseBarony
}

var BLeiningen莱宁根 feud.Barony = &莱宁根LeiningenBarony{}

func init() {
	f := BLeiningen莱宁根.(*莱宁根LeiningenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leiningen",
		TitleName: "莱宁根",
		TitleCode: "b_leiningen",
	}
}
