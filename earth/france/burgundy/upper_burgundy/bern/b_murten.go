package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆尔滕MurtenBarony struct {
	feud.BaseBarony
}

var BMurten穆尔滕 feud.Barony = &穆尔滕MurtenBarony{}

func init() {
    f := BMurten穆尔滕.(*穆尔滕MurtenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murten",
		TitleName: "穆尔滕",
		TitleCode: "b_murten",
	}
}
