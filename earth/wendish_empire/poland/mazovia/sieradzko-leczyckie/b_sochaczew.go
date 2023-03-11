package sieradzko-leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索哈切夫SochaczewBarony struct {
	feud.BaseBarony
}

var BSochaczew索哈切夫 feud.Barony = &索哈切夫SochaczewBarony{}

func init() {
    f := BSochaczew索哈切夫.(*索哈切夫SochaczewBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sochaczew",
		TitleName: "索哈切夫",
		TitleCode: "b_sochaczew",
	}
}
