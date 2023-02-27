package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰尔姆TermesBarony struct {
	feud.BaseBarony
}

var BTermes泰尔姆 feud.Barony = &泰尔姆TermesBarony{}

func init() {
    f := BTermes泰尔姆.(*泰尔姆TermesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "termes",
		TitleName: "泰尔姆",
		TitleCode: "b_termes",
	}
}
