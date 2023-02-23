package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍迦尔ShekhsarBarony struct {
	feud.BaseBarony
}

var BShekhsar舍迦尔 feud.Barony = &舍迦尔ShekhsarBarony{}

func init() {
	f := BShekhsar舍迦尔.(*舍迦尔ShekhsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shekhsar",
		TitleName: "舍迦尔",
		TitleCode: "b_shekhsar",
	}
}
