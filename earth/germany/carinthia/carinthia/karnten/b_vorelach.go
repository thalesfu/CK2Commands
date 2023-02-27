package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费拉赫VorelachBarony struct {
	feud.BaseBarony
}

var BVorelach费拉赫 feud.Barony = &费拉赫VorelachBarony{}

func init() {
    f := BVorelach费拉赫.(*费拉赫VorelachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vorelach",
		TitleName: "费拉赫",
		TitleCode: "b_vorelach",
	}
}
