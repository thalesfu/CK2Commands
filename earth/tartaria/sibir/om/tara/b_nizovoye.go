package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼佐沃耶NizovoyeBarony struct {
	feud.BaseBarony
}

var BNizovoye尼佐沃耶 feud.Barony = &尼佐沃耶NizovoyeBarony{}

func init() {
	f := BNizovoye尼佐沃耶.(*尼佐沃耶NizovoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizovoye",
		TitleName: "尼佐沃耶",
		TitleCode: "b_nizovoye",
	}
}
