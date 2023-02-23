package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本德雷利VendrellBarony struct {
	feud.BaseBarony
}

var BVendrell本德雷利 feud.Barony = &本德雷利VendrellBarony{}

func init() {
	f := BVendrell本德雷利.(*本德雷利VendrellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vendrell",
		TitleName: "本德雷利",
		TitleCode: "b_vendrell",
	}
}
