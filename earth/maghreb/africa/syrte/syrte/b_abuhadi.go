package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布哈迪AbuhadiBarony struct {
	feud.BaseBarony
}

var BAbuhadi阿布哈迪 feud.Barony = &阿布哈迪AbuhadiBarony{}

func init() {
    f := BAbuhadi阿布哈迪.(*阿布哈迪AbuhadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abuhadi",
		TitleName: "阿布哈迪",
		TitleCode: "b_abuhadi",
	}
}
