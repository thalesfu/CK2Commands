package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔努瓦SaulnoisBarony struct {
	feud.BaseBarony
}

var BSaulnois索尔努瓦 feud.Barony = &索尔努瓦SaulnoisBarony{}

func init() {
    f := BSaulnois索尔努瓦.(*索尔努瓦SaulnoisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saulnois",
		TitleName: "索尔努瓦",
		TitleCode: "b_saulnois",
	}
}
