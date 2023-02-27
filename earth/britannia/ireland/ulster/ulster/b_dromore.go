package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德罗莫尔DromoreBarony struct {
	feud.BaseBarony
}

var BDromore德罗莫尔 feud.Barony = &德罗莫尔DromoreBarony{}

func init() {
    f := BDromore德罗莫尔.(*德罗莫尔DromoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dromore",
		TitleName: "德罗莫尔",
		TitleCode: "b_dromore",
	}
}
