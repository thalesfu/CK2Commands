package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔巴迪AlarbadiBarony struct {
	feud.BaseBarony
}

var BAlarbadi阿尔巴迪 feud.Barony = &阿尔巴迪AlarbadiBarony{}

func init() {
    f := BAlarbadi阿尔巴迪.(*阿尔巴迪AlarbadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alarbadi",
		TitleName: "阿尔巴迪",
		TitleCode: "b_alarbadi",
	}
}
