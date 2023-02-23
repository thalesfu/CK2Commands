package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔扎尔迪特TazaditBarony struct {
	feud.BaseBarony
}

var BTazadit塔扎尔迪特 feud.Barony = &塔扎尔迪特TazaditBarony{}

func init() {
	f := BTazadit塔扎尔迪特.(*塔扎尔迪特TazaditBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tazadit",
		TitleName: "塔扎尔迪特",
		TitleCode: "b_tazadit",
	}
}
