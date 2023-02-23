package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛尔施LorschBarony struct {
	feud.BaseBarony
}

var BLorsch洛尔施 feud.Barony = &洛尔施LorschBarony{}

func init() {
	f := BLorsch洛尔施.(*洛尔施LorschBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lorsch",
		TitleName: "洛尔施",
		TitleCode: "b_lorsch",
	}
}
