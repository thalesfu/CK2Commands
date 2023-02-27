package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科捷利尼科沃KotelnikovoBarony struct {
	feud.BaseBarony
}

var BKotelnikovo科捷利尼科沃 feud.Barony = &科捷利尼科沃KotelnikovoBarony{}

func init() {
    f := BKotelnikovo科捷利尼科沃.(*科捷利尼科沃KotelnikovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotelnikovo",
		TitleName: "科捷利尼科沃",
		TitleCode: "b_kotelnikovo",
	}
}
