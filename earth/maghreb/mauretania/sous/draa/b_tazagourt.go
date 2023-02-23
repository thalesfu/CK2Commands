package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔扎古尔特TazagourtBarony struct {
	feud.BaseBarony
}

var BTazagourt塔扎古尔特 feud.Barony = &塔扎古尔特TazagourtBarony{}

func init() {
	f := BTazagourt塔扎古尔特.(*塔扎古尔特TazagourtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tazagourt",
		TitleName: "塔扎古尔特",
		TitleCode: "b_tazagourt",
	}
}
