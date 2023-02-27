package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米凯利MikkeliBarony struct {
	feud.BaseBarony
}

var BMikkeli米凯利 feud.Barony = &米凯利MikkeliBarony{}

func init() {
    f := BMikkeli米凯利.(*米凯利MikkeliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mikkeli",
		TitleName: "米凯利",
		TitleCode: "b_mikkeli",
	}
}
