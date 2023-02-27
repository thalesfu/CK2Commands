package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芬斯特伦FinstromBarony struct {
	feud.BaseBarony
}

var BFinstrom芬斯特伦 feud.Barony = &芬斯特伦FinstromBarony{}

func init() {
    f := BFinstrom芬斯特伦.(*芬斯特伦FinstromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "finstrom",
		TitleName: "芬斯特伦",
		TitleCode: "b_finstrom",
	}
}
