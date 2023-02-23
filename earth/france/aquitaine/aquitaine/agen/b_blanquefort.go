package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布朗克福BlanquefortBarony struct {
	feud.BaseBarony
}

var BBlanquefort布朗克福 feud.Barony = &布朗克福BlanquefortBarony{}

func init() {
	f := BBlanquefort布朗克福.(*布朗克福BlanquefortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blanquefort",
		TitleName: "布朗克福",
		TitleCode: "b_blanquefort",
	}
}
