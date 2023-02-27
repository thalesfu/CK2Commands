package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘陀罗迦KodalakaBarony struct {
	feud.BaseBarony
}

var BKodalaka拘陀罗迦 feud.Barony = &拘陀罗迦KodalakaBarony{}

func init() {
    f := BKodalaka拘陀罗迦.(*拘陀罗迦KodalakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kodalaka",
		TitleName: "拘陀罗迦",
		TitleCode: "b_kodalaka",
	}
}
