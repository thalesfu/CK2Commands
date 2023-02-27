package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆特纳亚MutnayaBarony struct {
	feud.BaseBarony
}

var BMutnaya穆特纳亚 feud.Barony = &穆特纳亚MutnayaBarony{}

func init() {
    f := BMutnaya穆特纳亚.(*穆特纳亚MutnayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mutnaya",
		TitleName: "穆特纳亚",
		TitleCode: "b_mutnaya",
	}
}
