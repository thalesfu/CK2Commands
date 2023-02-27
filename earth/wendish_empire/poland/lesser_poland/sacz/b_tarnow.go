package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔努夫TarnowBarony struct {
	feud.BaseBarony
}

var BTarnow塔尔努夫 feud.Barony = &塔尔努夫TarnowBarony{}

func init() {
    f := BTarnow塔尔努夫.(*塔尔努夫TarnowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarnow",
		TitleName: "塔尔努夫",
		TitleCode: "b_tarnow",
	}
}
