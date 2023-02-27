package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥帕图夫OpatowBarony struct {
	feud.BaseBarony
}

var BOpatow奥帕图夫 feud.Barony = &奥帕图夫OpatowBarony{}

func init() {
    f := BOpatow奥帕图夫.(*奥帕图夫OpatowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "opatow",
		TitleName: "奥帕图夫",
		TitleCode: "b_opatow",
	}
}
