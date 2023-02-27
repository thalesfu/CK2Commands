package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昭莫多ZuunmodBarony struct {
	feud.BaseBarony
}

var BZuunmod昭莫多 feud.Barony = &昭莫多ZuunmodBarony{}

func init() {
    f := BZuunmod昭莫多.(*昭莫多ZuunmodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuunmod",
		TitleName: "昭莫多",
		TitleCode: "b_zuunmod",
	}
}
