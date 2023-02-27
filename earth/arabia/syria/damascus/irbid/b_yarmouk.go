package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶尔穆克YarmoukBarony struct {
	feud.BaseBarony
}

var BYarmouk耶尔穆克 feud.Barony = &耶尔穆克YarmoukBarony{}

func init() {
    f := BYarmouk耶尔穆克.(*耶尔穆克YarmoukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yarmouk",
		TitleName: "耶尔穆克",
		TitleCode: "b_yarmouk",
	}
}
