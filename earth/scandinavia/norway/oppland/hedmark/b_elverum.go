package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔沃吕姆ElverumBarony struct {
	feud.BaseBarony
}

var BElverum埃尔沃吕姆 feud.Barony = &埃尔沃吕姆ElverumBarony{}

func init() {
    f := BElverum埃尔沃吕姆.(*埃尔沃吕姆ElverumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elverum",
		TitleName: "埃尔沃吕姆",
		TitleCode: "b_elverum",
	}
}
