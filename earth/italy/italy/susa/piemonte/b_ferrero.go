package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费雷罗FerreroBarony struct {
	feud.BaseBarony
}

var BFerrero费雷罗 feud.Barony = &费雷罗FerreroBarony{}

func init() {
    f := BFerrero费雷罗.(*费雷罗FerreroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ferrero",
		TitleName: "费雷罗",
		TitleCode: "b_ferrero",
	}
}
