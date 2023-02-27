package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普费芬根PfeffingenBarony struct {
	feud.BaseBarony
}

var BPfeffingen普费芬根 feud.Barony = &普费芬根PfeffingenBarony{}

func init() {
    f := BPfeffingen普费芬根.(*普费芬根PfeffingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pfeffingen",
		TitleName: "普费芬根",
		TitleCode: "b_pfeffingen",
	}
}
