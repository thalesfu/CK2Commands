package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞夫鲁SefrouBarony struct {
	feud.BaseBarony
}

var BSefrou塞夫鲁 feud.Barony = &塞夫鲁SefrouBarony{}

func init() {
	f := BSefrou塞夫鲁.(*塞夫鲁SefrouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sefrou",
		TitleName: "塞夫鲁",
		TitleCode: "b_sefrou",
	}
}
