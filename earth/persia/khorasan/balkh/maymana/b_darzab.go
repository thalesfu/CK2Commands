package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜尔札卜DarzabBarony struct {
	feud.BaseBarony
}

var BDarzab杜尔札卜 feud.Barony = &杜尔札卜DarzabBarony{}

func init() {
    f := BDarzab杜尔札卜.(*杜尔札卜DarzabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darzab",
		TitleName: "杜尔札卜",
		TitleCode: "b_darzab",
	}
}
