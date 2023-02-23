package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔本加AlbengaBarony struct {
	feud.BaseBarony
}

var BAlbenga阿尔本加 feud.Barony = &阿尔本加AlbengaBarony{}

func init() {
	f := BAlbenga阿尔本加.(*阿尔本加AlbengaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albenga",
		TitleName: "阿尔本加",
		TitleCode: "b_albenga",
	}
}
