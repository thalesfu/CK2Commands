package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布佉罗PokharaBarony struct {
	feud.BaseBarony
}

var BPokhara布佉罗 feud.Barony = &布佉罗PokharaBarony{}

func init() {
	f := BPokhara布佉罗.(*布佉罗PokharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pokhara",
		TitleName: "布佉罗",
		TitleCode: "b_pokhara",
	}
}
