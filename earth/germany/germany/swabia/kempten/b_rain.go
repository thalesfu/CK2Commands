package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖恩RainBarony struct {
	feud.BaseBarony
}

var BRain赖恩 feud.Barony = &赖恩RainBarony{}

func init() {
    f := BRain赖恩.(*赖恩RainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rain",
		TitleName: "赖恩",
		TitleCode: "b_rain",
	}
}
