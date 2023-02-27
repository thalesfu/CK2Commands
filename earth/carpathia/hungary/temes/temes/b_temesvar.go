package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰梅什堡TemesvarBarony struct {
	feud.BaseBarony
}

var BTemesvar泰梅什堡 feud.Barony = &泰梅什堡TemesvarBarony{}

func init() {
    f := BTemesvar泰梅什堡.(*泰梅什堡TemesvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "temesvar",
		TitleName: "泰梅什堡",
		TitleCode: "b_temesvar",
	}
}
