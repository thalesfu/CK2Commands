package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏剌多姞利呬SuratgarhBarony struct {
	feud.BaseBarony
}

var BSuratgarh苏剌多姞利呬 feud.Barony = &苏剌多姞利呬SuratgarhBarony{}

func init() {
	f := BSuratgarh苏剌多姞利呬.(*苏剌多姞利呬SuratgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suratgarh",
		TitleName: "苏剌多姞利呬",
		TitleCode: "b_suratgarh",
	}
}
