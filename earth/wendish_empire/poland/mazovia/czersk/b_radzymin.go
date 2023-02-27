package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉济明RadzyminBarony struct {
	feud.BaseBarony
}

var BRadzymin拉济明 feud.Barony = &拉济明RadzyminBarony{}

func init() {
    f := BRadzymin拉济明.(*拉济明RadzyminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radzymin",
		TitleName: "拉济明",
		TitleCode: "b_radzymin",
	}
}
