package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维堡ViborgBarony struct {
	feud.BaseBarony
}

var BViborg维堡 feud.Barony = &维堡ViborgBarony{}

func init() {
	f := BViborg维堡.(*维堡ViborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viborg",
		TitleName: "维堡",
		TitleCode: "b_viborg",
	}
}
