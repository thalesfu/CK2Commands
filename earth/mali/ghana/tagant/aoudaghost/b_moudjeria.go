package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆杰里亚MoudjeriaBarony struct {
	feud.BaseBarony
}

var BMoudjeria穆杰里亚 feud.Barony = &穆杰里亚MoudjeriaBarony{}

func init() {
	f := BMoudjeria穆杰里亚.(*穆杰里亚MoudjeriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moudjeria",
		TitleName: "穆杰里亚",
		TitleCode: "b_moudjeria",
	}
}
