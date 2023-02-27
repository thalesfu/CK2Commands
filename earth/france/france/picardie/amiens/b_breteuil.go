package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布勒特伊BreteuilBarony struct {
	feud.BaseBarony
}

var BBreteuil布勒特伊 feud.Barony = &布勒特伊BreteuilBarony{}

func init() {
    f := BBreteuil布勒特伊.(*布勒特伊BreteuilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "breteuil",
		TitleName: "布勒特伊",
		TitleCode: "b_breteuil",
	}
}
