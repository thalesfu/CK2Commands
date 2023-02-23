package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒什塔尔ShushtarBarony struct {
	feud.BaseBarony
}

var BShushtar舒什塔尔 feud.Barony = &舒什塔尔ShushtarBarony{}

func init() {
	f := BShushtar舒什塔尔.(*舒什塔尔ShushtarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shushtar",
		TitleName: "舒什塔尔",
		TitleCode: "b_shushtar",
	}
}
