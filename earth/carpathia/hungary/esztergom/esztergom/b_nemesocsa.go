package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈迈绍乔NemesocsaBarony struct {
	feud.BaseBarony
}

var BNemesocsa奈迈绍乔 feud.Barony = &奈迈绍乔NemesocsaBarony{}

func init() {
    f := BNemesocsa奈迈绍乔.(*奈迈绍乔NemesocsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nemesocsa",
		TitleName: "奈迈绍乔",
		TitleCode: "b_nemesocsa",
	}
}
