package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里切夫KrychawBarony struct {
	feud.BaseBarony
}

var BKrychaw克里切夫 feud.Barony = &克里切夫KrychawBarony{}

func init() {
	f := BKrychaw克里切夫.(*克里切夫KrychawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krychaw",
		TitleName: "克里切夫",
		TitleCode: "b_krychaw",
	}
}
