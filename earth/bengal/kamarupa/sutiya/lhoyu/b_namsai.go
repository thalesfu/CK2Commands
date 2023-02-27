package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南赛NamsaiBarony struct {
	feud.BaseBarony
}

var BNamsai南赛 feud.Barony = &南赛NamsaiBarony{}

func init() {
    f := BNamsai南赛.(*南赛NamsaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namsai",
		TitleName: "南赛",
		TitleCode: "b_namsai",
	}
}
