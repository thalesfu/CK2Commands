package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆捷MoutierBarony struct {
	feud.BaseBarony
}

var BMoutier穆捷 feud.Barony = &穆捷MoutierBarony{}

func init() {
    f := BMoutier穆捷.(*穆捷MoutierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moutier",
		TitleName: "穆捷",
		TitleCode: "b_moutier",
	}
}
