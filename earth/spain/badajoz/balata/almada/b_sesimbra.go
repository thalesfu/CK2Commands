package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛辛布拉SesimbraBarony struct {
	feud.BaseBarony
}

var BSesimbra赛辛布拉 feud.Barony = &赛辛布拉SesimbraBarony{}

func init() {
	f := BSesimbra赛辛布拉.(*赛辛布拉SesimbraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sesimbra",
		TitleName: "赛辛布拉",
		TitleCode: "b_sesimbra",
	}
}
