package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥列马OlemaBarony struct {
	feud.BaseBarony
}

var BOlema奥列马 feud.Barony = &奥列马OlemaBarony{}

func init() {
    f := BOlema奥列马.(*奥列马OlemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olema",
		TitleName: "奥列马",
		TitleCode: "b_olema",
	}
}
