package tui

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍尔加奇SaurgachiBarony struct {
	feud.BaseBarony
}

var BSaurgachi绍尔加奇 feud.Barony = &绍尔加奇SaurgachiBarony{}

func init() {
	f := BSaurgachi绍尔加奇.(*绍尔加奇SaurgachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saurgachi",
		TitleName: "绍尔加奇",
		TitleCode: "b_saurgachi",
	}
}
