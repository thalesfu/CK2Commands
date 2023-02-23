package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 建提底KantitBarony struct {
	feud.BaseBarony
}

var BKantit建提底 feud.Barony = &建提底KantitBarony{}

func init() {
	f := BKantit建提底.(*建提底KantitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kantit",
		TitleName: "建提底",
		TitleCode: "b_kantit",
	}
}
