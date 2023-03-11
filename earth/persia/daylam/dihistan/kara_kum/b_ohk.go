package kara_kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥赫OhkBarony struct {
	feud.BaseBarony
}

var BOhk奥赫 feud.Barony = &奥赫OhkBarony{}

func init() {
    f := BOhk奥赫.(*奥赫OhkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ohk",
		TitleName: "奥赫",
		TitleCode: "b_ohk",
	}
}
