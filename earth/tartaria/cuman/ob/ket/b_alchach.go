package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔恰奇AlchachBarony struct {
	feud.BaseBarony
}

var BAlchach阿尔恰奇 feud.Barony = &阿尔恰奇AlchachBarony{}

func init() {
    f := BAlchach阿尔恰奇.(*阿尔恰奇AlchachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alchach",
		TitleName: "阿尔恰奇",
		TitleCode: "b_alchach",
	}
}
