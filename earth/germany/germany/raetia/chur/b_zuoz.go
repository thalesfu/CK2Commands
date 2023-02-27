package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚奥茨ZuozBarony struct {
	feud.BaseBarony
}

var BZuoz楚奥茨 feud.Barony = &楚奥茨ZuozBarony{}

func init() {
    f := BZuoz楚奥茨.(*楚奥茨ZuozBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuoz",
		TitleName: "楚奥茨",
		TitleCode: "b_zuoz",
	}
}
