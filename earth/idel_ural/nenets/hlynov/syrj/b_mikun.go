package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米昆MikunBarony struct {
	feud.BaseBarony
}

var BMikun米昆 feud.Barony = &米昆MikunBarony{}

func init() {
    f := BMikun米昆.(*米昆MikunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mikun",
		TitleName: "米昆",
		TitleCode: "b_mikun",
	}
}
