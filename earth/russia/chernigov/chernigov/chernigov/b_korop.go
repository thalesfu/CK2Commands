package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科罗普KoropBarony struct {
	feud.BaseBarony
}

var BKorop科罗普 feud.Barony = &科罗普KoropBarony{}

func init() {
    f := BKorop科罗普.(*科罗普KoropBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korop",
		TitleName: "科罗普",
		TitleCode: "b_korop",
	}
}
