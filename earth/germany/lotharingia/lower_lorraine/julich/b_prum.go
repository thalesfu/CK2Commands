package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普吕姆PrumBarony struct {
	feud.BaseBarony
}

var BPrum普吕姆 feud.Barony = &普吕姆PrumBarony{}

func init() {
    f := BPrum普吕姆.(*普吕姆PrumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prum",
		TitleName: "普吕姆",
		TitleCode: "b_prum",
	}
}
