package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮山PishanBarony struct {
	feud.BaseBarony
}

var BPishan皮山 feud.Barony = &皮山PishanBarony{}

func init() {
    f := BPishan皮山.(*皮山PishanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pishan",
		TitleName: "皮山",
		TitleCode: "b_pishan",
	}
}
