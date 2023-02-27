package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨浪伽补罗SarangpurBarony struct {
	feud.BaseBarony
}

var BSarangpur萨浪伽补罗 feud.Barony = &萨浪伽补罗SarangpurBarony{}

func init() {
    f := BSarangpur萨浪伽补罗.(*萨浪伽补罗SarangpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarangpur",
		TitleName: "萨浪伽补罗",
		TitleCode: "b_sarangpur",
	}
}
