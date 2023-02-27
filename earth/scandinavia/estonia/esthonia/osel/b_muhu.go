package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆胡MuhuBarony struct {
	feud.BaseBarony
}

var BMuhu穆胡 feud.Barony = &穆胡MuhuBarony{}

func init() {
    f := BMuhu穆胡.(*穆胡MuhuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muhu",
		TitleName: "穆胡",
		TitleCode: "b_muhu",
	}
}
