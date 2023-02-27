package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德布雷塞特DebrezeyitBarony struct {
	feud.BaseBarony
}

var BDebrezeyit德布雷塞特 feud.Barony = &德布雷塞特DebrezeyitBarony{}

func init() {
    f := BDebrezeyit德布雷塞特.(*德布雷塞特DebrezeyitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debrezeyit",
		TitleName: "德布雷塞特",
		TitleCode: "b_debrezeyit",
	}
}
