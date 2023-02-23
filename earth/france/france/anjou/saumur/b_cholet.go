package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍莱CholetBarony struct {
	feud.BaseBarony
}

var BCholet绍莱 feud.Barony = &绍莱CholetBarony{}

func init() {
	f := BCholet绍莱.(*绍莱CholetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cholet",
		TitleName: "绍莱",
		TitleCode: "b_cholet",
	}
}
