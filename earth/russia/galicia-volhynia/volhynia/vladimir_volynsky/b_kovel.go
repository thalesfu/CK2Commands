package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科韦利KovelBarony struct {
	feud.BaseBarony
}

var BKovel科韦利 feud.Barony = &科韦利KovelBarony{}

func init() {
    f := BKovel科韦利.(*科韦利KovelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kovel",
		TitleName: "科韦利",
		TitleCode: "b_kovel",
	}
}
